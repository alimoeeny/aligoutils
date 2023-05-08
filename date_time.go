package aliU

import (
	"fmt"
	"time"
)

func DayStr(t time.Time) string {
	return fmt.Sprintf("%04d_%02d_%02d", t.Year(), t.Month(), t.Day())
}
