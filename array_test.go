package aliUtils

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_StringArrayContains(t *testing.T) {
	RegisterTestingT(t)

	var a []string
	Expect(StringArrayContains(a, "ali")).To(BeFalse())

	a = []string{"ali"}
	Expect(StringArrayContains(a, "ali")).To(BeTrue())
	Expect(StringArrayContains(a, "mali")).To(BeFalse())
	Expect(StringArrayContains(a, "")).To(BeFalse())
}

func Test_StringArraysHaveSameElements(t *testing.T) {
	RegisterTestingT(t)

	{
		var a []string
		var b []string
		Expect(StringArraysHaveSameElements(a, b)).To(Equal(true))
	}

	{
		var a = []string{}
		var b = []string{}
		Expect(StringArraysHaveSameElements(a, b)).To(Equal(true))
	}

	{
		var a = []string{}
		var b = []string{}
		Expect(StringArraysHaveSameElements(a, b)).To(Equal(true))
	}

	{
		var a = []string{"a"}
		var b = []string{"a"}
		Expect(StringArraysHaveSameElements(a, b)).To(Equal(true))
	}

	{
		var a = []string{"a"}
		var b = []string{"b"}
		Expect(StringArraysHaveSameElements(a, b)).To(Equal(false))
	}

	{
		var a = []string{"a", "b", "c"}
		var b = []string{"b", "c", "a"}
		Expect(StringArraysHaveSameElements(a, b)).To(Equal(true))
	}

	{
		var a = []string{"a", "b", "c", "d"}
		var b = []string{"b", "c", "a"}
		Expect(StringArraysHaveSameElements(a, b)).To(Equal(false))
	}

}

func Test_MeanFloat64(t *testing.T) {
	RegisterTestingT(t)

	{
		var a []float64
		Expect(MeanFloat64(a)).To(Equal(0.0))
	}
	{
		a := []float64{}
		Expect(MeanFloat64(a)).To(Equal(0.0))
	}
	{
		a := []float64{-1, 1}
		Expect(MeanFloat64(a)).To(Equal(0.0))
	}
	{
		a := []float64{0, 1, 2}
		Expect(MeanFloat64(a)).To(Equal(1.0))
	}
	{
		a := []float64{-1, 0, 1, 2, 3}
		Expect(MeanFloat64(a)).To(Equal(1.0))
	}
}

func Test_MedianFloat64(t *testing.T) {
	RegisterTestingT(t)

	{
		var a []float64
		Expect(MedianFloat64(a)).To(Equal(0.0))
	}
	{
		a := []float64{}
		Expect(MedianFloat64(a)).To(Equal(0.0))
	}
	{
		a := []float64{-1, 1}
		Expect(MedianFloat64(a)).To(Equal(0.0))
	}
	{
		a := []float64{0, 1, 2}
		Expect(MedianFloat64(a)).To(Equal(1.0))
	}
	{
		a := []float64{-1, 0, 1, 2, 3}
		Expect(MedianFloat64(a)).To(Equal(1.0))
	}
	{
		a := []float64{-1, 0, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 100000}
		Expect(MedianFloat64(a)).To(Equal(3.0))
	}
}

func Test_filterFloat64Array(t *testing.T) {
	RegisterTestingT(t)

	{
		var a []float64
		Expect(FliterFloat64Array(a, func(x float64) bool { return false })).To(BeEmpty())
	}
	{
		a := []float64{}
		Expect(FliterFloat64Array(a, func(x float64) bool { return false })).To(BeEmpty())
	}
	{
		a := []float64{-1, 1}
		Expect(FliterFloat64Array(a, func(x float64) bool { return false })).To(BeEmpty())
	}
	{
		a := []float64{0, 1, 2}
		Expect(FliterFloat64Array(a, func(x float64) bool { return x > 0 })).To(HaveLen(2))
	}
	{
		a := []float64{-1, 0, 1, 2, 3}
		Expect(FliterFloat64Array(a, func(x float64) bool { return x > 2 })).To(HaveLen(1))
	}
	{
		a := []float64{-1, 0, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 100000}
		Expect(FliterFloat64Array(a, func(x float64) bool { return x < 3 })).To(HaveLen(4))
	}
}

func Test_UniqueArray(t *testing.T) {
	RegisterTestingT(t)

	// nil array
	{
		var a []string
		Expect(a).To(Equal(UniqueArray(a)))
	}
	// empty array
	{
		a := []string{}
		Expect(a).To(Equal(UniqueArray(a)))
	}
	// already unique string array
	{
		a := []string{"a", "b", "c"}
		Expect(a).To(ContainElements(UniqueArray(a)))
		Expect(UniqueArray(a)).To(ContainElements(a))
	}
	// canonical usecase
	{
		a := []string{"a", "a", "c", "b", "b", "b", "c", "c"}
		Expect([]string{"a", "c", "b"}).To(ContainElements(UniqueArray(a)))
		Expect(UniqueArray(a)).To(ContainElements([]string{"a", "c", "b"}))
	}
	// canonical usecase
	{
		a := []int{1, 1, 3, 2, 2, 2, 3, 3}
		Expect([]int{1, 3, 2}).To(ContainElements(UniqueArray(a)))
		Expect(UniqueArray(a)).To(ContainElements([]int{1, 2, 3}))
	}

}
