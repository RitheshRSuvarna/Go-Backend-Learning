package common

import "time"

type Time struct {
	value time.Time
}

func NewTime(t time.Time) Time {
	return Time{value: t}
}

func Now() Time {
	return Time{value: time.Now()}
}

func (t Time) Time() time.Time {
	return t.value
}

func (t Time) Format(Layout string) string {
	return t.value.Format(Layout)
}