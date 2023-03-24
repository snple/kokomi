package util

import "time"

func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
}

func ParseDate(value string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", value, time.Local)
}

func TimeFormat(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04:05")
}

type MyTime time.Time

func (myT MyTime) MarshalText() (data []byte, err error) {
	t := time.Time(myT)
	data = []byte(TimeFormat(t))
	return
}

func (myT *MyTime) UnmarshalText(text []byte) (err error) {
	t := (*time.Time)(myT)
	*t, err = ParseTime(string(text))
	return
}

func TryConvertTimeZone(before string) string {
	t, err := ParseTime(before)
	if err == nil {
		return t.UTC().Format("2006-01-02T15:04:05Z")
	}

	return before
}
