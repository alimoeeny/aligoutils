package aliU

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func Test_DayStr(t *testing.T) {
	RegisterFailHandler(FastFail)

	{
		x := DayStr(time.Now())
		y := strings.Split(x, "_")
		fmt.Println(x)
		Expect(len(y)).To(Equal(3))
		Expect(len(y[0])).To(Equal(4))
		Expect(len(y[1])).To(Equal(2))
		Expect(len(y[2])).To(Equal(2))
		Expect(y[0]).To(MatchRegexp(`\d{4}`))
		Expect(y[1]).To(MatchRegexp(`\d{2}`))
		Expect(y[2]).To(MatchRegexp(`\d{2}`))
		Expect(y[0]).To(Equal(fmt.Sprintf("%04d", time.Now().Year())))
		Expect(y[1]).To(Equal(fmt.Sprintf("%02d", time.Now().Month())))
		Expect(y[2]).To(Equal(fmt.Sprintf("%02d", time.Now().Day())))
	}
}

func Test_ParseDuration(t *testing.T) {
	RegisterFailHandler(FastFail)

	{
		s := "3:21:34"
		v, err := ParseDuration(s)
		Expect(err).To(BeNil())
		Expect(math.Round(v.Hours())).To(BeNumerically("==", 3.0))
		Expect(math.Round(v.Minutes())).To(BeNumerically("==", 202))
		Expect(math.Round(v.Seconds())).To(BeNumerically("==", 12094))
	}
	{
		s := "18:45"
		v, err := ParseDuration(s)
		Expect(err).To(BeNil())
		Expect(math.Round(v.Hours())).To(BeNumerically("==", 0))
		Expect(math.Round(v.Minutes())).To(BeNumerically("==", 19))
		Expect(math.Round(v.Seconds())).To(BeNumerically("==", 1125))
	}

}
