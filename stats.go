// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

// stats is a simple commandline helper script for calculating basic
// statistics on a data file expected to consist of a single column
// of floating point numbers.
// NOTE: Currently stats will read in all the data to compute the statistics
// and thus require memory on the order of the data set size.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		usage()
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	s, err := computeStats(file)
	if err != nil {
		log.Fatal(err)
	}

	// print results
	fmt.Printf("#elem: %d\n", s.numElem)
	fmt.Printf("min  : %e\n", s.min)
	fmt.Printf("max  : %e\n", s.max)
	fmt.Printf("mean : %e\n", s.mean)
	fmt.Printf("var  : %e\n", s.variance)
	fmt.Printf("med  : %e\n", s.median.val)
}

// usage prints basic usage info to stdout
func usage() {
	fmt.Println("usage: stats <options> filename")
}
