package exercise

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Dmaxiya/Just-write-the-parser/util"
	"github.com/Knetic/govaluate"
)

func getBracket(status int, bracket string) (int, string) {
	if rand.Intn(3) == 0 {
		return status, ""
	}
	if status == 0 {
		if bracket == "(" {
			return status + 1, "("
		}
		return status, ""
	}
	switch bracket {
	case "(":
		status++
	case ")":
		status--
	}
	return status, bracket
}

func genTestCases(caseNum int) []*util.TestCase {
	var (
		bracketStatus int // 用于标示生成的括号是否合法，生成一个 "(" 则 +1，")" 则 -1，当 bracketStatus == 0 时不能生成 ")"
		bracket       string
		input         bytes.Buffer
		opArr         = []byte{'+', '*'}
		testCases     = make([]*util.TestCase, 0, caseNum)
	)

	rand.Seed(time.Now().Unix())
	for i := 0; i < caseNum; i++ {
		input.Reset()
		bracketStatus, bracket = getBracket(bracketStatus, "(")
		input.WriteString(bracket)
		for j := 0; j < 4; j++ {
			if j != 0 {
				bracketStatus, bracket = getBracket(bracketStatus, "(")
				input.WriteByte(opArr[rand.Intn(2)])
				input.WriteString(bracket)
			}
			bracketStatus, bracket = getBracket(bracketStatus, ")")
			input.WriteString(strconv.FormatInt(rand.Int63()%100, 10))
			input.WriteString(bracket)
		}

		// 将所有括号匹配完全
		for bracketStatus != 0 {
			input.WriteByte(')')
			bracketStatus--
		}

		e, _ := govaluate.NewEvaluableExpression(input.String())
		result, _ := e.Evaluate(nil)
		testCases = append(testCases, &util.TestCase{
			Input:  input.String(),
			Expect: int(result.(float64)),
		})
	}
	return testCases
}

func TestGenTestCases(t *testing.T) {
	for _, tc := range genTestCases(5) {
		fmt.Println(tc.Input, "=", tc.Expect)
	}
	/*
	   output:
	   	(57+85)+75+86 = 303
	   	(85*60)+(42)+63 = 5205
	   	(39*(61+34+49)) = 5616
	   	(28)+(25)*(94)+(63) = 2441
	   	29+(61*(50)*73) = 222679
	*/
}
