package aliU

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_DotProductDistance(t *testing.T) {
	RegisterFailHandler(FastFail)
	{
		var a []float64
		var b []float64
		d, err := DotProductDistance(a, b)
		Expect(err).NotTo(BeNil())
		Expect(d).To(BeNumerically("==", -1))
	}
	{
		var a = []float64{0.1, 0.2}
		var b = []float64{0.3}
		d, err := DotProductDistance(a, b)
		Expect(err).NotTo(BeNil())
		Expect(d).To(BeNumerically("==", -1))
	}
	{
		var a = []float64{0.1, 0.2}
		var b = []float64{0.3, 0.4}
		d, err := DotProductDistance(a, b)
		Expect(err).To(BeNil())
		Expect(d).To(BeNumerically("==", 0.89))
	}

}
