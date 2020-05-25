package gen

import (
	"errors"
	"fmt"
)

const (
	TEMPLATE_PATH = "templates"
	PKG_PATH      = "pkg"
)

var (
	ErrEmpty = errors.New("file is empty")
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

func (a *Action) TouchTemplate() error {
	return Touch(a.Template)
}

func (a *Action) RemoveTemplate() error {
	return Remove(a.Template)
}

func (a *Action) RemoveGeneratedFile() error {
	return Remove(a.Path)
}

// Generate reads from the template and write to the destination.
func (a *Action) Exec(data interface{}) error {
	b, err := Read(a.Template)
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return ErrEmpty
	}
	return Write(a.Path, b, data)
}
