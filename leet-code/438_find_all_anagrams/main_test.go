package _38_find_all_anagrams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAnagrams(t *testing.T)  {
	assert.Equal(t, []int{0,1,2}, findAnagrams("abab", "ab"))
}