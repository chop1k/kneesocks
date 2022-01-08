package logger

import (
	"strconv"
	"time"
)

func getBasicParameters() map[string]string {
	now := time.Now()

	return map[string]string{
		"now_unix": strconv.FormatInt(now.Unix(), 10),
		"now":      now.Format("2006-01-02 15:04:05"),
		"now_date": now.Format("2006-01-02"),
		"now_time": now.Format("15:04:05"),
	}
}
