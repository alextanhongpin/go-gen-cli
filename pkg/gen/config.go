package gen

type Config struct {
	Templates []*Template `yaml:"templates"`
}

func NewConfig() *Config {
	return &Config{
		Templates: make([]*Template, 0),
	}
}
