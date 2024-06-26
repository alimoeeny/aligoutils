package aliU

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ChunkString(t *testing.T) {
	RegisterFailHandler(FastFail)

	{
		chunks := ChunkString("", 9)
		Expect(chunks).To(HaveLen(0))
	}

	{
		chunks := ChunkString("ali", 9)
		Expect(chunks).To(HaveLen(1))
	}

	{
		chunks := ChunkString("ali   ali", 9)
		Expect(chunks).To(HaveLen(1))
	}

	{
		chunks := ChunkString("ali   ali ali ali", 9)
		Expect(chunks).To(HaveLen(2))
	}

}
