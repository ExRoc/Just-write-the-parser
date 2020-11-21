package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveDescent(t *testing.T) {
	a := assert.New(t)

	for _, tc := range genTestCases(1000) {
		a.Equal(tc.Expect, ParsingByASinglePass(tc.Input))
	}
}
