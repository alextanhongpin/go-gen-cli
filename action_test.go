package gen_test

import (
	"os"
	"testing"

	"github.com/alextanhongpin/go-gen-cli"
	"github.com/stretchr/testify/assert"
)

func TestAction(t *testing.T) {
	act := gen.NewAction("my_service", "/templates/service.tmpl:/dst/service.go")

	t.Run("NewAction returns pointer to Action", func(t *testing.T) {
		assert := assert.New(t)
		assert.NotNil(act)
		assert.Equal("my_service", act.Name)
		assert.Equal("/templates/service.tmpl:/dst/service.go", act.Path)
	})

	t.Run("when path is valid returns source and destination", func(t *testing.T) {
		assert := assert.New(t)
		assert.Equal("/templates/service.tmpl", act.Source())
		assert.Equal("/dst/service.go", act.Destination())
	})

	t.Run("when path is invalid returns empty source and destination", func(t *testing.T) {
		assert := assert.New(t)

		act := gen.NewAction("my_service", "")
		assert.Equal("", act.Source())
		assert.Equal("", act.Destination())
	})
}

func TestActionResolve(t *testing.T) {
	assert := assert.New(t)

	err := os.Setenv("FROM_ENV", "hardcoded-src")
	assert.Nil(err)

	act := gen.NewAction("my_service", "/templates/{{ .Source }}/service.tmpl:/dst/{{ .Destination }}/service.go")
	act.Variables = map[string]interface{}{
		"Source":      `$FROM_ENV`,
		"Destination": "{{ .FromEnvironment }}",
	}
	newAct, err := act.Resolve(map[string]interface{}{
		"FromEnvironment": "hardcoded-dst",
	})
	assert.Nil(err)
	assert.Equal("/templates/hardcoded-src/service.tmpl", newAct.Source())
	assert.Equal("/dst/hardcoded-dst/service.go", newAct.Destination())
}
