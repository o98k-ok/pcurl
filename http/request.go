package http

import (
	"github.com/o98k-ok/pcurl/utils"
	"io"
	"net/http"
)

type Request struct {
	URL    string
	Method string
	Params map[string]string
	Header http.Header
	Data   io.Reader
	cli    *http.Client
}

func NewRequest(cli *http.Client, url string) *Request {
	return &Request{
		URL:    url,
		Method: http.MethodGet,
		Params: map[string]string{},
		Header: http.Header{},
		Data:   nil,
		cli:    cli,
	}
}

func (r *Request) WithMethod(method string) *Request {
	r.Method = method
	return r
}

func (r *Request) AddParam(k, v string) *Request {
	r.Params[k] = v
	return r
}

func (r *Request) AddHeader(k, v string) *Request {
	r.Header.Set(k, v)
	return r
}

func (r *Request) WithData(reader io.Reader) *Request {
	r.Data = reader
	return r
}

func (r *Request) FullURL() string {
	return r.URL + "?" + utils.Convert(r.Params, "=", "&")
}

func (r *Request) Do() (*Response, error) {
	var (
		request  *http.Request
		response *http.Response
		err      error
	)

	request, err = http.NewRequest(r.Method, r.FullURL(), r.Data)
	if err != nil {
		return nil, err
	}
	request.Header = r.Header

	response, err = r.cli.Do(request)
	if err != nil {
		return nil, err
	}

	return &Response{
		Code:    response.StatusCode,
		Headers: response.Header,
		Body:    response.Body,
	}, nil
}
