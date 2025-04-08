// This package jalaali conversion inspired from github.com/yaa110/go-persian-calendar library.
//
// Copyright (c) 2016 Navid Fathollahzade
package jalaali

import (
	"strings"
	"time"
)

// A Month specifies a month of the year starting from Farvardin = 1.
// type Month int

// A Weekday specifies a day of the week starting from Shanbe = 0.
type Weekday int

// List of days in a week.
const (
	Shanbeh Weekday = iota
	Yekshanbeh
	Doshanbeh
	Seshanbeh
	Charshanbeh
	Panjshanbeh
	Jomeh
)

var days = []string{
	"شنبه",
	"یک\u200cشنبه",
	"دوشنبه",
	"سه\u200cشنبه",
	"چهارشنبه",
	"پنج\u200cشنبه",
	"جمعه",
}

var shortDays = []string{
	"ش",
	"ی",
	"د",
	"س",
	"چ",
	"پ",
	"ج",
}

// String returns the Persian name of the day in week.
func (d Weekday) String() string {
	switch {
	case d < 0:
		return days[0]
	case d > 6:
		return days[6]
	default:
		return days[d]
	}
}

// Short returns the Persian short name of the day in week.
func (d Weekday) Short() string {
	switch {
	case d < 0:
		return shortDays[0]
	case d > 6:
		return shortDays[6]
	default:
		return shortDays[d]
	}
}

// Weekday get time.Weekday.
func (d Weekday) Weekday() time.Weekday {
	switch d {
	case Shanbeh:
		return time.Saturday
	case Yekshanbeh:
		return time.Sunday
	case Doshanbeh:
		return time.Monday
	case Seshanbeh:
		return time.Tuesday
	case Charshanbeh:
		return time.Wednesday
	case Panjshanbeh:
		return time.Thursday
	case Jomeh:
		return time.Friday
	}
	return 0
}

// JWeekday get weekday from time.Weekday.
func JWeekday(wd time.Weekday) Weekday {
	switch wd {
	case time.Saturday:
		return Shanbeh
	case time.Sunday:
		return Yekshanbeh
	case time.Monday:
		return Doshanbeh
	case time.Tuesday:
		return Seshanbeh
	case time.Wednesday:
		return Charshanbeh
	case time.Thursday:
		return Panjshanbeh
	case time.Friday:
		return Jomeh
	}
	return 0
}

func daysStr() string {
	return strings.Join(days, "|")
}

func shortDaysStr() string {
	return strings.Join(shortDays, "|")
}
