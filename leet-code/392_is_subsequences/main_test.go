package _92_is_subsequences

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	assert.True(t, isSubsequence2("abc", "ahbgdc"))
	assert.False(t, isSubsequence2("acb", "ahbgdc"))
	assert.False(t, isSubsequence2("axc", "ahbgdc"))
	assert.True(t, isSubsequence2("", "ahbgdc"))
	assert.False(t, isSubsequence2("ahbgdc", "a"))
	assert.False(t, isSubsequence2("ahbgdc", ""))
	assert.True(t, isSubsequence2("", ""))
}
