package main

import (
	"fmt"
	"hex/internal/adapters/core/arithmetic"
)

func main() {

	// ports
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 5)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
