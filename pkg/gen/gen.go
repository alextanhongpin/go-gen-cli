package gen

import "gopkg.in/yaml.v2"

func OverwriteConfig(name string, cfg *Config) error {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return Overwrite(name, b)
}
