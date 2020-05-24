package gen

type Config struct {
	Commands []*Command `yaml:"commands"`
}

type Command struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Template    string `yaml:"template"`
	Path        string `yaml:"path"`
}
