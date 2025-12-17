package main

import (
	"exc9/mapred"
	"fmt"
	"os"
	"strings"
)

// Main function
func main() {
	// todo read file
	content, err := os.ReadFile("res/meditations.txt")
	if err != nil {
		panic(err)
	}
	text := strings.Split(string(content), "\n")

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)
	// todo print your result to stdout
	fmt.Println(results)
}
