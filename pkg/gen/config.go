package gen

type Config struct {
	Templates []*Template `yaml:"templates"`
}

type Template struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Template    string    `yaml:"template"`
	Path        string    `yaml:"path"`
	Prompts     []*Prompt `yaml:"prompts"`
	Actions     []*Action `yaml:"actions"`
}

type Prompt struct {
	Name     string   `yaml:"name"`
	Message  string   `yaml:"message"`
	Type     string   `yaml:"type"`
	Required bool     `yaml:"required"`
	Options  []string `yaml:"options"`
}

type Action struct {
	Description string `yaml:"description"`
	Template    string `yaml:"template"`
	Path        string `yaml:"path"`
}
