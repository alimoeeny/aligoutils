package aliU

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func RandStr100(length int) string {
	length = MinInt(length, 100)
	length = MaxInt(length, 1)
	hash := sha512.New()
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Int63()
	s := strconv.FormatInt(i, 10)
	b := hash.Sum([]byte(s))
	he := hex.EncodeToString(b)
	permu := rand.Perm(len(he))
	var shuf string
	for i := 0; i < len(he); i++ {
		shuf = shuf + string(he[permu[i]])
	}
	return shuf[0:length]
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytesLC = "abcdefghijklmnopqrstuvwxyz"
const letterBytesUC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(length int, includeUpperCase bool, includeDigits bool) string {
	rand.Seed(time.Now().UTC().UnixNano())
	sampleSpace := letterBytesLC
	if includeUpperCase {
		sampleSpace += letterBytesUC
	}
	// so that we alwasys start with a letter not a digit
	firstChar := sampleSpace[rand.Int63()%int64(len(sampleSpace))]
	if includeDigits {
		sampleSpace = sampleSpace + "0123456789"
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = sampleSpace[rand.Int63()%int64(len(sampleSpace))]
	}
	b[0] = firstChar
	return string(b)
}

func FancyRandStr(length int) string {
	length = MinInt(length, 4096)
	length = MaxInt(length, 1)
	hash := sha512.New()
	rand.Seed(time.Now().UTC().UnixNano())
	str := ""
	isNotLongEnough := true
	for isNotLongEnough {
		i := rand.Int63()
		s := strconv.FormatInt(i, 10)
		b := hash.Sum([]byte(s))
		he := hex.EncodeToString(b)
		str += he
		isNotLongEnough = len(str) < length
	}
	permu := rand.Perm(len(str))
	var shuf string
	for i := 0; i < len(str); i++ {
		shuf = shuf + string(str[permu[i]])
	}
	return shuf[0:length]
}
