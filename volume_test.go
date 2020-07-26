package gen_test

import (
	"os"
	"testing"

	"github.com/alextanhongpin/go-gen-cli"
	"github.com/stretchr/testify/assert"
)

func TestVolume(t *testing.T) {
	assert := assert.New(t)

	vol := gen.Volume("/src:/dst")
	src, dst, err := vol.Split()

	assert.Equal("/src", src)
	assert.Equal("/dst", dst)
	assert.Nil(err)
}

func TestVolumeWithEnvironment(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(os.Setenv("PACKAGE", "user"))
	vol := gen.Volume("/src:/dst/$PACKAGE")
	src, dst, err := vol.Split()

	assert.Equal("/src", src)
	assert.Equal("/dst/user", dst)
	assert.Nil(err)
}
