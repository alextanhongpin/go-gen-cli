package gen

import "fmt"

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
		Actions:     make([]*Action, 0),
		Environment: make(map[string]interface{}),
	}
}

type ActionResult struct {
	Error  error
	Action *Action
}

func GenerateTemplates(tpl *Template) []ActionResult {
	var result []ActionResult
	for _, act := range tpl.Actions {
		err := Create(act.Template)
		result = append(result, ActionResult{
			Error:  err,
			Action: act,
		})
	}
	return result
}

func ClearTemplates(tpl *Template) []ActionResult {
	var result []ActionResult
	for _, act := range tpl.Actions {
		err := RemoveIfExists(act.Path)
		result = append(result, ActionResult{
			Error:  err,
			Action: act,
		})
	}
	return result
}

func (t *Template) ValidateEnvironment() []error {
	var errors []error
	for key, value := range t.Environment {
		if IsZero(value) {
			errors = append(errors, fmt.Errorf("env %s is specified but no value is provided", key))
		}
	}
	return errors
}
