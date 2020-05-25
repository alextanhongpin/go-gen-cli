package gen

type Template struct {
	Name        string                 `yaml:"name,omitempty"`
	Description string                 `yaml:"description,omitempty"`
	Prompts     []*Prompt              `yaml:"prompts,omitempty"`
	Actions     []*Action              `yaml:"actions,omitempty"`
	Environment map[string]interface{} `yaml:"environment,omitempty"`
}
