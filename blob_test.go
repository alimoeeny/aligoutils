package aliUtils

import (
	"net/url"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_S3Utils(t *testing.T) {
	RegisterTestingT(t)

	Expect(FiveCharPrefix("")).To(Equal(""))
	Expect(FiveCharPrefix("a")).To(Equal("aaaaa"))
	Expect(FiveCharPrefix("ab")).To(Equal("abaaa"))
	Expect(FiveCharPrefix("abc")).To(Equal("abcaa"))
	Expect(FiveCharPrefix("abcd")).To(Equal("abcda"))
	Expect(FiveCharPrefix("aaaaa")).To(Equal("aaaaa"))
	Expect(FiveCharPrefix("abcdefgh")).To(Equal("abcde"))

	url, _ := url.Parse("https://photos.hemati.ai/abcdefg.jpg")
	path := "abcde/abcdefg.jpg"
	Expect(StoragePathFromURL(url)).To(Equal(path))

	url, _ = url.Parse("https://photos.hemati.ai/z.jpg")
	path = "zzzzz/z.jpg"
	Expect(StoragePathFromURL(url)).To(Equal(path))

	url, _ = url.Parse("https://photos.hemati.ai/zy12.jpg")
	path = "zy12z/zy12.jpg"
	Expect(StoragePathFromURL(url)).To(Equal(path))

	url, _ = url.Parse("https://photos.hemati.ai/123456.ali.png")
	path = "12345/123456.ali.png"
	Expect(StoragePathFromURL(url)).To(Equal(path))
}

func Test_Folderify(t *testing.T) {
	RegisterTestingT(t)

	s := ""
	Expect(Folderify(s)).To(Equal(s))

	s = "ali"
	Expect(Folderify(s)).To(Equal("aliaa/ali"))
	s = "sha224"
	Expect(Folderify(s)).To(Equal("sha22/sha224"))

	s = "sasdfiaksdjhfjahgauhgaiug"
	Expect(Folderify(s)).To(Equal("sasdf/sasdfiaksdjhfjahgauhgaiug"))
	s = "sha224-i-12345"
	Expect(Folderify(s)).To(Equal("12345/sha224-i-12345"))
	s = "sha224-i-cbd28cdbaa3643ce1b184c8df6d11df6b59f4db156e3212589ed1501"
	Expect(Folderify(s)).To(Equal("cbd28/sha224-i-cbd28cdbaa3643ce1b184c8df6d11df6b59f4db156e3212589ed1501"))
	s = "sha224-i-db367ccd89fc39aa6ff9fff0cd9b11e6b5b6a41cff4bbb232bee7c93"
	Expect(Folderify(s)).To(Equal("db367/sha224-i-db367ccd89fc39aa6ff9fff0cd9b11e6b5b6a41cff4bbb232bee7c93"))

	s = "sha224-i-"
	Expect(Folderify(s)).To(Equal("sha22/sha224-i-"))

	// but it should not redo it
	s = "4d58b/sha224-i-4d58b0cee1e38f75f93e2e605d1e6ad44c24a975d0c3b53367de7c86"
	Expect(Folderify(s)).To(Equal(s))
	s = "6b470/sha224-i-6b470eaf441040cd2e3d0301ac7163e3999003f022a2c4f714aaae52"
	Expect(Folderify(s)).To(Equal(s))

	s = "sha224-s01-"
	Expect(Folderify(s)).To(Equal("sha22/sha224-s01-"))
	s = "sha224-s01-12345"
	Expect(Folderify(s)).To(Equal("12345/sha224-s01-12345"))
	s = "sha224-s01-cbd28cdbaa3643ce1b184c8df6d11df6b59f4db156e3212589ed1501"
	Expect(Folderify(s)).To(Equal("cbd28/sha224-s01-cbd28cdbaa3643ce1b184c8df6d11df6b59f4db156e3212589ed1501"))
	s = "sha224-s01-db367ccd89fc39aa6ff9fff0cd9b11e6b5b6a41cff4bbb232bee7c93"
	Expect(Folderify(s)).To(Equal("db367/sha224-s01-db367ccd89fc39aa6ff9fff0cd9b11e6b5b6a41cff4bbb232bee7c93"))

	// but it should not redo it
	s = "00000/sha224-s01-000002f32a7144ea0e82e3e573bcb8e54862bde6d9c0e302c5080c05"
	Expect(Folderify(s)).To(Equal(s))
}

func Test_ContentTypeFromFileName(t *testing.T) {
	RegisterTestingT(t)

	_, err := ContentTypeFromFileName("")
	if err == nil {
		t.Error("Expected error for an empty filename")
	}
	_, err = ContentTypeFromFileName("ali")
	if err == nil {
		t.Error("Expected error for filename with no extension")
	}
	ct, err := ContentTypeFromFileName("ali.jpg")
	if err != nil {
		t.Error("Expected no error for filename with extension jpg")
	}
	if ct != "image/jpeg" {
		t.Error("Expected image/jpeg for filename with extension jpg")
	}
}
