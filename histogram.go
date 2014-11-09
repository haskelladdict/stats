// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import (
	"fmt"
	"math"
)

// bin describes an individual histogram bin
type bin struct {
	left, right float64 // left and right boundaries of bin
	n           int     // number of elements in the bin
}

// hist describes a complete histogram.
// NOTE: The list of bins is sorted in ascending order
type hist []bin

// computeHist computes a histogram based on the parsed data
func computeHist(s *Stats, nBins int) (*hist, error) {
	if s.max <= s.min {
		return nil, fmt.Errorf("histogram: the date range is <= 0")
	}

	// create bins
	var h hist
	binWidth := (s.max - s.min) / float64(nBins)
	for i := 0; i < nBins; i++ {
		b := bin{left: s.min + float64(i)*binWidth, right: s.min + float64(i+1)*binWidth}
		h = append(h, b)
	}
	// make sure the last bin ends exactly at s.max
	h[nBins-1].right = s.max

	// fill histogram with data
	for _, v := range s.median.smaller {
		i := int(math.Floor((-v - s.min) / binWidth))
		// special case: map all s.max values back into largest bin since above
		// computation places them in nBins+1 bin.
		if i == nBins {
			i--
		}
		h[i].n++
	}
	for _, v := range s.median.larger {
		i := int(math.Floor((v - s.min) / binWidth))
		// special case: map all s.max values back into largest bin since above
		// computation places them in nBins+1 bin.
		if i == nBins {
			i--
		}
		h[i].n++
	}

	return &h, nil
}
