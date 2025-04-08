package jalaali

import (
	"math"
	"time"
)

func (jt jTime) IsZero() bool {
	return jt == jTime{}
}

func (jt jTime) IsLeap() bool {
	return isLeap(jt.year)
}

func (jt jTime) Since(t2 Jalaali) time.Duration {
	return time.Duration(math.Abs(float64(t2.Unix()-jt.Unix()))) * time.Second
}

func (jt jTime) AmPm() AmPm {
	if jt.hour > 12 || (jt.hour == 12 && (jt.min > 0 || jt.sec > 0)) {
		return Pm
	}
	return Am
}

func (jt jTime) Zone() (string, int) {
	return jt.Time().Zone()
}

func (jt jTime) In(loc *time.Location) Jalaali {
	res := jt.clone()
	if loc != nil {
		res.loc = loc
	}
	res.resetWeekday()
	return res
}

func (jt jTime) Add(d time.Duration) Jalaali {
	return New(jt.Time().Add(d))
}

func (jt jTime) AddTime(hour, min, sec, nsec int) Jalaali {
	hours := time.Duration(hour) * time.Hour
	mins := time.Duration(min) * time.Minute
	secs := time.Duration(sec) * time.Second
	nanos := time.Duration(nsec) * time.Nanosecond
	return jt.Add(hours + mins + secs + nanos)
}

func (jt jTime) AddDate(year, month, day int) Jalaali {
	return Date(
		jt.year+year, jt.month+Month(month), jt.day+day,
		jt.hour, jt.min, jt.sec, jt.nsec, jt.loc,
	)
}

func (jt jTime) AddDatetime(year, month, day, hour, min, sec, nsec int) Jalaali {
	return Date(
		jt.year+year, jt.month+Month(month), jt.day+day,
		jt.hour+hour, jt.min+min, jt.sec+sec,
		jt.nsec+nsec, jt.loc,
	)
}

func (jt jTime) Yesterday() Jalaali {
	return jt.AddDate(0, 0, -1)
}

func (jt jTime) Tomorrow() Jalaali {
	return jt.AddDate(0, 0, 1)
}

