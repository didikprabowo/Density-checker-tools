package main

import (
	"encoding/json"
	// "fmt"
	"github.com/gookit/color"
	"net/http"
	"regexp"
	"strings"
)

type Meta struct {
	Count      int `json:"count"`
	Characters int `json:"characters"`
}

type Density struct {
	Word    string  `json:"word"`
	Count   int     `json:"count"`
	Density float32 `json:"density"`
}

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
func Checker(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Words string `json:"words"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		strLine := payload.Words
		var densities []Density

		for index, element := range wordCount(strLine) {
			var densityes Density
			var density float32
			I := float32(element)
			C := float32(WordCount(strLine))
			density = I / C * 100
			color.Blue.Printf("Word: %v => Count : %v => Density %-6.2f\n", index, element, density)
			densityes.Word = index
			densityes.Count = element
			densityes.Density = density
			densities = append(densities, densityes)
		}
		red := color.FgRed.Render
		color.Cyan.Printf("Count Word: %v => Characters : %v \n", WordCount(strLine), red(len(strLine)))
		meta := Meta{
			Count:      WordCount(strLine),
			Characters: len(strLine),
		}

		results := map[string]interface{}{
			"meta":    meta,
			"results": densities,
		}
		json.NewEncoder(w).Encode(results)
		return
	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)

}

func Reu(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		next.ServeHTTP(w, r)
	}
}
func main() {

	http.HandleFunc("/", Reu(Checker))
	http.ListenAndServe(":8080", nil)
}
