package gen

import (
	"fmt"
	"os"
)

type Template struct {
	Name        string            `yaml:"name,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Prompts     []*Prompt         `yaml:"prompts,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Actions     []*Action         `yaml:"actions,omitempty"`
}

func NewTemplate(name string) *Template {
	return &Template{
		Name:        name,
		Description: fmt.Sprintf("%s action", name),
		Prompts:     make([]*Prompt, 0),
		Environment: make(map[string]string, 0),
		Actions:     make([]*Action, 0),
	}
}

func (t *Template) ParseEnvironment() []error {
	var errors []error
	for k, v := range t.Environment {
		v = os.ExpandEnv(v)
		if v == "" {
			errors = append(errors, fmt.Errorf("environmentError: %q is specified but no value is provided", k))
			continue
		}
		t.Environment[k] = v
	}
	return errors
}

func (t *Template) ParsePrompts() (map[string]interface{}, error) {
	return Prompts(t.Prompts)
}
