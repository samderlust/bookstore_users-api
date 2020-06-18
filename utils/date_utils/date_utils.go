package date_utils

import "time"

const (
	dateLayout = "2019-01-02T15:04:05Z"
)

//GetNowString return string of current time
func GetNowString() string {
	return GetNow().Format(dateLayout)
}

//GetNow get current time
func GetNow() time.Time {
	return time.Now().UTC()
}
