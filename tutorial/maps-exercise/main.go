package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	stringFields := strings.Fields(s)
	fmt.Println(stringFields)
	retMap := make(map[string]int)
	for _, k := range stringFields {
		curr := retMap[k]
		retMap[k] = curr + 1
	}
	return retMap
}

func main() {
	wc.Test(WordCount)
}
