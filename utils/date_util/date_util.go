package date_util

import "time"

const (
	apiDateLayout   = "2006-01-02T15:04:05Z"
	apiDBDateLayout = "2006-01-02 15:04:05"
)

//GetNow returns that time in UTC timezone
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString return time format
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

//GetNowDBString return time format suitable to be entered in to the DB
func GetNowDBFormat() string {
	return GetNow().Format(apiDBDateLayout)
}
