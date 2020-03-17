package main

import (
	"fmt"
	"github.com/jamieabc/go-segment-tree/pkg/segment"
	"math"
)

func main() {
	data := []int{7, 1, 3, -4, 5, -2, -10, 0, 12, 3}
	s := segment.New(data, min, math.MaxInt32)

	fmt.Println("data: ", data)

	rangeStart, rangeEnd := 1, 8
	fmt.Printf("Query min from range %d to %d is %d\n", rangeStart, rangeEnd, s.Query(rangeStart, rangeEnd, 0))

	fmt.Println("update index 6 element from -10 to 7")
	s.Update(6, 7)

	rangeStart, rangeEnd = 2, 9
	fmt.Printf("Query min from range %d to %d is %d\n", rangeStart, rangeEnd, s.Query(rangeStart, rangeEnd, 0))
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
