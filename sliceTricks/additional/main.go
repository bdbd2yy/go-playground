package main

import (
	"fmt"
)

func main() {
	// filtering without allocating
	a := []int{-1, -2, 3, 4, 5}
	// the b copies the slice header of a, sharing the same underlying array
	b := a[:0]
	for _, n := range a {
		if n > 0 {
			b = append(b, n)
		}
	}
	fmt.Println(b)

	// reverse
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
	fmt.Println(b)
	for i := 0; i < len(b)/2; i++ {
		opp := len(b) - 1 - i
		b[i], b[opp] = b[opp], b[i]
	}
	fmt.Println(b)

	// batch with minimal allocation
	actions := []int{1, 2, 3, 4, 5, 6, 7}
	batchSize := 3
	batchs := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)
	for batchSize > len(actions) {
		// the right part will compute first and then the left part will be assigned
		// these batchs share the underlying array and if you don't restrict the cap for each batch, the modification on earlier batches may affect the subsequent batches.
		actions, batchs = actions[batchSize:], append(batchs, actions[0:batchSize:batchSize])
	}
	batchs = append(batchs, actions)
	// for len(actions) > 0 {
	// 	if len(actions) < batchSize {
	// 		batchs = append(batchs, actions)
	// 		break
	// 	}
	// 	actions, batchs = actions[batchSize:], append(batchs, actions[:batchSize:batchSize])
	// }
}
