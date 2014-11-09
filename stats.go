// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

// stats is a simple commandline helper script for calculating basic
// statistics on a data file expected to consist of a single column
// of floating point numbers.
// NOTE: Currently stats will read in all the data to compute the statistics
// and thus require memory on the order of the data set size.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const version = 0.1

// flags for command line parser
var (
	wantStats bool // request statistics output
	wantHist  bool // request histogram plotting
	numBins   int  // number of bins for histogram
)

func init() {
	flag.BoolVar(&wantStats, "s", true, "print statistics")
	flag.BoolVar(&wantHist, "h", false, "compute and show histogram")
	flag.IntVar(&numBins, "b", 100, "number of bins for histogram")
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
		os.Exit(1)
	}
	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// compute stats. The actual data is stored in s.median.smaller and
	// s.median.larger in case we need it for computing the histogram
	s, err := computeStats(file)
	if err != nil {
		log.Fatal(err)
	}

	if wantStats {
		printStats(s)
	}

	if wantHist {
		h, err := computeHist(s, numBins)
		if err != nil {
			log.Fatal(err)
		}
		printHist(h)
	}
}

// usage prints basic usage info to stdout
func usage() {
	fmt.Printf("stats v%f  (C) 2014 Markus Dittrich\n", version)
	fmt.Println("usage: stats <options> filename")
	fmt.Printf("\noptions:\n")
	flag.PrintDefaults()
}
