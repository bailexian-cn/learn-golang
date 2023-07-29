package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	nurl "net/url"
	"strings"
	"time"
)

var HttpClientTimeout = 5 * time.Minute

func HttpCall(method, url string, h map[string]string, p map[string]string, body interface{}, res interface{}) error {
	reqStr, err := json.Marshal(body)
	if err != nil {
		return err
	}
	reader := strings.NewReader(string(reqStr))
	u, err := nurl.Parse(url)
	var client *http.Client
	if u.Scheme == "http" {
		client = &http.Client{
			Timeout: HttpClientTimeout,
		}
	} else if u.Scheme == "https" {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: HttpClientTimeout,
		}
	} else {
		return fmt.Errorf("cann't suppot the scheme %s", u.Scheme)
	}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest(method, u.String(), reader)
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")
	for k, v := range h {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return RequestErr{err}
	}
	if resp.StatusCode == http.StatusNotFound {
		return UrlNotFoundErr{fmt.Errorf("404 Not Found")}
	}
	respByteArr, err := io.ReadAll(resp.Body)
	if err != nil {
		return ReadRespErr{
			error: err,
			Resp:  string(respByteArr),
		}
	}
	log.Printf("url: %s\nmethod: %s\nheader: %s\nparams: %s\nbody: %s\nresponse: %s\n",
		url, method, h, p, string(reqStr), string(respByteArr))
	err = json.Unmarshal(respByteArr, res)
	if err != nil {
		return ParseRespErr{
			error: err,
			Resp:  string(respByteArr),
		}
	}
	return nil
}
