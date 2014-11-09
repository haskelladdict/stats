// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import (
	"bufio"
	"io"
	"math"
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
		}
		if d < s.min {
			s.min = d
		}

		s.median = updateMedian(s.median, d)
		s.mean += d

		// update variance
		k := float64(s.numElem)
		qk += (k - 1) * (d - mk) * (d - mk) / k
		mk += (d - mk) / k
	}

	if s.numElem == 0 {
		return &s, nil
	}

	s.mean /= float64(s.numElem)
	if s.numElem > 1 {
		s.variance = qk / float64(s.numElem-1)
	}
	return &s, nil
}
