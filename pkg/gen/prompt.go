package gen

import (
	"log"

	"github.com/tj/survey"
)

type Prompt struct {
	Name     string   `yaml:"name"`
	Message  string   `yaml:"message"`
	Help     string   `yaml:"help"`
	Default  string   `yaml:"default"`
	Selected []string `yaml:"selected"`
	Type     string   `yaml:"type"`
	Confirm  bool     `yaml:"confirm"`
	PageSize int      `yaml:"page_size"`
	Required bool     `yaml:"required"`
	Options  []string `yaml:"options"`
}

func Prompts(prompts []*Prompt) (map[string]interface{}, error) {
	var qs []*survey.Question
	answers := make(map[string]interface{})
	for _, p := range prompts {
		q := &survey.Question{
			Name: p.Name,
		}
		if p.Required {
			q.Validate = survey.Required
		}

		switch p.Type {
		case "input":
			q.Prompt = &survey.Input{
				Message: p.Message,
				Default: p.Default,
				Help:    p.Help,
			}
		case "select":
			q.Prompt = &survey.Select{
				Message: p.Message,
				Options: p.Options,
				Default: p.Default,
				Help:    p.Help,
			}
		case "password":
			q.Prompt = &survey.Password{
				Message: p.Message,
				Help:    p.Help,
			}
		case "multiselect":
			q.Prompt = &survey.MultiSelect{
				Message:  p.Message,
				Options:  p.Options,
				Default:  p.Selected,
				Help:     p.Help,
				PageSize: p.PageSize,
			}
		case "confirm":
			q.Prompt = &survey.Confirm{
				Message: p.Message,
				Default: p.Confirm,
				Help:    p.Help,
			}
		default:
			log.Println("not supported", p.Type)
		}
		qs = append(qs, q)
	}
	if err := survey.Ask(qs, &answers); err != nil {
		return nil, err
	}
	return answers, nil
}
