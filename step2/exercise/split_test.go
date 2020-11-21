package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingBySplit(t *testing.T) {
	a := assert.New(t)

	for _, tc := range genTestCases(1000) {
		a.Equal(tc.Expect, ParsingBySplit(tc.Input))
	}
}
