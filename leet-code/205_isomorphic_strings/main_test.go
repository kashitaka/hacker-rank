package _05_isomorphic_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	assert.True(t, isIsomorphic("egg", "add"))
	assert.False(t, isIsomorphic("foo", "bar"))
	assert.True(t, isIsomorphic("2[[", "?||"))
	assert.True(t, isIsomorphic("p", "q"))
	assert.False(t, isIsomorphic("ab", "aa"))
}
