package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for _, value := range pow {
		if value%2 == 0 {
			fmt.Printf("%d\n", value)
			continue
		}
	}
}
