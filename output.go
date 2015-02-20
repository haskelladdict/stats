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
func printStats(stats map[int]*Stats) {
	for i, s := range stats {
		fmt.Printf("col %-17d\n", i)
		fmt.Println()
		fmt.Printf("#elem : %-14d\n", s.NumElem)
		fmt.Printf("mean  : %-10.8g\n", s.Mean)
		fmt.Printf("med   : %-10.8g\n", s.Median.val)
		fmt.Printf("std   : %-10.8g\n", math.Sqrt(s.Variance))
		fmt.Printf("var   : %-10.8g\n", s.Variance)
		fmt.Printf("min   : %-10.8g\n", s.Min)
		fmt.Printf("max   : %-10.8g\n", s.Max)
		fmt.Printf("\n\n")
	}
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
	fmt.Println()
	for _, v := range *h {
		length := int(math.Floor(float64(v.n) * scale))
		barString, err := repeatChar(length, "*")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%4.2e  %s\n", v.left, barString)
		fmt.Printf("%4.2e  %s\n", v.right, barString)
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
