package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Print("Run example \n")

	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 10)
	copy(worst, scores[0:5])
	worst = append(worst, scores[1])
	fmt.Println(worst)
	fmt.Println(scores)
}
