package models

import (
	"fmt"
	"time"
)

/**
* 传入的时，时间戳
 */
func UnixToTime(date int) string {
	fmt.Println(date)

	t := time.Unix(int64(date), 0)

	return t.Format("2006-01-02 15:04:05")
}
