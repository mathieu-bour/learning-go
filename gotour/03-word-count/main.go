package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// Instructions: https://go.dev/tour/moretypes/23

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, w := range strings.Fields(s) {
		_, exists := m[w]

		if exists {
			m[w]++
		} else {
			m[w] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
