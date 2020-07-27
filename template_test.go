package gen_test

import (
	"fmt"
	"testing"

	"github.com/alextanhongpin/go-gen-cli"
	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	assert := assert.New(t)

	type data struct {
		Name string
	}

	tests := []struct {
		input  string
		snake  string
		camel  string
		kebab  string
		pascal string
	}{
		{"John Doe", "john_doe", "johnDoe", "john-doe", "JohnDoe"},
		{"party-relationship", "party_relationship", "partyRelationship", "party-relationship", "PartyRelationship"},
		{"party-100", "party_100", "party100", "party-100", "Party100"},
		{"user_service", "user_service", "userService", "user-service", "UserService"},
	}
	for _, tt := range tests {
		tpl := []byte(`{{ snakecase .Name }} {{ camelcase .Name }} {{ kebabcase .Name }} {{ pascalcase .Name }}`)
		res, err := gen.ParseTemplate(tpl, data{Name: tt.input})
		assert.Nil(err)
		assert.Equal(fmt.Sprintf("%s %s %s %s", tt.snake, tt.camel, tt.kebab, tt.pascal), string(res))
	}
}
