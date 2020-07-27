package gen_test

import (
	"testing"

	"github.com/alextanhongpin/go-gen-cli"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := gen.NewConfig()
	assert.NotNil(cfg)
	assert.True(len(cfg.Templates) == 0)
}

func TestConfigFind(t *testing.T) {
	t.Run("returns nil when there are no matches", func(t *testing.T) {
		assert := assert.New(t)

		cfg := gen.NewConfig()
		tpl := cfg.Find("hello world")
		assert.Nil(tpl)
	})

	t.Run("returns template when there are matches", func(t *testing.T) {
		assert := assert.New(t)

		cfg := gen.NewConfig()
		cfg.Templates = append(cfg.Templates, &gen.Template{
			Name: "domain",
		})
		tpl := cfg.Find("domain")
		assert.NotNil(tpl)
		assert.Equal("domain", tpl.Name)
	})
}
