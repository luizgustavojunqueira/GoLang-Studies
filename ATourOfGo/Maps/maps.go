package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words = strings.Fields(s)

	result := make(map[string]int)

	for _, v := range words {
		valor, ok := result[v]
		if ok {
			result[v] = valor + 1
		} else {
			result[v] = 1
		}
	}

	return result

}

func main() {
	wc.Test(WordCount)
}
