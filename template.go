package gen

import (
	"bytes"
	"fmt"
	"html/template"
)

type Template struct {
	Name        string                 `yaml:"name,omitempty"`
	Description string                 `yaml:"description,omitempty"`
	Prompts     []*Prompt              `yaml:"prompts,omitempty"`
	Actions     []*Action              `yaml:"actions,omitempty"`
	Environment map[string]interface{} `yaml:"environment,omitempty"`
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
		Environment: map[string]interface{}{
			"PKG": "$PKG",
		},
	}
}

func (t *Template) ParseEnvironment() []error {
	var errors []error
	for key, value := range t.Environment {
		if IsZero(value) {
			errors = append(errors, fmt.Errorf("env %s is specified but no value is provided", key))
		}
	}
	return errors
}

func (t *Template) ParsePrompts() (map[string]interface{}, error) {
	return Prompts(t.Prompts)
}

func ParseTemplate(b []byte, data interface{}) ([]byte, error) {
	t := template.Must(template.New("").Parse(string(b)))

	var bb bytes.Buffer
	if err := t.Execute(&bb, data); err != nil {
		return nil, err
	}

	return bb.Bytes(), nil
}
