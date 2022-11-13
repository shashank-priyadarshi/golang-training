package main

import "fmt"

type emp struct {
	id   int
	name string
	exp  float32
}

func main() {
	fmt.Println("Sorting by Id:\n", sortById([]emp{{4, "X", 1.5}, {2, "Y", 1.5}, {3, "Z", 2.5}, {1, "Z", 2.5}}))
}

func sortById(e []emp) []emp {
	for i := 0; i < len(e); i++ {
		for j := i + 1; j < len(e); j++ {
			if e[i].id > e[j].id {
				e[i], e[j] = e[j], e[i]
			}
		}
	}
	return e
}
