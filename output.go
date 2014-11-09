// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import "fmt"

// printStats prints the computed statistics to stdout
func printStats(s *Stats) {
	fmt.Printf("#elem: %d\n", s.numElem)
	fmt.Printf("min  : %e\n", s.min)
	fmt.Printf("max  : %e\n", s.max)
	fmt.Printf("mean : %e\n", s.mean)
	fmt.Printf("var  : %e\n", s.variance)
	fmt.Printf("med  : %e\n", s.median.val)
}
