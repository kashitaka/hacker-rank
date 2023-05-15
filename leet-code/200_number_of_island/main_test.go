package _00_number_of_island

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numIslands(t *testing.T)  {
	in := [][]byte{
		{48,49,48},
		{49,48,49},
		{48,49,48},
	}
	assert.Equal(t,4, numIslands(in) )
}