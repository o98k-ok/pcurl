package memo

import (
	"github.com/o98k-ok/pcurl/format"
	hp "github.com/o98k-ok/pcurl/http"
)

type RoundTrip interface {
	Load() error
	Do() error
	Sink() error
}

type BaseTrip struct {
	request   *hp.Request
	response  *hp.Response
	formatter *format.Formatter
}

func (b *BaseTrip) Load() error {
	return nil
}

func (b *BaseTrip) Do() error {
	var err error
	b.response, err = b.request.Do()
	return err
}

func (b *BaseTrip) Sink() error {
	return b.formatter.Jsonify(b.response.Body)
}
