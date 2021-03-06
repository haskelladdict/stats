// Copyright 2014 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details

package main

import (
	"container/heap"
	"log"
	"math"
)

// medData holds the data structures needed to compute a running median.
// Currently, the running median is implemented via a min and max heap data
// structure and thus requires storage on the order of the data set size
type medData struct {
	smaller, larger FloatHeap
	val             float64
}

// newMedData initializes the data structure for computing the running median
func newMedData() *medData {
	var m medData
	heap.Init(&m.smaller)
	heap.Init(&m.larger)
	return &m
}

// updateMedian updates the running median using two heaps the each keep
// track of elements smaller and larger than the current median.
func updateMedian(m *medData, v float64) *medData {
	if len(m.smaller) == 0 && len(m.larger) == 0 {
		// insert first element
		heap.Push(&m.smaller, -v)
	} else if len(m.smaller) == 0 {
		// insert second element (first case)
		if v > m.larger[0] {
			heap.Push(&m.smaller, -heap.Pop(&m.larger).(float64))
			heap.Push(&m.larger, v)
		} else {
			heap.Push(&m.smaller, -v)
		}
	} else if len(m.larger) == 0 {
		// insert second element (second case)
		if v < -m.smaller[0] {
			heap.Push(&m.larger, -heap.Pop(&m.smaller).(float64))
			heap.Push(&m.smaller, -v)
		} else {
			heap.Push(&m.larger, v)
		}
	} else {
		// insert third and following elements
		if v < m.val {
			heap.Push(&m.smaller, -v)
		} else if v > m.val {
			heap.Push(&m.larger, v)
		} else {
			if len(m.smaller) <= len(m.larger) {
				heap.Push(&m.smaller, -v)
			} else {
				heap.Push(&m.larger, v)
			}
		}
	}

	// fix up heaps if they differ in length by more than 2
	if len(m.smaller) == len(m.larger)+2 {
		heap.Push(&m.larger, -heap.Pop(&m.smaller).(float64))
	} else if len(m.larger) == len(m.smaller)+2 {
		heap.Push(&m.smaller, -heap.Pop(&m.larger).(float64))
	}

	// compute new median
	if len(m.smaller) == len(m.larger) {
		m.val = 0.5 * (m.larger[0] - m.smaller[0])
	} else if len(m.smaller) > len(m.larger) {
		m.val = -m.smaller[0]
	} else {
		m.val = m.larger[0]
	}

	if math.Abs(float64(len(m.smaller)-len(m.larger))) > 1 {
		log.Panic("median heaps differ by more than 2")
	}

	return m
}

// FloatHeap is a min-heap of float64
type FloatHeap []float64

// implement heap interface for FloatHeap
func (f FloatHeap) Len() int {
	return len(f)
}

func (f FloatHeap) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f FloatHeap) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *FloatHeap) Push(x interface{}) {
	*f = append(*f, x.(float64))
}

func (f *FloatHeap) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}
