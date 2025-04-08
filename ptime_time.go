// This package jalaali conversion inspired from github.com/yaa110/go-persian-calendar library.
//
// Copyright (c) 2016 Navid Fathollahzade
package jalaali

import "strings"

// A AmPm specifies the 12-Hour marker.
type AmPm int

// A DayTime represents a part of the day based on hour.
type DayTime int

// List of 12-Hour markers.
const (
	Am AmPm = 0 + iota
	Pm
)

// List of day times.
const (
	Midnight DayTime = iota
	Dawn
	Morning
	BeforeNoon
	Noon
	AfterNoon
	Evening
	Night
)

var daytimes = []string{
	"نیمه\u200cشب",
	"سحر",
	"صبح",
	"قبل از ظهر",
	"ظهر",
	"بعد از ظهر",
	"عصر",
	"شب",
}

var amPm = []string{
	"قبل از ظهر",
	"بعد از ظهر",
}

var shortAmPm = []string{
	"ق.ظ",
	"ب.ظ",
}

// String returns the Persian name of 12-Hour marker.
func (a AmPm) String() string {
	switch {
	case a < 0:
		return amPm[0]
	case a > 1:
		return amPm[1]
	default:
		return amPm[a]
	}
}

// Short returns the Persian short name of 12-Hour marker.
func (a AmPm) Short() string {
	switch {
	case a < 0:
		return shortAmPm[0]
	case a > 1:
		return shortAmPm[1]
	default:
		return shortAmPm[a]
	}
}

// String returns the Persian name of day time.
func (d DayTime) String() string {
	switch {
	case d < 0:
		return daytimes[0]
	case d > 7:
		return daytimes[7]
	default:
		return daytimes[d]
	}
}

func daytimeStr() string {
	return strings.Join(daytimes, "|")
}

func amPmStr() string {
	return strings.Join(amPm, "|")
}

func shortAmPmStr() string {
	return strings.Join(shortAmPm, "|")
}

func parseAmPm(values ...string) AmPm {
	for _, value := range values {
		if value == "بعد از ظهر" ||
			value == "ب.ظ" {
			return Pm
		}
	}
	return Am
}
