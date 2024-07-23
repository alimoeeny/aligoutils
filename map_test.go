package aliU

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_mergeMaps(t *testing.T) {
	RegisterFailHandler(FastFail)

	{
		a := map[string]string{"a": "A", "b": "B"}
		b := map[string]string{"x": "X", "y": "Y"}
		m := mergeMaps(a, b)
		Expect(m["a"]).To(Equal("A"))
		Expect(m["b"]).To(Equal("B"))
		Expect(m["x"]).To(Equal("X"))
		Expect(m["y"]).To(Equal("Y"))
	}

	{
		a := map[string]string{"a": "A", "b": "B"}
		var b map[string]string
		m := mergeMaps(a, b)
		Expect(m["a"]).To(Equal("A"))
		Expect(m["b"]).To(Equal("B"))
	}

	{
		a := map[string]string{"a": "A", "b": "B"}
		var b map[string]string
		m := mergeMaps(b, a)
		Expect(m["a"]).To(Equal("A"))
		Expect(m["b"]).To(Equal("B"))
	}

	{
		a := map[string]int64{"a": 1721738090648502000, "b": 1721727290648505000}
		b := map[string]int64{"x": 1721738104338683000, "y": 4877411690648528000}
		m := mergeMaps(a, b)
		Expect(m["a"]).To(BeNumerically("==", 1721738090648502000))
		Expect(m["b"]).To(BeNumerically("==", 1721727290648505000))
		Expect(m["x"]).To(BeNumerically("==", 1721738104338683000))
		Expect(m["y"]).To(BeNumerically("==", 4877411690648528000))
	}

}
