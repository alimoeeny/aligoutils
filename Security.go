package aliU

import "math/rand"

func PasswordWorthyString(length int) string {
	sampleSpace := letterBytesLC
	sampleSpace += letterBytesUC

	// so that we alwasys start with a letter not a digit
	firstChar := sampleSpace[rand.Int63()%int64(len(sampleSpace))]

	sampleSpace += sampleSpace + "0123456789"
	sampleSpace += specialByteUC

	b := make([]byte, length)
	for i := range b {
		b[i] = sampleSpace[rand.Int63()%int64(len(sampleSpace))]
	}
	b[0] = firstChar
	return string(b)
}
