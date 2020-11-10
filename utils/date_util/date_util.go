package date_util

import "time"

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString return time format
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
