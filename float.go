package aliU

import (
	"fmt"
	"math"
	"strconv"
)

// from https://stackoverflow.com/a/77009270/487855
// prec controls the number of digits (excluding the exponent)
// prec of -1 uses the smallest number of digits
func TruncFloat(f float64, prec int) (float64, error) {
	floatBits := 64

	if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) {
		return 0, fmt.Errorf("bad float val %f", f)
	}

	fTruncStr := strconv.FormatFloat(f, 'f', prec+1, floatBits)
	fTruncStr = fTruncStr[:len(fTruncStr)-1]
	fTrunc, err := strconv.ParseFloat(fTruncStr, floatBits)
	if err != nil {
		return 0, err
	}

	return fTrunc, nil
}
