package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	assert := assert.New(t)
	p := &Product{ID: 4, Name: "ji", Price: 4, SKU: "abc-absd-dfsdf"}

	err := p.Validate()
	assert.NoError(err)
}
