package ginlog

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Env        string
	LogPath    string
	MaxSize    int  // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int  // 日志文件最多保存多少个备份
	MaxAge     int  // 文件最多保存多少天
	Compress   bool // 是否压缩
}

func (config *Config) setDefault() {
	if config.MaxSize == 0 {
		config.MaxSize = 1024
	}
	if config.MaxBackups == 0 {
		config.MaxBackups = 5
	}
	if config.MaxAge == 0 {
		config.MaxAge = 365
	}
}

const (
	EnvProduct string = "product"
	EnvDevelop string = "develop"
	EnvDebug   string = "debug"
)

var ZapLogger *zap.Logger
var loggerSourceDir = "ginlog/logger.go"

func FileWithLineNum() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.Contains(file, loggerSourceDir) || strings.HasSuffix(file, "_test.go")) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}

func InitLogger(config Config) {
	config.setDefault()
	zapEncoderConfig := zap.NewDevelopmentEncoderConfig()
	zapEncoder := zapcore.NewConsoleEncoder(zapEncoderConfig)
	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	if config.Env == EnvProduct {
		zapEncoderConfig = zap.NewProductionEncoderConfig()
		zapEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		zapEncoder = zapcore.NewJSONEncoder(zapEncoderConfig)
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
	var writers []zapcore.WriteSyncer
	writers = append(writers, zapcore.AddSync(os.Stderr))
	if len(config.LogPath) > 0 {
		hook := lumberjack.Logger{
			Filename:   config.LogPath,    // 日志文件路径
			MaxSize:    config.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: config.MaxBackups, // 日志文件最多保存多少个备份
			MaxAge:     config.MaxAge,     // 文件最多保存多少天
			Compress:   config.Compress,   // 是否压缩
		}
		writers = append(writers, zapcore.AddSync(&hook))
	}
	core := zapcore.NewCore(
		zapEncoder,                              // 编码器配置
		zapcore.NewMultiWriteSyncer(writers...), // 打印到控制台和文件
		level,
	)
	ZapLogger = zap.New(core, zap.AddStacktrace(zap.ErrorLevel))
}

type CxtZapLogger struct {
	Zaplogger *zap.Logger
	ctx       context.Context
}

func (logger *CxtZapLogger) Info(msg string, fields ...zap.Field) {
	appendFileds := []zap.Field{}
	appendFileds = append(appendFileds, zap.String("request_id", GetRequestID(logger.ctx)))
	appendFileds = append(appendFileds, zap.String("span_id", GetSpanID(logger.ctx)))
	appendFileds = append(appendFileds, zap.String("caller", FileWithLineNum()))
	taskID := GetTaskID(logger.ctx)
	if len(taskID) > 0 {
		appendFileds = append(appendFileds, zap.String("TaskID", taskID))
	}

	appendFileds = append(appendFileds, fields...)
	ZapLogger.Info(msg, appendFileds...)
}

func (logger *CxtZapLogger) Warn(msg string, fields ...zap.Field) {
	appendFileds := []zap.Field{}
	appendFileds = append(appendFileds, zap.String("request_id", GetRequestID(logger.ctx)))
	appendFileds = append(appendFileds, zap.String("span_id", GetSpanID(logger.ctx)))
	taskID := GetTaskID(logger.ctx)
	if len(taskID) > 0 {
		appendFileds = append(appendFileds, zap.String("TaskID", taskID))
	}
	appendFileds = append(appendFileds, fields...)
	ZapLogger.Warn(msg, appendFileds...)
}

func (logger *CxtZapLogger) Error(msg string, fields ...zap.Field) {
	appendFileds := []zap.Field{}
	appendFileds = append(appendFileds, zap.String("request_id", GetRequestID(logger.ctx)))
	appendFileds = append(appendFileds, zap.String("span_id", GetSpanID(logger.ctx)))
	taskID := GetTaskID(logger.ctx)
	if len(taskID) > 0 {
		appendFileds = append(appendFileds, zap.String("TaskID", taskID))
	}
	appendFileds = append(appendFileds, fields...)
	ZapLogger.Error(msg, appendFileds...)
}

func (logger *CxtZapLogger) Errorln(msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

func (logger *CxtZapLogger) Infoln(v ...interface{}) {
	logger.Info("println", zap.String("message", fmt.Sprintln(v...)))
}

func CtxLogger(ctx context.Context) *CxtZapLogger {
	return &CxtZapLogger{
		Zaplogger: ZapLogger,
		ctx:       ctx,
	}
}
