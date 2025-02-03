package aliU

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PasswordWorthyString(t *testing.T) {
	RegisterTestingT(t)

	for i := 1; i < 1001; i++ {
		rs := PasswordWorthyString(i)

		Expect(len(rs)).To(Equal(i))
		Expect(strings.Contains("!@#$%^&*{}<>0123456789", string(rs[0]))).To(BeFalse())                              //ElementOf([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}))
		Expect(strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", string(rs[0]))).To(BeTrue()) //([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}))

		// this is silly I guess but whatever
		for i := 1; i < len(rs); i++ {
			Expect(strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*{}<>0123456789", string(rs[i]))).To(BeTrue())
		}
	}
}
