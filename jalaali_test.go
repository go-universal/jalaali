package jalaali_test

import (
	"testing"
	"time"

	"github.com/go-universal/jalaali"
	"github.com/stretchr/testify/assert"
)

func TestJalaali(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.TehranTz())
		expected := "1403-01-15T20:14:00+03:30"
		assert.Equal(t, expected, date.String())
	})

	t.Run("FormatDate", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.TehranTz())
		expected := "1403-01-15 20:14:00"
		assert.Equal(t, expected, date.Format(time.DateTime))
	})

	t.Run("AddDate", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 1, jalaali.TehranTz())
		newDate := date.AddDate(1, 1, 1)
		expected := "1404-02-16T20:14:00+03:30"
		assert.Equal(t, expected, newDate.String())
	})

	t.Run("AddTime", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.TehranTz()).
			AddTime(1, 30, 0, 0)
		expected := "1403-01-15T21:44:00+03:30"
		assert.Equal(t, expected, date.String())
	})

	t.Run("SetTime", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.TehranTz())
		date.SetTime(-1, 30, 10, 10)
		expected := "1403-01-15T20:30:10.000000010+03:30"
		assert.Equal(t, expected, date.Format(time.RFC3339Nano))
	})

	t.Run("BeginningOfMonth", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.TehranTz()).
			BeginningOfMonth()
		expected := "1403-01-01T00:00:00+03:30"
		assert.Equal(t, expected, date.String())
	})

	t.Run("EndOfMonth", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.KabulTz())
		newDate := date.EndOfMonth()
		expected := "1403-01-31T23:59:59.999999999+04:30"
		assert.Equal(t, expected, newDate.Format(time.RFC3339Nano))
	})

	t.Run("IsLeap", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.TehranTz())
		assert.True(t, date.IsLeap())
	})

	t.Run("Unix", func(t *testing.T) {
		date := jalaali.Date(1403, 01, 15, 20, 14, 0, 0, jalaali.TehranTz())
		assert.Equal(t, date.Unix(), jalaali.Unix(date.Unix(), 0).Unix())
	})

	t.Run("Now", func(t *testing.T) {
		now := jalaali.Now()
		assert.Equal(t, time.Now().Year(), now.Time().Year())
	})
}

func TestDateFormat(t *testing.T) {
	tests := []struct {
		layout   string
		expected string
	}{
		{"2006", "1403"},
		{"06", "03"},
		{"January", "شهریور"},
		{"Jan", "شهر"},
		{"01", "06"},
		{"1", "6"},
		{"02", "03"},
		{"_2", " 3"},
		{"2", "3"},
		{"Monday", "شنبه"},
		{"Mon", "ش"},
	}

	date := jalaali.Date(1403, jalaali.Shahrivar, 3, 0, 0, 0, 0, nil)
	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			assert.Equal(t, test.expected, date.Format(test.layout))
		})
	}
}

func TestTimeFormat(t *testing.T) {
	tests := []struct {
		layout   string
		expected string
	}{
		{"15", "14"},
		{"03", "02"},
		{"3", "2"},
		{"04", "05"},
		{"4", "5"},
		{"05", "06"},
		{"5", "6"},
		{".999999999", ".01034567"},
		{".999999", ".010345"},
		{".999", ".01"},
		{".000000000", ".010345670"},
		{".000000", ".010345"},
		{".000", ".010"},
		{"Morning", "ظهر"},
		{"PM", "بعد از ظهر"},
		{"pm", "ب.ظ"},
	}

	date := jalaali.Date(1400, jalaali.Farvardin, 1, 14, 5, 6, 10345670, jalaali.TehranTz())
	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			assert.Equal(t, test.expected, date.Format(test.layout))
		})
	}
}

func TestZoneFormat(t *testing.T) {
	tests := []struct {
		layout   string
		expected string
	}{
		{"MST", "Asia/Tehran"},
		{"Z070000", "+033000"},
		{"Z0700", "+0330"},
		{"Z07:00:00", "+03:30:00"},
		{"Z07:00", "+03:30"},
		{"Z07", "+03"},
		{"-070000", "+033000"},
		{"-0700", "+0330"},
		{"-07:00:00", "+03:30:00"},
		{"-07:00", "+03:30"},
		{"-07", "+03"},
	}

	date := jalaali.Date(1400, jalaali.Farvardin, 1, 0, 0, 0, 0, jalaali.TehranTz())
	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			assert.Equal(t, test.expected, date.Format(test.layout))
		})
	}
}
