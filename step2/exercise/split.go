package exercise

import (
	"strconv"
	"strings"
)

func ParsingBySplit(input string) int {
	var sum int
	for _, term := range strings.Split(input, "+") {
		prod := 1
		for _, factor := range strings.Split(term, "*") {
			factorInt, _ := strconv.ParseInt(factor, 10, 64)
			prod *= int(factorInt)
		}
		sum += prod
	}
	return sum
}