func (jt jTime) BeginningOfDay() Jalaali {
	res := jt.clone()
	res.SetTime(0, 0, 0, 0)
	return res
}
func (jt jTime) EndOfDay() Jalaali {
	res := jt.clone()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (jt jTime) FirstWeekDay() Jalaali {
	if jt.wday == Shanbeh {
		return jt.clone()
	}
	return jt.AddDate(0, 0, int(Shanbeh-jt.wday))
}

func (jt jTime) LastWeekDay() Jalaali {
	if jt.wday == Jomeh {
		return jt.clone()
	}
	return jt.AddDate(0, 0, int(Jomeh-jt.wday))
}

func (jt jTime) BeginningOfWeek() Jalaali {
	res := jt.FirstWeekDay()
	res.SetTime(0, 0, 0, 0)
	return res
}

func (jt jTime) EndOfWeek() Jalaali {
	res := jt.LastWeekDay()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (jt jTime) FirstMonthDay() Jalaali {
	if jt.day == 1 {
		return jt.clone()
	}
	return Date(
		jt.year, jt.month, 1,
		jt.hour, jt.min, jt.sec,
		jt.nsec, jt.loc,
	)
}

func (jt jTime) LastMonthDay() Jalaali {
	dIndex := 0
	if jt.IsLeap() {
		dIndex = 1
	}

	mIndex := jt.month - 1
	if mIndex < 0 {
		mIndex = 0
	} else if mIndex > 11 {
		mIndex = 11
	}

	lastDay := monthMeta[mIndex][dIndex]
	if lastDay == jt.day {
		return jt.clone()
	}
	return Date(
		jt.year, jt.month, lastDay,
		jt.hour, jt.min, jt.sec,
		jt.nsec, jt.loc,
	)
}

func (jt jTime) BeginningOfMonth() Jalaali {
	res := jt.FirstMonthDay()
	res.SetTime(0, 0, 0, 0)
	return res
}

func (jt jTime) EndOfMonth() Jalaali {
	res := jt.LastMonthDay()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (jt jTime) FirstYearDay() Jalaali {
	if jt.month == Farvardin && jt.day == 1 {
		return jt.clone()
	}
	return Date(
		jt.year, Farvardin, 1,
		jt.hour, jt.min, jt.sec,
		jt.nsec, jt.loc,
	)
}

func (jt jTime) LastYearDay() Jalaali {
	dIndex := 0
	if jt.IsLeap() {
		dIndex = 1
	}
	lastDay := monthMeta[Esfand-1][dIndex]
	if jt.month == Esfand && jt.day == lastDay {
		return jt.clone()
	}
	return Date(
		jt.year, Esfand, lastDay,
		jt.hour, jt.min, jt.sec,
		jt.nsec, jt.loc,
	)
}

func (jt jTime) BeginningOfYear() Jalaali {
	res := jt.FirstYearDay()
	res.SetTime(0, 0, 0, 0)
	return res
}

func (jt jTime) EndOfYear() Jalaali {
	res := jt.LastYearDay()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (jt *jTime) SetYear(year int) {
	jt.year = year
	jt.normalizeDay()
	jt.resetWeekday()
}

func (jt *jTime) SetMonth(month Month) {
	jt.month = month
	jt.normalizeMonth()
	jt.normalizeDay()
	jt.resetWeekday()
}

func (jt *jTime) SetDay(day int) {
	jt.day = day
	jt.normalizeDay()
	jt.resetWeekday()
}

func (jt *jTime) SetHour(hour int) {
	jt.hour = hour
	jt.normalizeHour()
}

func (jt *jTime) SetMinute(min int) {
	jt.min = min
	jt.normalizeMin()
}

func (jt *jTime) SetSecond(sec int) {
	jt.sec = sec
	jt.normalizeSec()
}

func (jt *jTime) SetNanosecond(nsec int) {
	jt.nsec = nsec
	jt.normalizeNano()
}

func (jt *jTime) SetTime(hour, min, sec, nsec int) {
	if hour >= 0 {
		jt.SetHour(hour)
	}
	if min >= 0 {
		jt.SetMinute(min)
	}
	if sec >= 0 {
		jt.SetSecond(sec)
	}
	if nsec >= 0 {
		jt.SetNanosecond(nsec)
	}
}

func (jt *jTime) SetDate(year, month, day int) {
	if year > 0 {
		jt.SetYear(year)
	}
	if month > 0 {
		jt.SetMonth(Month(month))
	}
	if day > 0 {
		jt.SetDay(day)
	}
}

func (jt *jTime) SetDateTime(year, month, day, hour, min, sec, nsec int) {
	jt.SetDate(year, month, day)
	jt.SetTime(hour, min, sec, nsec)
}

func (jt jTime) Year() int {
	return jt.year
}

func (jt jTime) YearDay() int {
	month := jt.month - 1
	if month < 0 {
		month = 0
	} else if month > 11 {
		month = 11
	}
	return monthMeta[month][2] + jt.day
}

func (jt jTime) YearRemainDays() int {
	days := 365
	if jt.IsLeap() {
		days = 366
	}
	return days - jt.YearDay()
}

func (jt jTime) Month() Month {
	return jt.month
}

func (jt jTime) Weekday() Weekday {
	return jt.wday
}

func (jt jTime) MonthWeek() int {
	return int(math.Ceil(float64(jt.day+int(jt.FirstMonthDay().Weekday())) / 7.0))
}

func (jt jTime) YearWeek() int {
	return int(math.Ceil(float64(jt.YearDay()+int(jt.FirstYearDay().Weekday())) / 7.0))
}

func (jt jTime) YearRemainWeeks() int {
	return 52 - jt.YearWeek()
}

func (jt jTime) Day() int {
	return jt.day
}

func (jt jTime) MonthRemainDays() int {
	dIndex := 0
	if jt.IsLeap() {
		dIndex = 1
	}

	mIndex := jt.month - 1
	if mIndex < 0 {
		mIndex = 0
	} else if mIndex > 11 {
		mIndex = 11
	}

	return monthMeta[mIndex][dIndex] - jt.day
}

func (jt jTime) Hour() int {
	return jt.hour
}

func (jt jTime) Hour12() int {
	if jt.hour > 12 {
		return jt.hour - 12
	} else {
		return jt.hour
	}
}

func (jt jTime) Minute() int {
	return jt.min
}

func (jt jTime) Second() int {
	return jt.sec
}

func (jt jTime) Nanosecond() int {
	return jt.nsec
}

func (jt jTime) DayTime() DayTime {
	return DayTime(jt.hour / 3)
}

func (jt jTime) Location() *time.Location {
	return jt.loc
}

func (jt jTime) Date() (int, Month, int) {
	return jt.year, jt.month, jt.day
}

func (jt jTime) Clock() (int, int, int) {
	return jt.hour, jt.min, jt.sec
}

func (jt jTime) Unix() int64 {
	return jt.Time().Unix()
}

func (jt jTime) UnixNano() int64 {
	return jt.Time().UnixNano()
}

func (jt jTime) Time() time.Time {
	// Handle empty date
	if jt.IsZero() {
		return time.Time{}
	}

	var year, month, day int

	// Convert the Shamsi to the corresponding Julian Day Number (JDN)
	jdn := convertShamsiToJDN(jt.year, int(jt.month), jt.day)

	// Convert the JDN to a Gregorian testDate
	if jdn > gregorianReformJulianDay {
		year, month, day = convertJDNToGregorianPostReform(jdn)
	} else {
		year, month, day = convertJDNToGregorianPreReform(jdn)
	}

	// Use the location stored in the Time struct, or default to the local time zone
	loc := jt.loc
	if loc == nil {
		loc = time.Local
	}

	// Return the corresponding time.Time object
	return time.Date(
		year, time.Month(month), day,
		jt.hour, jt.min, jt.sec,
		jt.nsec, loc,
	)
}

func (jt jTime) String() string {
	return jt.Format(time.RFC3339)
}
