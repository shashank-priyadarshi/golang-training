package main

import (
	"errors"
	"log"
)

// var ErrLengthMismatch = log.New("length mismatch")
var ErrLengthMismatch = errors.New("length mismatch")

func main() {
	m, err := buildMap([]int{1, 2, 3, 1, 2}, []string{"x", "y", "z", "a", "b"})
	if err != nil {
		log.Println("Error building map:", err)
	}
	
	for key, value := range m {
		log.Println(key, value)
	}
	
	m, err = buildMap([]int{1, 2, 3, 1}, []string{"x", "y", "z", "a", "b"})
	if err != nil {
		log.Println("Error building map:", err)
	}
}

func buildMap(values []int, keys []string) (map[int]string, error) {
	m := make(map[int]string)
	if len(values) != len(keys) {
		return nil, ErrLengthMismatch
	}
	for i, key := range keys {
		m[values[i]] = key
	}
	return m, nil
}
