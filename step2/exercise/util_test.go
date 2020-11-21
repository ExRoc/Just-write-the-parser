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

func genTestCases(caseNum int) []*util.TestCase {
	var (
		input     bytes.Buffer
		opArr     = []byte{'+', '*'}
		testCases = make([]*util.TestCase, 0, caseNum)
	)

	rand.Seed(time.Now().Unix())
	for i := 0; i < caseNum; i++ {
		input.Reset()
		for j := 0; j < 4; j++ {
			if j != 0 {
				input.WriteByte(opArr[rand.Intn(2)])
			}
			input.WriteString(strconv.FormatInt(rand.Int63()%100, 10))
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
	   69*1+16*81 = 1365
	   30+70+4+51 = 155
	   65+27*88*87 = 206777
	   88+5+66+49 = 208
	   71+62+80*14 = 1253
	*/
}
