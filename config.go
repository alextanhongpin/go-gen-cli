package gen

type Config struct {
	Version   string      `yaml:"version"`
	Templates []*Template `yaml:"templates,omitempty"`
}

func NewConfig() *Config {
	return &Config{
		Templates: make([]*Template, 0),
	}
}

func (c *Config) Find(name string) *Template {
	for _, tpl := range c.Templates {
		if tpl.Name == name {
			return tpl
		}
	}
	return nil
}
