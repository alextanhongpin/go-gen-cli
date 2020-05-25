package gen

import "sort"

type Config struct {
	Templates []*Template `yaml:"templates"`
}

func (c *Config) Find(name string) *Template {
	for _, tpl := range c.Templates {
		if tpl.Name == name {
			return tpl
		}
	}
	return nil
}

func (c *Config) Remove(name string) {
	idx := c.Index(name)
	if idx == -1 {
		return
	}
	c.Templates = append(c.Templates[:idx], c.Templates[idx+1:]...)
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
