package exercise

import (
	"fmt"
	"testing"

	"github.com/Dmaxiya/Just-write-the-parser/util"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input  string
	expect []interface{}
}

func TestBuildSimpleAST(t *testing.T) {
	a := assert.New(t)

	testCases := []*testCase{
		{
			input: "2*1+2*4*8+6*5+2*4+6*7",
			expect: []interface{}{
				"+",
				[]interface{}{"*", 2, 1},
				[]interface{}{"*", 2, 4, 8},
				[]interface{}{"*", 6, 5},
				[]interface{}{"*", 2, 4},
				[]interface{}{"*", 6, 7},
			},
		},
		{
			input: "5*67*4+126*4+75*4*2*7",
			expect: []interface{}{
				"+",
				[]interface{}{"*", 5, 67, 4},
				[]interface{}{"*", 126, 4},
				[]interface{}{"*", 75, 4, 2, 7},
			},
		},
	}

	for _, tc := range testCases {
		a.Equal(tc.expect, BuildSimpleAST(tc.input))
		fmt.Println(util.Marshal(tc.expect))
	}
	/*
		output:
		["+",["*",2,1],["*",2,4,8],["*",6,5],["*",2,4],["*",6,7]]
		["+",["*",5,67,4],["*",126,4],["*",75,4,2,7]]
	*/
}
