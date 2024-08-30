package aliU

import "math/rand"

func ContainsString(s string, a []string) bool {
	r := false
	for _, i := range a {
		if s == i {
			return true
		}
	}
	return r
}

type SortableStringArray []string

func (a SortableStringArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortableStringArray) Len() int {
	return len(a)
}
func (a SortableStringArray) Less(i, j int) bool {
	return a[i] < a[j]
}

func StringArrayContains(a []string, s string) bool {
	for _, aa := range a {
		if aa == s {
			return true
		}
	}
	return false
}

func StringArraysHaveSameElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, aa := range a {
		if !StringArrayContains(b, aa) {
			return false
		}
	}
	for _, bb := range b {
		if !StringArrayContains(a, bb) {
			return false
		}
	}
	return true
}

type SortableInt64Array []int64

func (a SortableInt64Array) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortableInt64Array) Len() int {
	return len(a)
}
func (a SortableInt64Array) Less(i, j int) bool {
	return a[i] < a[j]
}

type SortableFloat64Array []float64

func (a SortableFloat64Array) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortableFloat64Array) Len() int {
	return len(a)
}
func (a SortableFloat64Array) Less(i, j int) bool {
	return a[i] < a[j]
}

func FliterFloat64Array(a []float64, filter func(f float64) bool) []float64 {
	r := []float64{}
	for _, x := range a {
		if filter(x) {
			r = append(r, x)
		}
	}
	return r
}

func UniqueArray[K comparable](a []K) []K {
	if a == nil {
		return a
	}
	m := map[K]bool{}
	for _, x := range a {
		m[x] = true
	}
	r := make([]K, len(m))
	index := 0
	for k := range m {
		r[index] = k
		index++
	}
	return r
}

func Shuffle[K comparable](a []K) []K {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}
