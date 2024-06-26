package aliU

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func Test_SafeMarshalUnmarshal(t *testing.T) {
	RegisterTestingT(t)

	{
		ogStr := "ALIIIIIII"

		safeM, err := SafeMarshal(ogStr)
		Expect(err).To(BeNil())

		var back string
		err = SafeUnmarshal(safeM, &back)
		Expect(err).To(BeNil())
		Expect(ogStr).To(Equal(back))
	}

	{

		type someType struct {
			ID     string
			Title  string
			TS     int64
			Length float64
		}

		ts := time.Now().UnixNano()
		ogStr := someType{
			"1234",
			"ABCD",
			ts,
			3.141592,
		}

		safeM, err := SafeMarshal(ogStr)
		Expect(err).To(BeNil())

		var back someType
		err = SafeUnmarshal(safeM, &back)
		Expect(err).To(BeNil())
		Expect(ogStr).To(Equal(back))
	}

}

func Test_SafeMarshalUnmarshalBytes(t *testing.T) {
	RegisterTestingT(t)

	{
		ogStr := "ALIIIIIII"

		safeM, err := SafeMarshalBytes(ogStr)
		Expect(err).To(BeNil())

		var back string
		err = SafeUnmarshalBytes(safeM, &back)
		Expect(err).To(BeNil())
		Expect(ogStr).To(Equal(back))
	}

	{

		type someType struct {
			ID     string
			Title  string
			TS     int64
			Length float64
		}

		ts := time.Now().UnixNano()
		ogStr := someType{
			"1234",
			"ABCD",
			ts,
			3.141592,
		}

		safeM, err := SafeMarshalBytes(ogStr)
		Expect(err).To(BeNil())

		var back someType
		err = SafeUnmarshalBytes(safeM, &back)
		Expect(err).To(BeNil())
		Expect(ogStr).To(Equal(back))
	}

}
