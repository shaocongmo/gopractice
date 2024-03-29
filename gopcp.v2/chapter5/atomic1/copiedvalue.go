package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var countVal atomic.Value
	countVal.Store([]int{1, 3, 5, 7})
	antherStore(countVal)
	fmt.Printf("The count value: %+v\n", countVal.Load())
}

func antherStore(countVal atomic.Value) {
	countVal.Store([]int{2, 4, 6, 8})
}
