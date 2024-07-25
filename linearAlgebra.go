package aliU

import "fmt"

// https://www.pinecone.io/learn/vector-similarity/

func DotProductDistance(a, b []float64) (float64, error) {
	d, err := DotProductSimilarity(a, b)
	if err != nil {
		return -1, err
	}
	return 1.0 - d, nil
}

func DotProductSimilarity(a, b []float64) (float64, error) {
	if len(a) != len(b) || len(a) == 0 || len(b) == 0 {
		return -1, fmt.Errorf("expected two equally sized arrays")
	}
	d := float64(0)
	for idx := range a {
		if a[idx] > 1.0 || b[idx] > 1.0 || a[idx] < 0.0 || b[idx] < 0 {
			return -1, fmt.Errorf("expected floats between 0 and 1 inclusive")
		}
		d += a[idx] * b[idx]
	}
	return d, nil
}
