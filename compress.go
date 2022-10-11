package aliU

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io/ioutil"
)

func Zipb64(s string) string {
	if s == "" {
		return ""
	}
	var zBytes bytes.Buffer
	w := zlib.NewWriter(&zBytes)
	w.Write([]byte(s))
	w.Close()
	return base64.StdEncoding.EncodeToString(zBytes.Bytes())
}

func Unzipb64(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	zr, err := zlib.NewReader(bytes.NewReader([]byte(b)))
	if err != nil {
		return "", err
	}
	uzBytes, err := ioutil.ReadAll(zr)
	return string(uzBytes), err
}
