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
	// Split the string by colon
	parts := strings.Split(durationStr, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid duration format")
	}

	durationFormatted := fmt.Sprintf("%sh%sm%ss", parts[0], parts[1], parts[2])

	duration, err := time.ParseDuration(durationFormatted)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
