package gen_test

import (
	"testing"

	genlib "github.com/alextanhongpin/go-gen-cli"
	"github.com/alextanhongpin/go-gen-cli/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	path := "./testdata/gen.yaml"
	g := genlib.New(path)
	assert.Nil(g.Remove(path))
	err := gen.Init(path, false)
	assert.Nil(err)
}
