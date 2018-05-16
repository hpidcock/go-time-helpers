package timehelpers

import (
	"time"
)

// Unix casts time value to units since UNIX epoch.
func Unix(t time.Time, unit time.Duration) int64 {
	nanoDuration := time.Duration(t.UnixNano()) * time.Nanosecond
	unitDuration := nanoDuration / unit
	return int64(unitDuration)
}

// UnixMillisecond casts time value to milliseconds since UNIX epoch.
func UnixMillisecond(t time.Time) int64 {
	return Unix(t, time.Millisecond)
}

// UnixMicrosecond casts time value to microseconds since UNIX epoch.
func UnixMicrosecond(t time.Time) int64 {
	return Unix(t, time.Microsecond)
}

// UnixNanosecond casts time value to nanoseconds since UNIX epoch.
func UnixNanosecond(t time.Time) int64 {
	return t.UnixNano()
}

// UnixSecond casts time value to seconds since UNIX epoch.
func UnixSecond(t time.Time) int64 {
	return t.Unix()
}
