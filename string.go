package aliU

// // https://github.com/golang/example/blob/master/stringutil/reverse.go#L21
// func reverse(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(input string) string {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	return output
}

// ChunkString splits a string into chunks of the given size or smaller
func ChunkString(s string, size int) []string {
	var chunks = make([]string, 0)
	for len(s) > size {
		chunks = append(chunks, s[:size])
		s = s[size:]
	}
	if len(s) > 0 {
		chunks = append(chunks, s)
	}
	return chunks
}
