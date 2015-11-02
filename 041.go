// This program does not work in localhost.
// Paste onto https://go-tour-jp.appspot.com/#41

package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int  {
	wordCounts := map[string]int {}

	words := strings.Fields(s)
	for _, word := range words {
		_, exists := wordCounts[word]
		if exists {
			wordCounts[word] += 1
		} else {
			wordCounts[word] = 1
		}
	}
	return wordCounts
}

func main()  {
	wc.Test(WordCount)
}
