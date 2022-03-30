package http

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

type StubReader struct {
	io.Reader
}

func (s *StubReader) Close() error {
	return nil
}

func NewStubReader(str string) *StubReader {
	return &StubReader{
		strings.NewReader(str),
	}
}

func TestResponse_String(t *testing.T) {
	cases := []struct {
		name   string
		expect string
		reader *StubReader
	}{
		{
			"normal case",
			"string",
			NewStubReader("string"),
		},
		{
			"empty case",
			"",
			NewStubReader(""),
		},
		{
			"nil case",
			"",
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := &Response{Body: c.reader}
			assert.Equal(t, c.expect, r.String())
		})
	}
}
