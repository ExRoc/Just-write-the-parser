package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestCase struct {
	Input  string
	Expect int
}

func Assert(flag bool, err string) {
	if !flag {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func Number(x uint8) int {
	return int(x - '0')
}

func Marshal(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

func MarshalIndent(data interface{}) string {
	b, _ := json.MarshalIndent(data, "", "\t")
	return string(b)
}
