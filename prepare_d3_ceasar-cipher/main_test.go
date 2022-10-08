package main

import (
	"fmt"
	"testing"
)

func Test_caesarCipher(t *testing.T) {
	in := "1X7T4VrCs23k4vv08D6yQ3S19G4rVP188M9ahuxB6j1tMGZs1m10ey7eUj62WV2exLT4C83zl7Q80M"
	fmt.Println(caesarCipher(in, 27))
}
