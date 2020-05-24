package gen

type Config struct {
	Templates []*Template `yaml:"templates"`
}

type Template struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Template    string `yaml:"template"`
	Path        string `yaml:"path"`
}
