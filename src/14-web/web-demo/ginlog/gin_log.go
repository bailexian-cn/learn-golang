// Package ginzap provides log handling using zap package.
// Code structure based on ginrus package.
package ginlog

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

var MaxBodyLogLen = 4 * 1024 // 4 kB
const RequestIDHeader = "X-Request-ID"
const SpanIDHeader = "X-Span-ID"

// BodyLogWriter是为了记录返回数据到log中进行了双写
type BodyLogWriter struct {
	gin.ResponseWriter
	body      *bytes.Buffer
	writedLen int
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func (w *BodyLogWriter) Write(b []byte) (int, error) {
	if w.writedLen < MaxBodyLogLen {
		n, _ := w.body.Write(b[:min(MaxBodyLogLen-w.writedLen, len(b))])
		w.writedLen += n
	}
	return w.ResponseWriter.Write(b)
}

func (w *BodyLogWriter) String() string {
	return w.body.String()
}

func (w *BodyLogWriter) Bytes() []byte {
	return w.body.Bytes()
}

type BodyLogReader struct {
	reader     io.ReadCloser
	data       *bytes.Buffer
	preReadLen int
	reReadLen  int
}

func (r *BodyLogReader) PreRead() {
	temp := make([]byte, MaxBodyLogLen)
	var err error
	n := 0
	r.preReadLen = 0
	for {
		n, err = r.reader.Read(temp[r.preReadLen:])
		r.data.Write(temp[r.preReadLen : r.preReadLen+n])
		r.preReadLen += n
		if r.preReadLen >= MaxBodyLogLen || err == io.EOF {
			break
		}
	}
}

func (r *BodyLogReader) Read(p []byte) (n int, err error) {
	if r.reReadLen < r.preReadLen {
		n, err = r.data.Read(p)
		r.reReadLen += n
	} else {
		return r.reader.Read(p)
	}
	return
}
func (r BodyLogReader) Close() error {
	return r.reader.Close()
}

func (r BodyLogReader) String() string {
	return r.data.String()
}

type key int

const (
	KeyHttpCtx key = iota
	RequestID
	SpanID    //string ecs0.ebs1.ecs0.ebs1
	SpanIndex // int 0 1 2 3
	TaskID
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

func randStr(length int) string {
	b := make([]rune, length)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ExtractCtxValue(oldctx context.Context) context.Context {
	ctx := context.WithValue(context.Background(), RequestID, GetRequestID(oldctx))
	ctx = context.WithValue(ctx, SpanID, GetSpanID(oldctx))
	ctx = context.WithValue(ctx, SpanIndex, getSpanIndex(oldctx))
	ctx = context.WithValue(ctx, TaskID, GetTaskID(oldctx))
	return ctx
}

func GetRequestID(ctx context.Context) string {
	val := ctx.Value(RequestID)
	if val == nil {
		return ""
	}
	if id, ok := val.(string); ok {
		return id
	}
	return ""
}

func GetTaskID(ctx context.Context) string {
	val := ctx.Value(TaskID)
	if val == nil {
		return ""
	}
	if id, ok := val.(string); ok {
		return id
	}
	return ""
}

func GetContext(c *gin.Context) context.Context {
	if HTTPCtx, ok := c.Get("ctx"); ok {
		if ctx, ok := HTTPCtx.(context.Context); ok {
			return ctx
		}
	}
	return c.Request.Context()
}

func NewSpan(ctx context.Context, name string) context.Context {
	var spanIndex uint32 = 0
	ctx = context.WithValue(ctx, SpanIndex, &spanIndex)
	return context.WithValue(ctx, SpanID, GetSpanID(ctx)+"."+name)
}

func NextSpan(ctx context.Context) context.Context {
	var newSpanIndex uint32 = 0
	spanIndex := getSpanIndex(ctx)
	if spanIndex != nil {
		newSpanIndex = atomic.AddUint32(spanIndex, 1)
	}
	oldSpanID := GetSpanID(ctx)
	newSpanID := oldSpanID + fmt.Sprintf("%d", newSpanIndex)
	return context.WithValue(ctx, SpanID, newSpanID)
}

func getSpanIndex(ctx context.Context) *uint32 {
	val := ctx.Value(SpanIndex)
	if val == nil {
		return nil
	}
	if spanIndex, ok := val.(*uint32); ok {
		if spanIndex != nil {
			return spanIndex
		}
	}
	return nil
}

func GetSpanID(ctx context.Context) string {
	val := ctx.Value(SpanID)
	if val == nil {
		return ""
	}
	if spanID, ok := val.(string); ok {
		return spanID
	}
	return ""
}

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Ginzap(logger *zap.Logger, serviceName string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		reader := BodyLogReader{reader: c.Request.Body, data: bytes.NewBufferString("")}
		contentType := c.Request.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/octet-stream") {
			reader.PreRead()
			c.Request.Body = &reader
		}

		blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		requestID := c.Request.Header.Get(RequestIDHeader)
		if len(requestID) == 0 {
			requestID = uuid.Must(uuid.NewV4()).String()
		}
		spanID := c.Request.Header.Get(SpanIDHeader)
		if len(spanID) == 0 {
			spanID = randStr(5)
		}
		ctx := context.WithValue(c.Request.Context(), RequestID, requestID)
		ctx = context.WithValue(ctx, SpanID, spanID)
		ctx = NewSpan(ctx, serviceName)
		c.Set("ctx", ctx)
		c.Set("request", reader)
		c.Set("response", blw)
		// request
		logger.Info("request",
			zap.String("request_id", GetRequestID(ctx)),
			zap.String("span_id", GetSpanID(ctx)),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.ByteString("parameter", reader.data.Bytes()),
		)

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			// response log
			logger.Info("response",
				zap.String("request_id", GetRequestID(ctx)),
				zap.String("span_id", GetSpanID(ctx)),
				zap.String("method", c.Request.Method),
				zap.Duration("latency", latency),
				zap.ByteString("response", blw.body.Bytes()),
			)
		}
	}
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.String("request_id", GetRequestID(GetContext(c))),
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.String("request_id", GetRequestID(GetContext(c))),
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
