package date_util

import "time"

const (
	apiDateLayout   = "02-01-2006T15:04:05Z"
	apiDBDateLayout = "02-01-2006 15:04:05"
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
func GetNowDBString() string {
	return GetNow().Format(apiDBDateLayout)
}
