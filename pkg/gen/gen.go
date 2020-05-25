package gen

import (
	"os"
	"sort"

	"gopkg.in/yaml.v2"
)

type Gen struct {
	cfg *Config
}

func New() *Gen {
	return &Gen{cfg: NewConfig()}
}

func (c *Gen) Read(path string) error {
	b, err := Read(path)
	if err != nil {
		return err
	}
	b = []byte(os.ExpandEnv(string(b)))

	err = yaml.Unmarshal(b, &c.cfg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Gen) Write(path string) error {
	b, err := yaml.Marshal(c.cfg)
	if err != nil {
		return err
	}
	return Overwrite(path, b)
}

func (c *Gen) FindTemplate(name string) *Template {
	for _, tpl := range c.cfg.Templates {
		if tpl.Name == name {
			return tpl
		}
	}
	return nil
}

func (c *Gen) RemoveTemplate(name string) {
	var templates []*Template
	for _, tpl := range c.cfg.Templates {
		if tpl.Name == name {
			continue
		}
		templates = append(templates, tpl)
	}
	c.cfg.Templates = templates
}

func (c *Gen) AddTemplate(tpl *Template) {
	c.cfg.Templates = append(c.cfg.Templates, tpl)
	sort.SliceStable(c.cfg.Templates, func(i, j int) bool {
		return c.cfg.Templates[i].Name < c.cfg.Templates[j].Name
	})
}

func (c *Gen) ListTemplates() []string {
	var result []string
	for _, tpl := range c.cfg.Templates {
		result = append(result, tpl.Name)
	}
	return result
}
