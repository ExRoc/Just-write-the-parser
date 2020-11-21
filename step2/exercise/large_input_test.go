package exercise

import (
	"fmt"
	"testing"
	"time"
)

func TestLargeInputs(t *testing.T) {
	testCases := genTestCases(1000000)

	start := time.Now()
	for _, tc := range testCases {
		ParsingBySplit(tc.Input)
	}
	fmt.Println("Time of ParsingBySplit:", time.Since(start))
	// Time of ParsingBySplit: 495.392285ms

	start = time.Now()
	for _, tc := range testCases {
		ParsingByASinglePass(tc.Input)
	}
	fmt.Println("Time of ParsingByASinglePass:", time.Since(start))
	// Time of ParsingByASinglePass: 984.182014ms
}
