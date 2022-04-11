package format

import (
	"encoding/json"
	"github.com/o98k-ok/lazy/v2/format"
	"io"
)

type Formatter struct {
	f *format.JsonFormatter
}

func NewFormatter(w io.Writer) *Formatter {
	f := format.NewEncoder(w).StringColor(format.Green).MapKeyColor(format.Blue).
		NumberColor(format.Yellow).BoolColor(format.Cyan)
	return &Formatter{
		f: f,
	}
}

func (f *Formatter) WithIndent(indent string) *Formatter {
	f.f.WithIndent(indent)
	return f
}

func (f *Formatter) WithoutColor() *Formatter {
	f.f.DisableColor()
	return f
}

func (f *Formatter) Jsonify(reader io.Reader) error {
	var obj map[string]interface{}
	err := json.NewDecoder(reader).Decode(&obj)
	if err != nil {
		return err
	}

	return f.f.Encode(obj)
}
