package aliU

import "sort"

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MeanFloat64(a []float64) float64 {
	if len(a) < 1 {
		return 0.0
	}
	agg := 0.0
	for _, f := range a {
		agg += f
	}
	return agg / float64(len(a))
}

func MedianFloat64(input []float64) (median float64) {
	c := SortableFloat64Array(input)
	sort.Sort(c)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	l := len(c)
	if l == 0 {
		return 0.0
	} else if l%2 == 0 {
		return MeanFloat64(c[l/2-1 : l/2+1])
	} else {
		return c[l/2]
	}
}
