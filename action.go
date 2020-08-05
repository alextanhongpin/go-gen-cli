package gen

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrInvalidPath = errors.New("path is invalid")

type Action struct {
	Name      string
	Path      string
	Variables map[string]interface{}
}

func NewAction(name, path string) *Action {
	return &Action{
		Name: name,
		// We do not expand environment here, because the file is loaded from yaml.
		Path:      path,
		Variables: make(map[string]interface{}, 0),
	}
}

func (a *Action) splitPath() (src, dst string) {
	paths := filepath.SplitList(a.Path)
	if len(paths) != 2 {
		return
	}
	src, dst = paths[0], paths[1]
	return
}

func (a *Action) Source() string {
	src, _ := a.splitPath()
	return src
}

func (a *Action) Destination() string {
	_, dst := a.splitPath()
	return dst
}

// Resolve returns a new Action with the path and variables expanded with the environment variables
// and populated with the template data.
func (a *Action) Resolve(data map[string]interface{}) (*Action, error) {
	// Expand the variables.
	variables := make(map[string]interface{}, 0)
	// Copy the data to variables, overwrite if they have the same key.
	for k, v := range data {
		variables[k] = v
	}
	for k, v := range a.Variables {
		s, ok := v.(string)
		if ok {
			val, err := ParseString(os.ExpandEnv(s), data)
			if err != nil {
				return nil, err
			}
			if val == "" {
				return nil, fmt.Errorf("variableError: value is required for key %q", k)
			}
			variables[k] = val
		}
	}

	path, err := ParseString(a.Path, variables)
	if err != nil {
		return nil, err
	}

	act := NewAction(a.Name, path)
	act.Variables = variables
	return act, nil
}
