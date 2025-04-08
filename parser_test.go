package jalaali_test

import (
	"testing"

	"github.com/go-universal/jalaali"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateParse(t *testing.T) {
	tests := []struct {
		layout   string
		datetime string
	}{
		{"2006", "1403"},
		{"06", "03"},
		{"January", "اسفند"},
		{"Jan", "اسف"},
		{"01", "07"},
		{"1", "7"},
		{"02", "08"},
		{"_2", " 9"},
		{"2", "3"},
	}

	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			jalaaliDate, err := jalaali.Parse(test.layout, test.datetime)
			require.NoError(t, err, "unexpected error during parsing")

			formatted := jalaaliDate.Format(test.layout)
			assert.Equal(t, test.datetime, formatted, "formatted datetime mismatch")
		})
	}
}

func TestTimeParse(t *testing.T) {
	tests := []struct {
		layout   string
		datetime string
	}{
		{"15", "15"},
		{"03", "03"},
		{"3", "9"},
		{"04", "03"},
		{"4", "3"},
		{"05", "09"},
		{"5", "2"},
		{".000", ".010"},
		{".000000", ".012340"},
		{".000000000", ".012345600"},
		{".999", ".01"},
		{".999999", ".01234"},
		{".999999999", ".00123456"},
	}

	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			jalaaliTime, err := jalaali.Parse(test.layout, test.datetime)
			require.NoError(t, err, "unexpected error during parsing")

			formatted := jalaaliTime.Format(test.layout)
			assert.Equal(t, test.datetime, formatted, "formatted datetime mismatch")
		})
	}
}

func TestZoneParse(t *testing.T) {
	tests := []struct {
		layout   string
		datetime string
	}{
		{"MST", "UTC"},
		{"Z070000", "Z"},
		{"Z070000", "+033000"},
		{"Z0700", "Z"},
		{"Z0700", "+0330"},
		{"Z07:00:00", "Z"},
		{"Z07:00:00", "+03:30:00"},
		{"Z07:00", "Z"},
		{"Z07:00", "+03:30"},
		{"Z07", "Z"},
		{"Z07", "+03"},
		{"-070000", "+033000"},
		{"-0700", "+0330"},
		{"-07:00:00", "+03:30:00"},
		{"-07:00", "+03:30"},
		{"-07", "+03"},
	}

	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			jalaaliZone, err := jalaali.Parse(test.layout, test.datetime)
			require.NoError(t, err, "unexpected error during parsing")

			formatted := jalaaliZone.Format(test.layout)
			assert.Equal(t, test.datetime, formatted, "formatted datetime mismatch")
		})
	}
}
