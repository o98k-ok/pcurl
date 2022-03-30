package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func assertResponse(t *testing.T, resp *Response, expect string) {
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, expect, resp.String())
}

func TestRequest(t *testing.T) {
	cli := &http.Client{}
	t.Run("normal get cases", func(t *testing.T) {
		expect := "ok"
		srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, expect)
		}))

		resp, _ := NewRequest(cli, srv.URL).Do()
		assertResponse(t, resp, expect)
	})

	t.Run("normal get with params cases", func(t *testing.T) {
		key, value, expect := "id", "110", "ok"
		srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			v := request.URL.Query().Get(key)
			if v == value {
				fmt.Fprint(writer, expect)
			}
		}))

		resp, _ := NewRequest(cli, srv.URL).AddParam(key, value).Do()
		assertResponse(t, resp, expect)
	})

	t.Run("normal post with data", func(t *testing.T) {
		expect := "ok"
		srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			d, err := ioutil.ReadAll(request.Body)
			if err == nil && request.Method == http.MethodPost {
				fmt.Fprint(writer, string(d))
			}
		}))

		resp, _ := NewRequest(cli, srv.URL).WithMethod(http.MethodPost).WithData(strings.NewReader(expect)).Do()
		assertResponse(t, resp, expect)
	})

	t.Run("normal test with header", func(t *testing.T) {
		key, value, expect := "id", "110", "ok"
		srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if val := request.Header.Get(key); val == value {
				fmt.Fprint(writer, expect)
			}
		}))
		resp, _ := NewRequest(cli, srv.URL).AddHeader(key, value).Do()
		assertResponse(t, resp, expect)
	})
}
