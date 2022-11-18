package aliU

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_DotX(t *testing.T) {
	RegisterTestingT(t)

	Expect(DotX([]float64{}, 3.9)).To(BeEquivalentTo([]float64{}))

	Expect(DotX([]float64{1.0, 2.0, -1.0}, 3.9)).To(BeEquivalentTo([]float64{3.9, 7.8, -3.9}))
}

func Test_DotXX(t *testing.T) {
	RegisterTestingT(t)

	Expect(DotXX([][]float64{}, 3.9)).To(BeEquivalentTo([][]float64{}))

	Expect(DotXX([][]float64{{1.0, 2.0, -1.0}}, 3.9)).To(BeEquivalentTo([][]float64{{3.9, 7.8, -3.9}}))
}
