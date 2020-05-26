package gen

import "sort"

type Config struct {
	Templates []*Template `yaml:"templates,omitempty"`
}

func NewConfig() *Config {
	return &Config{
		Templates: make([]*Template, 0),
	}
}

func (c *Config) Add(tpl *Template) {
	c.Templates = append(c.Templates, tpl)
	sort.SliceStable(c.Templates, func(i, j int) bool {
		return c.Templates[i].Name < c.Templates[j].Name
	})
}

func (c *Config) Remove(name string) {
	var templates []*Template
	for _, tpl := range c.Templates {
		if tpl.Name == name {
			continue
		}
		templates = append(templates, tpl)
	}
	c.Templates = templates
}

func (c *Config) Find(name string) *Template {
	for _, tpl := range c.Templates {
		if tpl.Name == name {
			return tpl
		}
	}
	return nil
}
