package gen_test

import (
	"errors"
	"testing"

	"github.com/alextanhongpin/go-gen-cli"
	"github.com/stretchr/testify/assert"
)

func TestMultiErrorFactory(t *testing.T) {
	t.Run("returns empty error when initialize with no input", func(t *testing.T) {

		assert := assert.New(t)
		merr := gen.NewMultiError()
		assert.False(merr.HasError())
		assert.Equal("", merr.Error())
	})

	t.Run("returns error when initialize with multiple error", func(t *testing.T) {
		assert := assert.New(t)
		merr := gen.NewMultiError(errors.New("error 1"), errors.New("error 2"))
		assert.True(merr.HasError())
		assert.Equal("error 1\nerror 2", merr.Error())
	})
}

func TestMultiErrorAddError(t *testing.T) {
	assert := assert.New(t)
	merr := gen.NewMultiError()
	assert.True(merr.Add(errors.New("error 1")))
	assert.True(merr.Add(errors.New("error 2")))
	assert.False(merr.Add(nil))
	assert.True(merr.HasError())
	assert.Equal("error 1\nerror 2", merr.Error())
}
