package jalaali

import "time"

type jTime struct {
	year  int
	month Month
	day   int
	hour  int
	min   int
	sec   int
	nsec  int
	loc   *time.Location
	wday  Weekday
}

func (jt *jTime) setTime(t time.Time) {
	var year, month, day int
	jt.nsec = t.Nanosecond()
	jt.sec = t.Second()
	jt.min = t.Minute()
	jt.hour = t.Hour()
	jt.loc = t.Location()
	jt.wday = JWeekday(t.Weekday())

	var jdn int
	gy, gmm, gd := t.Date()
	gm := int(gmm)

	if isAfterGregorianReform(gy, gm, gd) {
		jdn = convertGregorianPostReformToJDN(gy, gm, gd)
	} else {
		jdn = convertGregorianPreReformToJDN(gy, gm, gd)
	}

	year, month, day = convertJDNToShamsi(jdn)

	jt.year = year
	jt.month = Month(month)
	jt.day = day
}

func (jt *jTime) set(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) {
	// helpers
	norm := func(hi, lo, base int) (int, int) {
		if lo < 0 {
			n := (-lo-1)/base + 1
			hi -= n
			lo += n * base
		}
		if lo >= base {
			n := lo / base
			hi += n
			lo -= n * base
		}
		return hi, lo
	}
	normDay := func(hi, lo, base int) (int, int) {
		if lo < 1 {
			n := (-lo-1)/base + 1
			hi -= n
			lo += n * base
		}
		if lo > base {
			n := lo / base
			hi += n
			lo -= n * base
		}
		return hi, lo
	}

	// Get current location if not passed
	if loc == nil {
		loc = time.Local
	}

	// Normalize nsec, sec, min, hour, overflowing into day.
	sec, nsec = norm(sec, nsec, 1e9)
	min, sec = norm(min, sec, 60)
	hour, min = norm(hour, min, 60)
	day, hour = norm(day, hour, 24)

	// Normalize month, overflowing into year.
	m := int(month) - 1
	year, m = norm(year, m, 12)

	if m < 0 {
		m = 0
	} else if m > 11 {
		m = 11
	}

	if isLeap(year) {
		m, day = normDay(m, day, monthMeta[m][1])
	} else {
		m, day = normDay(m, day, monthMeta[m][0])
	}
	year, m = norm(year, m, 12)
	month = Month(m) + 1
	jt.year = year
	jt.month = month
	jt.day = day
	jt.hour = hour
	jt.min = min
	jt.sec = sec
	jt.nsec = nsec
	jt.loc = loc
	jt.resetWeekday()
	jt.normalize()
}

func (jt jTime) clone() *jTime {
	return &jTime{
		year:  jt.year,
		month: jt.month,
		day:   jt.day,
		hour:  jt.hour,
		min:   jt.min,
		sec:   jt.sec,
		nsec:  jt.nsec,
		loc:   jt.loc,
		wday:  jt.wday,
	}
}

func (jt *jTime) normalize() {
	jt.normalizeNano()
	jt.normalizeSec()
	jt.normalizeMin()
	jt.normalizeHour()
	jt.normalizeMonth()
	jt.normalizeDay()
}

func (jt *jTime) normalizeNano() {
	if jt.nsec < 0 {
		jt.nsec = 0
	} else if jt.nsec > 999999999 {
		jt.nsec = 999999999
	}
}

func (jt *jTime) normalizeSec() {
	if jt.sec < 0 {
		jt.sec = 0
	} else if jt.sec > 59 {
		jt.sec = 59
	}
}

func (jt *jTime) normalizeMin() {
	if jt.min < 0 {
		jt.min = 0
	} else if jt.min > 59 {
		jt.min = 59
	}
}

func (jt *jTime) normalizeHour() {
	if jt.hour < 0 {
		jt.hour = 0
	} else if jt.hour > 23 {
		jt.hour = 23
	}
}

func (jt *jTime) normalizeMonth() {
	if jt.month < Farvardin {
		jt.month = Farvardin
	} else if jt.month > Esfand {
		jt.month = Esfand
	}
}

func (jt *jTime) normalizeDay() {
	index := 0
	if jt.IsLeap() {
		index = 1
	}

	mIndex := jt.month - 1
	if mIndex < 0 {
		mIndex = 0
	} else if mIndex > 11 {
		mIndex = 11
	}

	days := monthMeta[mIndex][index]
	if jt.day < 1 {
		jt.day = 1
	} else if jt.day > days {
		jt.day = days
	}
}

func (jt *jTime) resetWeekday() {
	jt.wday = JWeekday(jt.Time().Weekday())
}
