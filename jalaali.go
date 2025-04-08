// Package gojalaali provides an interface and implementation for manipulating
// Jalaali (Persian) calendar dates and times. It supports standard Go time package
// formats and includes functions for setting and getting various components of
// Jalaali dates and times, as well as converting between Jalaali and Gregorian
// dates. The package also includes utility functions for working with time zones
// specific to Tehran and Kabul.
//
// Conversion inspired from github.com/yaa110/go-persian-calendar library.
package jalaali

import "time"

// Jalaali represents an interface for
// manipulating Jalaali (Persian) calendar dates and times.
// It support standard go time package formats.
type Jalaali interface {
	// IsZero returns true if jalaali is zero time instance.
	IsZero() bool

	// IsLeap returns true if the year of t is a leap year.
	IsLeap() bool

	// Since returns the number of seconds between t and t2.
	Since(t2 Jalaali) time.Duration

	// AmPm returns the 12-Hour marker of instance.
	AmPm() AmPm

	// Zone returns the zone name and its offset
	// in seconds east of UTC of instance.
	Zone() (string, int)

	// In sets the location of jalaali date and returns a new instance.
	// If nil loc passed this method returns same instance.
	In(loc *time.Location) Jalaali

	// Add add duration amount to Jalaali and returns a new instance.
	Add(d time.Duration) Jalaali

	// AddTime add hour, minute, second and nanoseconds
	// to Jalaali and returns a new instance.
	AddTime(hour, min, sec, nsec int) Jalaali

	// AddDate add year, month and day
	// to Jalaali and returns a new instance.
	AddDate(year, month, day int) Jalaali

	// AddDatetime add the year, month, day,
	// hour, minute, second and nanoscond to Jalaali
	// and returns a new instance..
	AddDatetime(year, month, day, hour, min, sec, nsec int) Jalaali

	// Yesterday returns a new instance of Jalaali
	// representing a day before the day of instance.
	Yesterday() Jalaali

	// Tomorrow returns a new instance of Jalaali
	// representing a day after the day of instance.
	Tomorrow() Jalaali

	// BeginningOfDay returns a new instance of Jalaali
	// representing the 00:00:00.000000000 time of today.
	BeginningOfDay() Jalaali

	// EndOfDay returns a new instance of Jalaali
	// representing the 23:59:59.999999999 time of today.
	EndOfDay() Jalaali

	// FirstWeekDay returns a new instance of Jalaali
	// representing the first day of the week of instance.
	FirstWeekDay() Jalaali

	// LastWeekDay returns a new instance of Jalaali
	// representing the last day of the week of instance.
	LastWeekDay() Jalaali

	// BeginningOfWeek returns a new instance of Jalaali
	// representing the first day of the week of instance
	// and time is set to 00:00:00.000000000.
	BeginningOfWeek() Jalaali

	// EndOfWeek returns a new instance of Jalaali
	// representing the last day of the week of instance
	// and time is set to 23:59:59.999999999.
	EndOfWeek() Jalaali

	// FirstMonthDay returns a new instance of Jalaali
	// representing the first day of the month of instance.
	FirstMonthDay() Jalaali

	// LastMonthDay returns a new instance of Jalaali
	// representing the last day of the month of instance.
	LastMonthDay() Jalaali

	// BeginningOfMonth returns a new instance of Jalaali
	// representing the first day of the month of instance
	// and time is set to 00:00:00.000000000.
	BeginningOfMonth() Jalaali

	// EndOfMonth returns a new instance of Jalaali
	// representing the first day of the month of instance
	// and time is set to 23:59:59.999999999.
	EndOfMonth() Jalaali

	// FirstYearDay returns a new instance of Jalaali
	// representing the first day of the year of instance.
	FirstYearDay() Jalaali

	// LastYearDay returns a new instance of Jalaali
	// representing the last day of the year of instance.
	LastYearDay() Jalaali

	// BeginningOfYear returns a new instance of Jalaali
	// representing the first day of the year of instance
	// and time is set to 00:00:00.000000000.
	BeginningOfYear() Jalaali

	// EndOfYear returns a new instance of Jalaali
	// representing the first day of the year of instance
	// and time is set to 23:59:59.999999999.
	EndOfYear() Jalaali

	// SetYear sets the year of the instance.
	SetYear(year int)

	// SetMonth sets the month of the instance.
	SetMonth(month Month)

	// SetDay sets the day of the instance.
	SetDay(day int)

	// SetHour sets the hour offset of the Jalaali time.
	SetHour(hour int)

	// SetMinute sets the minute offset of the Jalaali time.
	SetMinute(min int)

	// SetSecond sets the second offset of the Jalaali time.
	SetSecond(sec int)

	// SetNanosecond sets the nanosecond offset of the Jalaali time.
	SetNanosecond(nsec int)

	// SetTime sets the hour, minute, second and nanosecond of the Jalaali time.
	// Pass -1 to ignore parameter.
	SetTime(hour, min, sec, nsec int)

	// SetDate sets the year, month, and day of the Jalaali date.
	// Pass -1 to ignore parameter.
	SetDate(year, month, day int)

	// SetDateTime sets the year, month, day,
	//  hour, minute, and second of the Jalaali date and time.
	// Pass -1 to ignore parameter.
	SetDateTime(year, month, day, hour, min, sec, nsec int)

	// Year returns the year of t.
	Year() int

	// YearDay returns the day of year of instance.
	YearDay() int

	// YearRemainDays returns the number of remaining days of the year of instance.
	YearRemainDays() int

	// Month returns the month of t in the range [1, 12].
	Month() Month

	// Weekday returns the weekday of instance.
	Weekday() Weekday

	// MonthWeek returns the week of month of instance.
	MonthWeek() int

	// YearWeek returns the week of year of instance.
	YearWeek() int

	// YearRemainWeeks returns the number of remaining weeks of the year of instance.
	YearRemainWeeks() int

	// Day returns the day of month of t.
	Day() int

	// MonthRemainDays returns the number of remaining days of the month of instance.
	MonthRemainDays() int

	// Hour returns the hour of t in the range [0, 23].
	Hour() int

	// Hour12 returns the hour of t in the range [0, 11].
	Hour12() int

	// Minute returns the minute offset of t in the range [0, 59].
	Minute() int

	// Second returns the seconds offset of t in the range [0, 59].
	Second() int

	// Nanosecond returns the nanoseconds offset of t in the range [0, 999999999].
	Nanosecond() int

	// DayTime returns the dayTime of that part of the day.
	// [0,3]   -> Midnight
	// [3,6]   -> Dawn
	// [6,9]   -> Morning
	// [9,12]  -> BeforeNoon
	// [12,15] -> Noon
	// [15,18] -> AfterNoon
	// [18,21] -> Evening
	// [21,24] -> Night
	DayTime() DayTime

	// Location returns a pointer to time.Location of instance.
	Location() *time.Location

	// Date returns the year, month, day of instance.
	Date() (int, Month, int)

	// Clock returns the hour, minute, seconds offsets of instance.
	Clock() (int, int, int)

	// Unix returns the number of seconds since January 1, 1970 UTC.
	Unix() int64

	// UnixNano seturns the number of nanoseconds since January 1, 1970 UTC.
	UnixNano() int64

	// Time converts the Shamsi (Solar Hijri) to Gregorian
	// and returns it as a Go time.Time object.
	Time() time.Time

	// String returns t in RFC3339 format.
	String() string

	// TimeFormat formats in standard time package layout.
	//
	// Year
	// 2006				Four-digit year								"1403"
	// 06				Two-digit year 								"03"
	//
	// Month
	// January			Full month name								"اسفند"
	// Jan				Three-letter abbreviation of the month		"اسف"
	// 01				Two-digit month with a leading 0			"07"
	// 1				One-digit month								"7"
	//
	// Day
	// 02				Two-digit month day with a leading 0		"08"
	// _2				Two-digit month day with a leading space	" 9"
	// 2				One-digit month day							"3"
	//
	// Weekday
	// Monday			Full weekday name							"شنبه"
	// Mon				abbreviation of the weekday					"ش"
	//
	// Hour
	// 15				Two-digit 24 hour format					"15"
	// 03				Two-digit 12 hour format					"03"
	// 3				One-digit 12 hour format					"9"
	//
	// Minute
	// 04				Two-digit minute with leading 0				"03"
	// 4				One-digit minute							"3"
	//
	// Second
	// 05				Two-digit second with leading 0				"09"
	// 5				One-digit second							"2"
	//
	// Milliseconds
	// .000				Millisecond									".120"
	// .000000			Microsecond									".123400"
	// .000000000		Nanosecond									".123456000"
	// .999				Trailing zeros removed millisecond			".12"
	// .999999			Trailing zeros removed microsecond			".1234"
	// .999999999		Trailing zeros removed nanosecond			".123456"
	//
	// Daytime
	// Morning			day time									"صبح"
	// PM				Full 12-Hour marker							"قبل از ظهر"
	// pm				Short 12-Hour marker						"ق.ظ"
	//
	// Timezone
	// MST				Abbreviation of the time zone				"UTC"
	// Z070000			zone offset	Hour, Minute and second			"Z" or "+033000"
	// Z0700			zone offset Hour and Minute					"Z" or "+0330"
	// Z07:00:00		zone offset	Hour, Minute and second			"Z" or "+03:30:00"
	// Z07:00			zone offset Hour and Minute					"Z" or "+03:30"
	// Z07				zone offset Hour							"Z" or "+03"
	// -070000			zone offset	Hour, Minute and second			"+033000"
	// -0700			zone offset Hour and Minute					"+0330"
	// -07:00:00		zone offset Hour, Minute and second			"+03:30:00"
	// -07:00			zone offset Hour and Minute					"+03:30"
	// -07				zone offset Hour							"+03"
	Format(layout string) string
}
