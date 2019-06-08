package binary_gap

import (
	"fmt"
	"strings"
)

func solution(n int) int {
	binary := strings.Split(strings.Trim(fmt.Sprintf("%b", n), "0"), "1")
	prevLength := 0
	for _, each := range binary {
		if len(each) > prevLength {
			prevLength = len(each)
		}
	}
	return prevLength
}
