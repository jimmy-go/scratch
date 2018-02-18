package scratch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := &Scratch{}
	assert.NotNil(t, c)
}
