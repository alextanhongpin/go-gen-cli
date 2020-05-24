package gen

import (
	"log"

	"github.com/tj/survey"
)

func Prompts(prompts []*Prompt) (map[string]interface{}, error) {
	var qs []*survey.Question
	answers := make(map[string]interface{})
	for _, p := range prompts {
		switch p.Type {
		case "input":
			q := &survey.Question{
				Name:   p.Name,
				Prompt: &survey.Input{Message: p.Message},
			}
			if p.Required {
				q.Validate = survey.Required
			}
			qs = append(qs, q)
		default:
			log.Println("not supported", p.Type)
		}
	}
	if err := survey.Ask(qs, &answers); err != nil {
		return nil, err
	}
	return answers, nil
}
