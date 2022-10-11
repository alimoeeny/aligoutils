package aliU

var spinniThings = []string{"|", "/", "-", "\\"}

var spinerI = 0

func Spinner() string {
	spinerI += 1
	return spinniThings[spinerI%len(spinniThings)]
}
