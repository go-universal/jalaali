package jalaali

import "time"

// New create new jalaali instance from time.
// If location is nil then the local time is used.
func New(t time.Time) Jalaali {
	if t.Year() < 1097 {
		return new(jTime)
	} else {
		driver := new(jTime)
		driver.setTime(t)
		return driver
	}

}

// Date create a new jalaali instance from jalaali date.
//
// year, month and day represent a day in Persian calendar.
//
// hour, min minute, sec seconds, nsec nanoseconds offsets represent a moment in time.
//
// loc is a pointer to time.Location, if loc is nil then the local time is used.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Jalaali {
	driver := new(jTime)
	driver.set(year, month, day, hour, min, sec, nsec, loc)
	return driver
}

// Unix create a new jalaali instance from unix timestamp.
//
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
func Unix(sec, nsec int64) Jalaali {
	return New(time.Unix(sec, nsec))
}

// Now create a new jalaali instance from current time.
func Now() Jalaali {
	return New(time.Now())

}

// TehranTz get tehran time zone.
func TehranTz() *time.Location {
	return time.FixedZone("Asia/Tehran", 12600) // UTC + 03:30
}

// KabulTz get kabul time zone.
func KabulTz() *time.Location {
	return time.FixedZone("Asia/Kabul", 16200) // UTC + 04:30
}
