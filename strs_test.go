// Steve Phillips / elimisteve
// 2017.01.08

package strs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOr(t *testing.T) {
	assert.Equal(t, "", Or())
	assert.Equal(t, "", Or(""))
	assert.Equal(t, "a", Or("", "a"))
	assert.Equal(t, "b", Or("b", ""))
	assert.Equal(t, "c", Or("", "c", ""))
}

func TestContains(t *testing.T) {
	assert.True(t, Contains([]string{"a", "b", "c"}, "a"))
	assert.True(t, Contains([]string{"a", "b", "c"}, "b"))
	assert.True(t, Contains([]string{"a", "b", "c"}, "c"))

	assert.False(t, Contains([]string{"a", "b", "c"}, ""))
	assert.False(t, Contains([]string{"a", "b", "c"}, "d"))
}
