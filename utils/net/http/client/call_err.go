package client

import (
	"net"
	"net/url"
)

// HttpCallErrType HttpCall函数错误分类
type HttpCallErrType string

// 创建类接口 TypeRequestTimeout、TypeParseRespErr、TypeReadRespErr时直接报错

const (
	TypeNil            = HttpCallErrType("Nil")
	TypeNetErr         = HttpCallErrType("NetErr")
	TypeRequestTimeout = HttpCallErrType("Request Timeout")
	TypeUrlNotFound    = HttpCallErrType("404 Not Found")
	TypeReadRespErr    = HttpCallErrType("Read Response Err")
	TypeParseRespErr   = HttpCallErrType("Parse Response Err")
	TypeDefault        = HttpCallErrType("Default")
)

func JudgeErrType(err error) HttpCallErrType {
	if err == nil {
		return TypeNil
	}
	if requestErr, ok := err.(RequestErr); ok {
		if urlErr, ok := requestErr.error.(*url.Error); ok && urlErr.Timeout() {
			return TypeRequestTimeout
		}
		if _, ok1 := requestErr.error.(net.Error); ok1 {
			return TypeNetErr
		}
	}

	if _, ok := err.(UrlNotFoundErr); ok {
		return TypeUrlNotFound
	}
	if _, ok := err.(ReadRespErr); ok {
		return TypeReadRespErr
	}
	if _, ok := err.(ParseRespErr); ok {
		return TypeParseRespErr
	}
	return TypeDefault
}

type UrlNotFoundErr struct {
	error
}

type RequestErr struct {
	error
}

type ParseRespErr struct {
	error
	Resp string
}

type ReadRespErr struct {
	error
	Resp string
}
