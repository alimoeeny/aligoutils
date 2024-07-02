package aliU

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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

func SafeMarshall(s interface{}) (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	gz.Write(jsonData)
	gz.Close()
	encoded := base64.StdEncoding.EncodeToString(b.Bytes())
	return encoded, nil
}

func SafeUnmarshal(encoded string, out interface{}) error {
	// Ensure out is a pointer
	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("out parameter must be a pointer")
	}
	compressedData, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}
	r, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return err
	}
	decompressedData, err := ioutil.ReadAll(r)

	err = json.Unmarshal(decompressedData, &out)
	return err
}
