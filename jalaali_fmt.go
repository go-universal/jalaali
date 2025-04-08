package jalaali

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func (jt jTime) Format(layout string) string {
	// Quick Format RFC3339 and RFC3339Nano
	if layout == time.RFC3339 || layout == time.RFC3339Nano {
		return jt.formatRFC3339(layout == time.RFC3339Nano)
	}

	// Format layout
	isDari := jt.Location().String() == KabulTz().String()
	return strings.NewReplacer(
		// Year
		"2006", formatYear(jt.year, 4),
		"06", formatYear(jt.year, 2),
		// Hour
		"15", fmt.Sprintf("%02d", jt.hour), // Put hour to render before month 1
		// Month
		"January", formatMonth(jt.month, false, isDari),
		"Jan", formatMonth(jt.month, true, isDari),
		"01", fmt.Sprintf("%02d", jt.month),
		"1", fmt.Sprintf("%d", jt.month),
		// Day
		"02", fmt.Sprintf("%02d", jt.day),
		"_2", fmt.Sprintf("%2d", jt.day),
		"2", fmt.Sprintf("%d", jt.day),
		// Weekday
		"Monday", jt.wday.String(),
		"Mon", jt.wday.Short(),
		// Hour
		"03", fmt.Sprintf("%02d", jt.Hour12()),
		"3", fmt.Sprintf("%d", jt.Hour12()),
		// Minute
		"04", fmt.Sprintf("%02d", jt.min),
		"4", fmt.Sprintf("%d", jt.min),
		// Second
		"05", fmt.Sprintf("%02d", jt.sec),
		"5", fmt.Sprintf("%d", jt.sec),
		// Milliseconds
		".999999999", formatFractional(jt.nsec, 9, true),
		".999999", formatFractional(jt.nsec, 6, true),
		".999", formatFractional(jt.nsec, 3, true),
		".000000000", formatFractional(jt.nsec, 9, false),
		".000000", formatFractional(jt.nsec, 6, false),
		".000", formatFractional(jt.nsec, 3, false),
		// Daytime
		"Morning", jt.DayTime().String(),
		"PM", jt.AmPm().String(),
		"pm", jt.AmPm().Short(),
		// Timezone
		"MST", jt.formatMST(),
		"Z070000", jt.formatOffset("Z070000"),
		"Z0700", jt.formatOffset("Z0700"),
		"Z07:00:00", jt.formatOffset("Z07:00:00"),
		"Z07:00", jt.formatOffset("Z07:00"),
		"Z07", jt.formatOffset("Z07"),
		"-070000", jt.formatOffset("-070000"),
		"-0700", jt.formatOffset("-0700"),
		"-07:00:00", jt.formatOffset("-07:00:00"),
		"-07:00", jt.formatOffset("-07:00"),
		"-07", jt.formatOffset("-07"),
	).Replace(layout)
}

// Helpers
func formatFractional(value, length int, trimmed bool) string {
	// validate and normalize value
	if value <= 0 {
		return ""
	} else if value > 999999999 {
		value = 999999999
	}

	// Format fractional
	str := strconv.Itoa(value)
	str = "." + strings.Repeat("0", 9-len(str)) + str
	str = str[:length+1]

	// Empty detection
	parsed, _ := strconv.ParseFloat(str, 32)
	if parsed == 0 {
		return ""
	}

	// Trailling right zero
	if trimmed {
		str = strings.TrimRight(str, "0")
	}

	return str
}

func formatYear(year, length int) string {
	str := fmt.Sprintf("%4d", year)
	if length == 4 {
		return str
	} else {
		return str[2:]
	}

}

func formatMonth(month Month, isShort, isDari bool) string {
	if isDari && isShort {
		return month.DariShort()
	} else if isDari {
		return month.Dari()
	} else if isShort {
		return month.Short()
	} else {
		return month.String()
	}
}

func (jt jTime) formatOffset(f ...string) string {
	format := "-07:00"
	valids := []string{
		"Z070000", "Z0700",
		"Z07:00:00", "Z07:00",
		"Z07",
		"-070000", "-0700",
		"-07:00:00", "-07:00",
		"-07",
	}
	if len(f) > 0 && slices.Contains(valids, f[0]) {
		format = f[0]
	}

	// Return zero offset
	_, offset := jt.Zone()
	if offset == 0 {
		switch format {
		case "-070000":
			return "+000000"
		case "-0700":
			return "+0000"
		case "-07:00:00":
			return "+00:00:00"
		case "-07:00":
			return "+00:00"
		case "-07":
			return "+00"
		case "Z070000", "Z0700", "Z07:00:00", "Z07:00", "Z07":
			return "Z"
		}
	}

	// Calculate offset
	hour := offset / 3600
	min := (offset % 3600) / 60
	sec := offset % 60
	switch format {
	case "Z070000", "-070000":
		return fmt.Sprintf("%+03d%02d%02d", hour, min, sec)
	case "Z0700", "-0700":
		return fmt.Sprintf("%+03d%02d", hour, min)
	case "Z07:00:00", "-07:00:00":
		return fmt.Sprintf("%+03d:%02d:%02d", hour, min, sec)
	case "Z07:00", "-07:00":
		return fmt.Sprintf("%+03d:%02d", hour, min)
	case "Z07", "-07":
		return fmt.Sprintf("%+03d", hour)
	default:
		return fmt.Sprintf("%+03d:%02d", hour, min)
	}
}

func (jt jTime) formatMST() string {
	zone, _ := jt.Zone()
	if zone == "" || strings.ToLower(zone) == "local" {
		return jt.formatOffset("-0700")
	}
	return zone
}

func (jt jTime) formatRFC3339(isNano bool) string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("%04d", jt.year))
	result.WriteString("-")
	result.WriteString(fmt.Sprintf("%02d", jt.month))
	result.WriteString("-")
	result.WriteString(fmt.Sprintf("%02d", jt.day))
	result.WriteString("T")
	result.WriteString(fmt.Sprintf("%02d", jt.hour))
	result.WriteString(":")
	result.WriteString(fmt.Sprintf("%02d", jt.min))
	result.WriteString(":")
	result.WriteString(fmt.Sprintf("%02d", jt.sec))
	// Nanosecond
	if isNano && jt.nsec > 0 {
		nanosec := fmt.Sprintf(".%09s", strconv.FormatInt(int64(jt.nsec), 10))
		result.WriteString(nanosec)
	}
	result.WriteString(jt.formatOffset("Z07:00"))
	return result.String()
}
