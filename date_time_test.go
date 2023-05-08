package aliU

import (
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func Test_DayStr(t *testing.T) {
	RegisterFailHandler(FastFail)

	{
		x := DayStr()
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
