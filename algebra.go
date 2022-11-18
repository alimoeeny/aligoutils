package aliU

func DotX(a []float64, b float64) []float64 {
	r := []float64{}
	for _, x := range a {
		r = append(r, x*b)
	}
	return r
}

func DotXX(a [][]float64, b float64) [][]float64 {
	r := [][]float64{}
	for _, x := range a {
		r = append(r, DotX(x, b))
	}
	return r
}
