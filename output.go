// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
)

// maximum length of histogram bars for displaying purposes
const maxLength = 70

// printStats prints the computed statistics to stdout
func printStats(s *Stats) {
	fmt.Printf("#elem: %d\n", s.numElem)
	fmt.Printf("min  : %e\n", s.min)
	fmt.Printf("max  : %e\n", s.max)
	fmt.Printf("mean : %e\n", s.mean)
	fmt.Printf("var  : %e\n", s.variance)
	fmt.Printf("med  : %e\n", s.median.val)
}

// printHist prints a simple ASCII art version of a histogram.
// This function is very simplistic for now. Hope to make it smarter in the future.
func printHist(h *hist) {

	// find max number of elements and compute scaling factor
	max := -1
	for _, v := range *h {
		if v.n > max {
			max = v.n
		}
	}
	scale := float64(maxLength) / float64(max)

	// print histogram
	for _, v := range *h {
		length := int(math.Floor(float64(v.n) * scale))
		barString, err := repeatChar(length, "\u25A9")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%4.4e   %s\n", v.left, barString)
		fmt.Printf("%4.4e   %s\n", v.right, barString)
	}
}

// repeatString returns a string with n repeats of the provided string
func repeatChar(n int, s string) (string, error) {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		_, err := buf.WriteString(s)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
