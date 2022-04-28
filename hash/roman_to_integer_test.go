package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	actual := romanToInt("MCMXCIV")
	assert.Equal(t, 1994, actual, "result")
}
