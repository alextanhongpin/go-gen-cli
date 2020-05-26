package gen

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
)

type Template struct {
	Name        string            `yaml:"name,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Prompts     []*Prompt         `yaml:"prompts,omitempty"`
	Actions     []*Action         `yaml:"actions,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
}

func NewTemplate(name string) *Template {
	return &Template{
		Name:        name,
		Description: fmt.Sprintf("%s template", name),
		Prompts:     make([]*Prompt, 0),
		Actions: []*Action{
			NewAction(name),
			NewAction(fmt.Sprintf("%s_test", name)),
		},
		Environment: map[string]string{
			"PKG": "$PKG",
		},
	}
}

func (t *Template) ParseEnvironment() []error {
	var errors []error
	for key, value := range t.Environment {
		if value == "" {
			errors = append(errors, fmt.Errorf("env %s is specified but no value is provided", key))
		} else {
			t.Environment[key] = os.ExpandEnv(value)
		}
	}
	return errors
}

func (t *Template) ParsePrompts() (map[string]interface{}, error) {
	return Prompts(t.Prompts)
}

func ParseTemplate(b []byte, data interface{}) ([]byte, error) {
	t := template.Must(template.New("").Funcs(sprig.FuncMap()).Parse(string(b)))

	var bb bytes.Buffer
	if err := t.Execute(&bb, data); err != nil {
		return nil, err
	}

	return bb.Bytes(), nil
}
