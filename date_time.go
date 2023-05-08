package aliU

import (
	"fmt"
	"time"
)

func DayStr() string {
	return fmt.Sprintf("%04d_%02d_%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
}
