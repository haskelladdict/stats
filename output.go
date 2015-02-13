// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strings"
)

// maximum length of histogram bars for displaying purposes
const maxLength = 70

// printStats prints the computed statistics to stdout
func printStats(stats []*Stats) {
	outStr := make([]string, 8)
	var spacer string
	for i, s := range stats {
		outStr[0] = fmt.Sprintf("%s%scol %-17d", outStr[0], spacer, i)
		outStr[1] = fmt.Sprintf("%s%s#elem: %-14d", outStr[1], spacer, s.NumElem)
		outStr[2] = fmt.Sprintf("%s%smin  : %-10.8e", outStr[2], spacer, s.Min)
		outStr[3] = fmt.Sprintf("%s%smax  : %-10.8e", outStr[3], spacer, s.Max)
		outStr[4] = fmt.Sprintf("%s%smean : %-10.8e", outStr[4], spacer, s.Mean)
		outStr[5] = fmt.Sprintf("%s%sstd  : %-10.8e", outStr[5], spacer, math.Sqrt(s.Variance))
		outStr[6] = fmt.Sprintf("%s%svar  : %-10.8e", outStr[6], spacer, s.Variance)
		outStr[7] = fmt.Sprintf("%s%smed  : %-10.8e", outStr[7], spacer, s.Median.val)
		spacer = strings.Repeat(" ", 5)
	}
	fmt.Println(strings.Join(outStr, "\n"))
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
