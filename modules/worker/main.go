// Command worker is a tiny job processor used to exercise the CI pipeline.
package main

import "fmt"

// Process doubles every input, standing in for real work.
func Process(in []int) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[i] = v * 2
	}
	return out
}

func main() {
	fmt.Println(Process([]int{1, 2, 3}))
}
