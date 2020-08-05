package gen

import (
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig"
)

func ParseTemplate(b []byte, data interface{}) ([]byte, error) {
	t := template.Must(template.New("").
		Funcs(sprig.FuncMap()).
		Funcs(template.FuncMap{
			"pascalcase": PascalCase,
			// Overrides sprig camelcase function, which is more like pascalcase.
			"camelcase": CamelCase,
		}).Parse(string(b)))

	var bb bytes.Buffer
	if err := t.Execute(&bb, data); err != nil {
		return nil, err
	}

	return bb.Bytes(), nil
}

func ParseString(s string, data interface{}) (string, error) {
	b, err := ParseTemplate([]byte(s), data)
	return string(b), err
}
