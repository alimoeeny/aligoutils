package aliU

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

func padRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func Folderify(filename string) string {
	if filename == "" {
		return ""
	}

	sha224SignatureIndex := strings.Index(filename, "sha224-")
	if sha224SignatureIndex == 6 {
		return filename
	}

	if sha224SignatureIndex == -1 {
		return FiveCharPrefix(filename) + "/" + filename
	}

	signatureEndIndex := sha224SignatureIndex + 8 + strings.Index(filename[sha224SignatureIndex+7:], "-")
	imageprefix := filename[sha224SignatureIndex:signatureEndIndex] // "sha224-s01-"

	fullPath := filename
	if len(fullPath) > len(imageprefix)+4 && (strings.Index(fullPath, imageprefix) == 0) {
		fullPath = fullPath[len(imageprefix):len(imageprefix)+5] + "/" + fullPath
	} else {
		fullPath = FiveCharPrefix(fullPath) + "/" + fullPath
	}
	return fullPath
}

func StoragePathFromURL(url *url.URL) (string, error) {
	fileName := url.Path
	fileName = strings.Trim(fileName, "/")
	return Folderify(fileName), nil //+ "/" + fileName, nil
}

func FiveCharPrefix(s string) string {
	ext := filepath.Ext(s)
	s = strings.TrimRight(s, ext)
	if len(s) < 1 {
		return ""
	}
	if len(s) < 5 {
		return padRight(s, string(s[0]), 5)
	}
	return s[:5]
}

func ContentTypeFromFileName(filename string) (string, error) {
	ext := filename[strings.LastIndex(filename, ".")+1:]
	switch ext {
	case "html":
		return "text/html", nil
	case "css":
		return "text/css", nil
	case "js":
		return "application/javascript", nil
	case "jpg", "jpeg":
		return "image/jpeg", nil
	case "png", "gif":
		return "image/" + ext, nil
	case "mp4", "webm":
		return "video/" + ext, nil
	case "mp3":
		return "audio/mpeg", nil
	case "ogg":
		return "audio/ogg", nil
	case "pdf":
		return "application/pdf", nil
	case "svs":
		return "application/octet-stream", nil
	default:
		return "", fmt.Errorf("unsupported file type: %s", ext)
	}
}

func UniqueIDForBlob(blob []byte) string {
	shaSum := fmt.Sprintf("%x", sha256.Sum224(blob))
	return fmt.Sprintf("sha224-%s.mp3", shaSum)
}
