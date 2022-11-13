package main

import "fmt"

func main() {
	fmt.Println(uniqueFunc(1, 2, "x", true, 3, 5.2))
}

func uniqueFunc(args ...interface{}) (int, []interface{}) {
	product := 1
	var unique []interface{}
	for _, arg := range args {
		switch arg.(type) {
		case int:
			product *= arg.(int)
		default:
			unique = append(unique, arg)
		}
	}
	return product, unique
}
