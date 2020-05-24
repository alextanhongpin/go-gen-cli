package gen

type Template struct {
	Name        string    `yaml:"name,omitempty"`
	Description string    `yaml:"description,omitempty"`
	Template    string    `yaml:"template,omitempty"`
	Path        string    `yaml:"path,omitempty"`
	Prompts     []*Prompt `yaml:"prompts,omitempty"`
	Actions     []*Action `yaml:"actions,omitempty"`
}
