// This package jalaali conversion inspired from github.com/yaa110/go-persian-calendar library.
//
// Copyright (c) 2016 Navid Fathollahzade
package jalaali

import (
	"slices"
	"strings"
)

// A Month specifies a month of the year starting from Farvardin = 1.
type Month int

// List of months in Persian calendar.
const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

// List of Dari months in Persian calendar.
const (
	Hamal Month = 1 + iota
	Sur
	Jauza
	Saratan
	Asad
	Sonboleh
	Mizan
	Aqrab
	Qos
	Jady
	Dolv
	Hut
)

var months = []string{
	"فروردین",
	"اردیبهشت",
	"خرداد",
	"تیر",
	"مرداد",
	"شهریور",
	"مهر",
	"آبان",
	"آذر",
	"دی",
	"بهمن",
	"اسفند",
}

var shortMonths = []string{
	"فرو",
	"ارد",
	"خرد",
	"تیر",
	"مرد",
	"شهر",
	"مهر",
	"آبا",
	"آذر",
	"دی",
	"بهم",
	"اسف",
}

var dariMonths = []string{
	"حمل",
	"ثور",
	"جوزا",
	"سرطان",
	"اسد",
	"سنبله",
	"میزان",
	"عقرب",
	"قوس",
	"جدی",
	"دلو",
	"حوت",
}

var shortDariMonths = []string{
	"حمل",
	"ثور",
	"جوز",
	"سرط",
	"اسد",
	"سنب",
	"میز",
	"عقر",
	"قوس",
	"جدی",
	"دلو",
	"حوت",
}

// {days, leap_days, days_before_start}
var monthMeta = [12][3]int{
	{31, 31, 0},   // Farvardin
	{31, 31, 31},  // Ordibehesht
	{31, 31, 62},  // Khordad
	{31, 31, 93},  // Tir
	{31, 31, 124}, // Mordad
	{31, 31, 155}, // Shahrivar
	{30, 30, 186}, // Mehr
	{30, 30, 216}, // Aban
	{30, 30, 246}, // Azar
	{30, 30, 276}, // Dey
	{30, 30, 306}, // Bahman
	{29, 30, 336}, // Esfand
}

// String returns the Persian name of the month.
func (m Month) String() string {
	switch {
	case m < 1:
		return months[0]
	case m > 11:
		return months[11]
	default:
		return months[m-1]
	}
}

// String returns the Persian short name of the month.
func (m Month) Short() string {
	switch {
	case m < 1:
		return shortMonths[0]
	case m > 11:
		return shortMonths[11]
	default:
		return shortMonths[m-1]
	}
}

// Dari returns the Dari name of the month.
func (m Month) Dari() string {
	switch {
	case m < 1:
		return dariMonths[0]
	case m > 11:
		return dariMonths[11]
	default:
		return dariMonths[m-1]
	}
}

// Dari returns the Dari short name of the month.
func (m Month) DariShort() string {
	switch {
	case m < 1:
		return shortDariMonths[0]
	case m > 11:
		return shortDariMonths[11]
	default:
		return shortDariMonths[m-1]
	}
}

func monthsStr() string {
	return strings.Join(slices.Concat(months, dariMonths), "|")
}

func shortMonthsStr() string {
	return strings.Join(slices.Concat(shortMonths, shortDariMonths), "|")
}

func parseMonth(values ...string) Month {
	for _, month := range values {
		switch month {
		case "فروردین", "فرو":
			return Farvardin
		case "اردیبهشت", "ارد":
			return Ordibehesht
		case "خرداد", "خرد":
			return Khordad
		case "تیر":
			return Tir
		case "مرداد", "مرد":
			return Mordad
		case "شهریور", "شهر":
			return Shahrivar
		case "مهر":
			return Mehr
		case "آبان", "آبا":
			return Aban
		case "آذر":
			return Azar
		case "دی":
			return Dey
		case "بهمن", "بهم":
			return Bahman
		case "اسفند", "اسف":
			return Esfand
		case "حمل":
			return Hamal
		case "ثور":
			return Sur
		case "جوزا", "جوز":
			return Jauza
		case "سرطان", "سرط":
			return Saratan
		case "اسد":
			return Asad
		case "سنبله", "سنب":
			return Sonboleh
		case "میزان", "میز":
			return Mizan
		case "عقرب", "عقر":
			return Aqrab
		case "قوس":
			return Qos
		case "جدی":
			return Jady
		case "دلو":
			return Dolv
		case "حوت":
			return Hut
		}
	}

	return 0
}
