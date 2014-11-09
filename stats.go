// Package stats is a simple commandline helper script for calculating basic
// statistics on a data file expected to consist of a single column
// of floating point numbers.
// NOTE: Currently stats will read in all the data to compute the statistics
// and thus require memory on the order of the data set size.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// Stats tracks the statistics for the analyzed dataset
type Stats struct {
	numElem  int64
	mean     float64
	variance float64
	median   *medData // running median
	min      float64
	max      float64
}

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
	fmt.Printf("med  : %e\n", s.median.median)
}

// computeStats determined relevant stats on the input file
func computeStats(r io.Reader) (*Stats, error) {
	s := Stats{max: -math.MaxFloat64, min: math.MaxFloat64, median: newMedData()}
	var mk, qk float64 // helper values for one pass variance computation
	var d float64
	var err error
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if d, err = strconv.ParseFloat(sc.Text(), 64); err != nil {
			return nil, err
		}
		s.numElem++

		// update min/max
		if d > s.max {
			s.max = d
		} else if d < s.min {
			s.min = d
		}

		s.median = updateMedian(s.median, d)

		// update mean
		s.mean += d

		// update variance
		k := float64(s.numElem)
		qk += (k - 1) * (d - mk) * (d - mk) / k
		mk += (d - mk) / k
	}
	s.mean /= float64(s.numElem)
	s.variance = qk / float64(s.numElem-1)

	return &s, nil
}

// usage prints basic usage info to stdout
func usage() {
	fmt.Println("usage: stats <options> filename")
}

/*






*/
