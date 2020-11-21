package exercise

import (
	"strconv"
	"strings"
)

const (
	plus     = "+"
	multiply = "*"
)

func BuildSimpleAST(input string) []interface{} {
	ast := []interface{}{plus}

	for _, term := range strings.Split(input, plus) {
		partAST := []interface{}{multiply}
		for _, factor := range strings.Split(term, multiply) {
			factorInt, _ := strconv.ParseInt(factor, 10, 64)
			partAST = append(partAST, int(factorInt))
		}
		ast = append(ast, partAST)
	}

	return ast
}
