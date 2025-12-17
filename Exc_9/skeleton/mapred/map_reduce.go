package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
func (mr *MapReduce) Run(input []string) map[string]int {
	// Map Phase
	mapResults := make(chan []KeyValue, len(input))
	var wg sync.WaitGroup

	for _, line := range input {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			mapResults <- mr.wordCountMapper(l)
		}(line)
	}

	wg.Wait()
	close(mapResults)

	// Shuffle Phase
	shuffled := make(map[string][]int)
	for res := range mapResults {
		for _, kv := range res {
			shuffled[kv.Key] = append(shuffled[kv.Key], kv.Value)
		}
	}

	// Reduce Phase
	reduceResults := make(chan KeyValue, len(shuffled))
	var wgReduce sync.WaitGroup

	for key, values := range shuffled {
		wgReduce.Add(1)
		go func(k string, v []int) {
			defer wgReduce.Done()
			reduceResults <- mr.wordCountReducer(k, v)
		}(key, values)
	}

	wgReduce.Wait()
	close(reduceResults)

	finalResult := make(map[string]int)
	for res := range reduceResults {
		finalResult[res.Key] = res.Value
	}

	return finalResult
}

func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	processedString := reg.ReplaceAllString(strings.ToLower(text), " ")
	words := strings.Fields(processedString)

	var results []KeyValue
	for _, word := range words {
		results = append(results, KeyValue{Key: word, Value: 1})
	}
	return results
}

func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	count := 0
	for _, v := range values {
		count += v
	}
	return KeyValue{Key: key, Value: count}
}
