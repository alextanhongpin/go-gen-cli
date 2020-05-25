package gen

import (
	"os"
	"sort"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Templates []*Template `yaml:"templates"`
}

func NewConfig(name string) (*Config, error) {
	b, err := Read(name)
	if err != nil {
		return nil, err
	}
	b = []byte(os.ExpandEnv(string(b)))

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) Find(name string) *Template {
	for _, tpl := range c.Templates {
		if tpl.Name == name {
			return tpl
		}
	}
	return nil
}

func (c *Config) Remove(name string) bool {
	idx := c.Index(name)
	if idx == -1 {
		return false
	}
	c.Templates = append(c.Templates[:idx], c.Templates[idx+1:]...)
	return true
}

func (c *Config) Index(name string) int {
	for i, tpl := range c.Templates {
		if tpl.Name == name {
			return i
		}
	}
	return -1
}

func (c *Config) Add(tpl *Template) {
	c.Templates = append(c.Templates, tpl)
	c.Sort()
}

func (c *Config) Sort() {
	sort.SliceStable(c.Templates, func(i, j int) bool {
		return c.Templates[i].Name < c.Templates[j].Name
	})
}

func (c *Config) ListTemplates() []string {
	var result []string
	for _, tpl := range c.Templates {
		result = append(result, tpl.Name)
	}
	return result
}

func (c *Config) WriteConfigIfNotExists(name string) error {
	f, err := Open(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return f.Write(b)
}
