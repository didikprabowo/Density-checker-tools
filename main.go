package main

import (
	"fmt"
	"regexp"
	"strings"
)

func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}

	}
	return counts
}
func WordCount(value string) int {
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(value, -1)
	return len(results)
}

func main() {
	strLine := "keyword is a combination of phrases or words"
	for index, element := range wordCount(strLine) {
		var density float32
		I := float32(element)
		C := float32(WordCount(strLine))
		density = I / C * 100
		fmt.Printf("Word: %v => Count : %v => Density %-6.2f\n", index, element, density)
	}
	fmt.Printf("Count Word: %v => Characters : %v \n", WordCount(strLine), len(strLine))
}
