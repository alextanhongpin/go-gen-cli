package gen_test

import (
	"testing"

	"github.com/alextanhongpin/go-gen-cli"
	"github.com/stretchr/testify/assert"
)

func TestFormatSource(t *testing.T) {
	assert := assert.New(t)

	result, err := gen.FormatSource([]byte(`package main
func main() {
	fmt.Println("hello world")
}`))
	assert.Nil(err)
	assert.Equal(`package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
`, string(result))
}
