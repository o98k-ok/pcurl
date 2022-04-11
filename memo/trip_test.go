package memo

import (
	"bytes"
	"fmt"
	"github.com/o98k-ok/pcurl/format"
	hp "github.com/o98k-ok/pcurl/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrip(t *testing.T) {
	cli := &http.Client{}
	result := "{\"name\":\"ok\"}"
	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, result)
	}))

	req := hp.NewRequest(cli, srv.URL).WithMethod(http.MethodGet)
	buf := &bytes.Buffer{}
	rt := &BaseTrip{request: req, formatter: format.NewFormatter(buf).WithoutColor()}

	var err error
	err = rt.Load()
	assert.NoError(t, err)

	err = rt.Do()
	assert.NoError(t, err)

	err = rt.Sink()
	assert.NoError(t, err)
	assert.Equal(t, `{
    "name": "ok"
}`, buf.String())
}
