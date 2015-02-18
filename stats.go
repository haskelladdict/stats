// Copyright 2015 Markus Dittrich
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
	"math"
	"os"
	"strconv"
	"strings"
)

const majorVersion = 0
const minorVersion = 1

// flags for command line parser
var (
	rowRangeStr string // specify a specific range of rows to average over
	wantStats   bool   // request statistics output
	wantHist    bool   // request histogram plotting
	numBins     int    // number of bins for histogram
)

type rowRange struct {
	minRow, maxRow int
}

func init() {
	flag.BoolVar(&wantStats, "s", true, "print statistics")
	flag.StringVar(&rowRangeStr, "r", "", "provide row range of type start:end")
	//	flag.BoolVar(&wantHist, "h", false, "compute and show histogram")
	//	flag.IntVar(&numBins, "b", 100, "number of bins for histogram")
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
		os.Exit(1)
	}

	rows, err := parseRowRange(rowRangeStr)
	if err != nil {
		log.Fatal(err)
	}

	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// compute stats. The actual data is stored in s.median.smaller and
	// s.median.larger in case we need it for computing the histogram
	s, err := computeStats(file, &rows)
	if err != nil {
		log.Fatal(err)
	}

	if wantStats {
		printStats(s)
	}
	/*
		if wantHist {
			h, err := computeHist(s, numBins)
			if err != nil {
				log.Fatal(err)
			}
			printHist(h)
		}
	*/
}

// parseRowRange parses a row range string and returns the corresponding
// rowRange struct.
// NOTE: If the passed in string is empty return a default rowRange that
// corresponds to processing the complete file.l
func parseRowRange(input string) (rowRange, error) {
	rows := rowRange{0, math.MaxInt64}
	if len(input) == 0 {
		return rows, nil
	}
	rowSpecs := strings.Split(input, ":")
	if len(rowSpecs) != 2 {
		return rows, fmt.Errorf("invalid row range specifier: %s", input)
	}
	var err error
	if rowSpecs[0] != "" { // keep the default in case it's missing
		if rows.minRow, err = strconv.Atoi(rowSpecs[0]); err != nil {
			return rows, err
		}
	}
	if rowSpecs[1] != "" { // keep the default in case it's missing
		if rows.maxRow, err = strconv.Atoi(rowSpecs[1]); err != nil {
			return rows, err
		}
	}
	return rows, nil
}

// usage prints basic usage info to stdout
func usage() {
	fmt.Printf("stats v%d.%d  (C) 2014 Markus Dittrich\n", majorVersion,
		minorVersion)
	fmt.Println("usage: stats <options> filename")
	fmt.Printf("\noptions:\n")
	flag.PrintDefaults()
}
