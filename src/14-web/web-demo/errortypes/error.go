package errortypes

import "github.com/pkg/errors"

var SqlNotFound = errors.New("sql not found")
var SqlDuplicate = errors.New("sql duplicate")
var OPNotFound = errors.New("op resource not found")
var OPDeleteNotPermit = errors.New("delete resource not permit")
var CrdNotFound = errors.New("crd data not found")
