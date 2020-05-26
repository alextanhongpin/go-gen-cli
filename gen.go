package gen

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Gen struct {
	cfgPath string
}

func New(cfgPath string) *Gen {
	return &Gen{
		cfgPath: cfgPath,
	}
}

func (g *Gen) Resolve(name string) string {
	if name == g.cfgPath {
		return name
	}
	dir := filepath.Dir(g.cfgPath)
	return filepath.Join(dir, name)
}

func (g *Gen) Write(w io.Writer, content []byte) (int, error) {
	return w.Write(content)
}

func (g *Gen) Read(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

func (g *Gen) Touch(name string) error {
	rel := g.Resolve(name)

	if err := os.MkdirAll(path.Dir(rel), os.ModePerm); err != nil {
		return err
	}

	f, err := os.OpenFile(rel, os.O_RDONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	return nil
}

func (g *Gen) ReadWriteFile(name string) (*os.File, error) {
	rel := g.Resolve(name)
	return os.OpenFile(rel, os.O_RDWR, 0644)
}

func (g *Gen) ReadOnlyFile(name string) (*os.File, error) {
	rel := g.Resolve(name)
	return os.OpenFile(rel, os.O_RDONLY, 0644)
}

func (g *Gen) WriteOnlyFile(name string) (*os.File, error) {
	rel := g.Resolve(name)
	if err := os.MkdirAll(path.Dir(rel), os.ModePerm); err != nil {
		return nil, err
	}
	return os.OpenFile(rel, os.O_WRONLY|os.O_TRUNC, 0644)
}

func (g *Gen) UnmarshalConfig(b []byte, cfg *Config) error {
	return yaml.Unmarshal(b, &cfg)
}

func (g *Gen) MarshalConfig(cfg *Config) ([]byte, error) {
	return yaml.Marshal(&cfg)
}

func (g *Gen) Remove(name string) error {
	rel := g.Resolve(name)
	return os.Remove(rel)
}

func (g *Gen) LoadConfig() (*Config, error) {
	f, err := g.ReadOnlyFile(g.cfgPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := g.Read(f)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = g.UnmarshalConfig(b, &cfg)
	return &cfg, err
}

func (g *Gen) WriteConfig(cfg *Config) error {
	f, err := g.WriteOnlyFile(g.cfgPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := g.MarshalConfig(cfg)
	if err != nil {
		return err
	}

	_, err = g.Write(f, b)
	return err
}

func (g *Gen) Copy(r io.Reader, w io.Writer, parser func([]byte) ([]byte, error)) error {
	b, err := g.Read(r)
	if err != nil {
		return err
	}

	b, err = parser(b)
	if err != nil {
		return err
	}

	_, err = g.Write(w, b)
	return err
}
