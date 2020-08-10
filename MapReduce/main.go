/*
* Counts the number of apparitions of each word in
* multiple files concurrently
 */

package main

import (
	"fmt"
	"strings"
	"time"
)

type dict map[string]int
type mapfn func(string, chan dict)
type redfn func(dict, dict)

// creates the dictionary (word:number of apparitions)
// and sends the data via a channel to the reduce phase
func count(rawdata string, c chan dict) { // mapfn
	m := make(dict)
	words := strings.Fields(rawdata)

	for _, word := range words {
		m[word]++
	}

	c <- m
}

// adds the number of apparitions for each dictionary
// of the processed chunk of data
func add(finalMap dict, m dict) { // redfn
	for key, value := range m {
		finalMap[key] += value
	}
}

// Applies the function to every chunk in a goroutine and sends to a channel
func _map(fn mapfn, chunks []string, c chan dict) {
	for _, chunk := range chunks {
		go fn(chunk, c)
	}
}

// Applies the function to every chunk in a goroutine
func _reduce(fn redfn, total int, c chan dict) dict {
	finalMap := make(dict)

	for worker := 0; worker < total; worker++ {
		m := <-c
		fn(finalMap, m)
	}

	return finalMap
}

func perf(fn func() dict) dict {
	start := time.Now()
	result := fn()
	elapsed := time.Since(start)
	fmt.Println("[Performance] Execution took", elapsed.Milliseconds())

	return result
}

func main() {
	c := make(chan dict)
	data1 := "Donec consectetur lacus non leo iaculis feugiat. Vestibulum quis eleifend felis, vel feugiat massa. Etiam ut purus dapibus, iaculis tellus eget, imperdiet risus. Aliquam vitae nunc ultrices, mattis sapien sed."
	data2 := "In vestibulum vestibulum erat vestibulum mattis. Suspendisse eu volutpat libero, at commodo eros. Etiam faucibus blandit nisi. Fusce consectetur pellentesque laoreet. In nec felis pretium, tempus purus aliquam, tristique nibh. "
	data3 := "Proin volutpat metus enim, eget placerat tortor tempus eget. eu Nullam ut elementum erat, nec fermentum neque. Vivamus id massa est. Mauris elementum, orci eu congue faucibus, risus leo cursus metus."

	chunks := []string{data1, data2, data3}

	result := perf(func() dict {
		_map(count, chunks, c)
		return _reduce(add, len(chunks), c)
	})

	close(c)
	fmt.Println("[Result]", result)
}
