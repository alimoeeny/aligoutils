package aliU

import (
	"fmt"
	"strings"
	"time"
)

func DayStr(t time.Time) string {
	return fmt.Sprintf("%04d_%02d_%02d", t.Year(), t.Month(), t.Day())
}

// ParseDuration gets a string in the format of 01:03:04 and returns a time.Duration assuming this is hour:minutes:seconds
func ParseDuration(durationStr string) (time.Duration, error) {

	durationFormatted := durationStr

	parts := strings.Split(durationStr, ":")
	if len(parts) == 3 {
		durationFormatted = fmt.Sprintf("%sh%sm%ss", parts[0], parts[1], parts[2])
	} else if len(parts) == 2 {
		durationFormatted = fmt.Sprintf("0h%sm%ss", parts[0], parts[1])
	} else if len(parts) == 1 {
		durationFormatted = fmt.Sprintf("0h0m%ss", parts[0])
	} else {
		return 0, fmt.Errorf("invalid duration format")
	}

	duration, err := time.ParseDuration(durationFormatted)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
