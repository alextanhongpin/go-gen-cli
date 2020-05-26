package gen

import (
	"fmt"

	"github.com/tj/survey"
)

type Prompt struct {
	Name         string   `yaml:"name,omitempty"`
	Message      string   `yaml:"message,omitempty"`
	Help         string   `yaml:"help,omitempty"`
	Default      string   `yaml:"default,omitempty"`
	Selected     []string `yaml:"selected,omitempty"`
	Type         string   `yaml:"type,omitempty"`
	Confirmation bool     `yaml:"confirm,omitempty"`
	PageSize     int      `yaml:"page_size,omitempty"`
	Required     bool     `yaml:"required,omitempty"`
	Options      []string `yaml:"options,omitempty"`
}

func (p *Prompt) Input() *survey.Question {
	q := &survey.Question{
		Name: p.Name,
		Prompt: &survey.Input{
			Message: p.Message,
			Default: p.Default,
			Help:    p.Help,
		},
	}
	if p.Required {
		q.Validate = survey.Required
	}
	return q
}

func (p *Prompt) Select() *survey.Question {
	q := &survey.Question{
		Name: p.Name,
		Prompt: &survey.Select{
			Message: p.Message,
			Options: p.Options,
			Default: p.Default,
			Help:    p.Help,
		},
	}
	if p.Required {
		q.Validate = survey.Required
	}
	return q

}

func (p *Prompt) Password() *survey.Question {
	q := &survey.Question{
		Name: p.Name,
		Prompt: &survey.Password{
			Message: p.Message,
			Help:    p.Help,
		},
	}
	if p.Required {
		q.Validate = survey.Required
	}
	return q
}

func (p *Prompt) MultiSelect() *survey.Question {
	q := &survey.Question{
		Name: p.Name,
		Prompt: &survey.MultiSelect{
			Message:  p.Message,
			Options:  p.Options,
			Default:  p.Selected,
			Help:     p.Help,
			PageSize: p.PageSize,
		},
	}
	if p.Required {
		q.Validate = survey.Required
	}
	return q
}

func (p *Prompt) Confirm() *survey.Question {
	q := &survey.Question{
		Name: p.Name,
		Prompt: &survey.Confirm{
			Message: p.Message,
			Default: p.Confirmation,
			Help:    p.Help,
		},
	}
	if p.Required {
		q.Validate = survey.Required
	}
	return q
}

func Prompts(prompts []*Prompt) (map[string]interface{}, error) {
	var qs []*survey.Question
	answers := make(map[string]interface{})
	for _, p := range prompts {
		var q *survey.Question
		switch p.Type {
		case "input":
			q = p.Input()
		case "select":
			q = p.Select()
		case "password":
			q = p.Password()
		case "multiselect":
			q = p.MultiSelect()
		case "confirm":
			q = p.Confirm()
		default:
			return nil, fmt.Errorf("prompt %q is not supported", p.Name)
		}
		qs = append(qs, q)
	}
	if err := survey.Ask(qs, &answers); err != nil {
		return nil, err
	}
	return answers, nil
}
