package gen

import "fmt"

const (
	TEMPLATE_PATH = "templates"
	PKG_PATH      = "pkg"
)

type Action struct {
	Description string `yaml:"description"`
	Template    string `yaml:"template"`
	Path        string `yaml:"path"`
}

func NewAction(name string) *Action {
	return &Action{
		Description: fmt.Sprintf("creates a %s", name),
		Template:    fmt.Sprintf("%s/%s.tmpl", TEMPLATE_PATH, name),
		Path:        fmt.Sprintf("%s/%s.go", PKG_PATH, name),
	}
}
