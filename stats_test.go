// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import (
	"bytes"
	"math"
	"testing"
)

// check no data case
func Test_statsEmpty(t *testing.T) {

	data := ""
	testBuffer := bytes.NewBufferString(data)
	s, err := computeStats(testBuffer)
	if err != nil {
		t.Error(err)
	}
	if s.mean != 0.0 {
		t.Errorf("mean incorrect: expected 0.0 got %15.15f\n", s.mean)
	}
	if s.variance != 0.0 {
		t.Errorf("variance incorrect: expected 0.0 got %15.15f\n", s.variance)
	}
	if s.median.val != 0.0 {
		t.Errorf("median incorrect: expected 0.0 got %15.15f\n", s.median.val)
	}
	if s.max != -math.MaxFloat64 {
		t.Errorf("max incorrect: expected 0.0 got %15.15f\n", s.max)
	}
	if s.min != math.MaxFloat64 {
		t.Errorf("min incorrect: expected 0.0 got %15.15f\n", s.min)
	}
}

// check single data case
func Test_statsSingle(t *testing.T) {

	data := "1.0\n"
	testBuffer := bytes.NewBufferString(data)
	s, err := computeStats(testBuffer)
	if err != nil {
		t.Error(err)
	}
	if s.mean != 1.0 {
		t.Errorf("mean incorrect: expected 1.0 got %15.15f\n", s.mean)
	}
	if s.variance != 0.0 {
		t.Errorf("variance incorrect: expected 0.0 got %15.15f\n", s.variance)
	}
	if s.median.val != 1.0 {
		t.Errorf("median incorrect: expected 1.0 got %15.15f\n", s.median.val)
	}
	if s.max != 1.0 {
		t.Errorf("max incorrect: expected 1.0 got %15.15f\n", s.max)
	}
	if s.min != 1.0 {
		t.Errorf("min incorrect: expected 1.0 got %15.15f\n", s.min)
	}
}

// check two data case
func Test_statsDouble(t *testing.T) {

	data := "1.0\n2.0\n"
	testBuffer := bytes.NewBufferString(data)
	s, err := computeStats(testBuffer)
	if err != nil {
		t.Error(err)
	}
	if s.mean != 1.5 {
		t.Errorf("mean incorrect: expected 1.5 got %15.15f\n", s.mean)
	}
	if s.variance != 0.5 {
		t.Errorf("variance incorrect: expected 0.5 got %15.15f\n", s.variance)
	}
	if s.median.val != 1.5 {
		t.Errorf("median incorrect: expected 1.5 got %15.15f\n", s.median.val)
	}
	if s.max != 2.0 {
		t.Errorf("max incorrect: expected 2.0 got %15.15f\n", s.max)
	}
	if s.min != 1.0 {
		t.Errorf("min incorrect: expected 1.0 got %15.15f\n", s.min)
	}
}

// check three data case
func Test_statsTriple(t *testing.T) {

	data := "1.0\n2.0\n12.0\n"
	testBuffer := bytes.NewBufferString(data)
	s, err := computeStats(testBuffer)
	if err != nil {
		t.Error(err)
	}
	if s.mean != 5.0 {
		t.Errorf("mean incorrect: expected 5.0 got %15.15f\n", s.mean)
	}
	if s.variance != 37.0 {
		t.Errorf("variance incorrect: expected 37.0 got %15.15f\n", s.variance)
	}
	if s.median.val != 2.0 {
		t.Errorf("median incorrect: expected 2.0 got %15.15f\n", s.median.val)
	}
	if s.max != 12.0 {
		t.Errorf("max incorrect: expected 12.0 got %15.15f\n", s.max)
	}
	if s.min != 1.0 {
		t.Errorf("min incorrect: expected 1.0 got %15.15f\n", s.min)
	}
}

// check multi data case
func Test_statsMulti(t *testing.T) {

	data := "1.0\n2.0\n12\n22\n17.0\n4\n8\n77.0\n13.0\n7.0\n"
	testBuffer := bytes.NewBufferString(data)
	s, err := computeStats(testBuffer)
	if err != nil {
		t.Error(err)
	}
	if s.mean != 16.3 {
		t.Errorf("mean incorrect: expected 16.3 got %15.15f\n", s.mean)
	}
	if !floatEqual(s.variance, 499.12222222222) {
		t.Errorf("variance incorrect: expected 499.12222222 got %15.15f\n", s.variance)
	}
	if s.median.val != 10.0 {
		t.Errorf("median incorrect: expected 10.0 got %15.15f\n", s.median.val)
	}
	if s.max != 77.0 {
		t.Errorf("max incorrect: expected 77.0 got %15.15f\n", s.max)
	}
	if s.min != 1.0 {
		t.Errorf("min incorrect: expected 1.0 got %15.15f\n", s.min)
	}
}

// floatEqual compares two float numbers for equality
// NOTE: the floating point comparison is based on an epsilon
//       which was chosen empirically so it's not rigorous
func floatEqual(a1, a2 float64) bool {
	epsilon := 1e-13
	if math.Abs(a2-a1) > epsilon*math.Abs(a1) {
		return false
	}
	return true
}
