package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var sum, sumSq float64
	n := 0
	for _, val := range os.Args[1:] {
		fval, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
		}
		sum += fval
		sumSq += fval * fval
		n++
	}
	fmt.Printf("Std: %g\n", math.Sqrt(sumSq/float64(n)-(sum*sum)/float64(n*n)))
}
