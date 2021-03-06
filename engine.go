// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
	"unicode"
)

// Stats tracks the statistics for the analyzed dataset
type Stats struct {
	NumElem  int64
	Mean     float64
	Variance float64
	qk, mk   float64  // variance helper variables
	Median   *medData // running median
	Min      float64
	Max      float64
}

// computeStats determined relevant stats on the input file
func computeStats(r io.Reader, rows *rowRange, cols colSpec) (statsMap, error) {
	sc := bufio.NewScanner(r)
	var rowCount int

	// parse first line, determine number of columns, and add it to stats
	sc.Scan()
	buf := bytes.FieldsFunc(bytes.TrimSpace(sc.Bytes()), unicode.IsSpace)
	if len(buf) == 0 {
		return nil, fmt.Errorf("error parsing input")
	}

	// if the colSpecs are empty we consider all columns
	if len(cols) == 0 {
		for i := 0; i < len(buf); i++ {
			cols[i] = struct{}{}
		}
	}

	// initialize map with column statistics
	stats := make(statsMap)
	for k := range cols {
		if k >= len(buf) {
			return stats, fmt.Errorf("Requested column id %d out of range.", k)
		}
		stats[k] = &Stats{Max: -math.MaxFloat64, Min: math.MaxFloat64, Median: newMedData()}
	}

	if rowCount >= rows.minRow && rowCount <= rows.maxRow {
		if err := updateStats(stats, buf); err != nil {
			return nil, err
		}
	}

	// parse the rest of the file
	for sc.Scan() {
		rowCount++
		if rowCount < rows.minRow || rowCount > rows.maxRow {
			continue
		}
		buf := bytes.FieldsFunc(bytes.TrimSpace(sc.Bytes()), unicode.IsSpace)
		if len(buf) == 0 {
			return nil, fmt.Errorf("error parsing input")
		}
		if err := updateStats(stats, buf); err != nil {
			return nil, err
		}
	}
	finalizeStats(stats)
	return stats, nil
}

// updateStats updates the stats for all columns based on the provided values
func updateStats(stats statsMap, buf [][]byte) error {
	if len(buf) < len(stats) {
		return fmt.Errorf("Insufficient number of columns encountered.")
	}

	for i, s := range stats {
		if err := s.update(buf[i]); err != nil {
			return err
		}
	}
	return nil
}

// finalizeStats is called once all the data has been parsed and thus the total
// number of elements is known
func finalizeStats(stats statsMap) {
	for _, s := range stats {
		s.finalize()
	}
}

// update updates the statistic s based on the byteslice b which has to be
// convertible into a float value
func (s *Stats) update(b []byte) error {

	var d float64
	var err error
	if d, err = strconv.ParseFloat(string(b), 64); err != nil {
		return err
	}
	s.NumElem++

	// update min/max
	if d > s.Max {
		s.Max = d
	}
	if d < s.Min {
		s.Min = d
	}

	s.Median = updateMedian(s.Median, d)
	s.Mean += d

	// update variance
	k := float64(s.NumElem)
	s.qk += (k - 1) * (d - s.mk) * (d - s.mk) / k
	s.mk += (d - s.mk) / k

	return nil
}

// finalize computes the mean and variance for the statistic s based on the
// current number of elements
func (s *Stats) finalize() {
	s.Mean /= float64(s.NumElem)
	if s.NumElem > 1 {
		s.Variance = s.qk / float64(s.NumElem-1)
	}
}
