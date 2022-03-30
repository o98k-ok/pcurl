package http

import (
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Response struct {
	Code    int
	Headers http.Header
	Body    io.ReadCloser
}

func (r *Response) String() string {
	if r.Body == nil || reflect.ValueOf(r.Body).IsNil() {
		return ""
	}

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ""
	}

	return string(d)
}
