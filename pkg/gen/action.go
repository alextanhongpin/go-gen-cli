package gen

import "fmt"

type Action struct {
	Description string `yaml:"description"`
	Template    string `yaml:"template"`
	Path        string `yaml:"path"`
}

func NewAction(name string) *Action {
	return &Action{
		Description: fmt.Sprintf("creates a %s", name),
		Template:    fmt.Sprintf("templates/%s.go", name),
		Path:        fmt.Sprintf("pkg/%s.go", name),
	}
}
