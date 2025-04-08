package jalaali

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Parse parse jalaali datetime from string with layout.www
// It returns a Jalaali instance and an error if the parsing fails.
func Parse(layout, datetime string) (Jalaali, error) {
	// Skip empty layout
	if strings.TrimSpace(layout) == "" {
		return nil, errors.New("layout cannot be empty")
	}

	// Skip empty datetime
	if strings.TrimSpace(datetime) == "" {
		return nil, errors.New("datetime cannot be empty")
	}

	// Proccess layout
	expression := getLayoutExpression(layout)
	rx, err := regexp.Compile(expression)
	if err != nil {
		return nil, errors.New("invalid layout")
	}

	// Get layour args
	matches := rx.FindStringSubmatch(datetime)
	if matches == nil {
		return nil, errors.New("input does not match layout")
	}

	// Resolve parts
	resultMap := make(map[string]string)
	for i, name := range rx.SubexpNames() {
		if i > 0 && name != "" {
			resultMap[name] = matches[i]
		}
	}

	// Parse year
	var year int
	if v, _ := strconv.Atoi(resultMap["2006"]); v > 0 {
		year = v
	} else if v, _ := strconv.Atoi(resultMap["06"]); v > 0 {
		year = 1400 + v
	}

	// Parse month
	var month int
	if v, _ := strconv.Atoi(resultMap["01"]); v > 0 {
		month = v
	} else if v, _ := strconv.Atoi(resultMap["1"]); v > 0 {
		month = v
	} else if v := parseMonth(resultMap["January"], resultMap["Jan"]); v > 0 {
		month = int(v)
	} else {
		month = 1
	}

	// Parse day
	var day int
	if v, _ := strconv.Atoi(resultMap["02"]); v > 0 {
		day = v
	} else if v, _ := strconv.Atoi(strings.TrimSpace(resultMap["_2"])); v > 0 {
		day = v
	} else if v, _ := strconv.Atoi(resultMap["2"]); v > 0 {
		day = v
	} else {
		day = 1
	}

	// Parse hour
	isPm := parseAmPm(resultMap["PM"], resultMap["pm"]) == Pm
	var hour int
	if v, _ := strconv.Atoi(resultMap["15"]); v > 0 {
		hour = v
	} else if v, _ := strconv.Atoi(resultMap["03"]); v > 0 {
		if isPm {
			v = v + 12
		}
		hour = v
	} else if v, _ := strconv.Atoi(resultMap["3"]); v > 0 {
		if isPm {
			v = v + 12
		}
		hour = v
	}

	// Parse minute
	var minute int
	if v, _ := strconv.Atoi(resultMap["04"]); v > 0 {
		minute = v
	} else if v, _ := strconv.Atoi(resultMap["4"]); v > 0 {
		minute = v
	}

	// Parse second
	var second int
	if v, _ := strconv.Atoi(resultMap["05"]); v > 0 {
		second = v
	} else if v, _ := strconv.Atoi(resultMap["5"]); v > 0 {
		second = v
	}

	// Parse nanoseconds
	nsec := parseNanosec(
		resultMap["999999999"], resultMap["999999"], resultMap["999"],
		resultMap["000000000"], resultMap["000000"], resultMap["000"],
	)

	// Parse timezone
	timezone := parseTimezone(
		resultMap["Z070000"], resultMap["Z0700"],
		resultMap["Z07_00_00"], resultMap["Z07_00"], resultMap["Z07"],
		resultMap["070000"], resultMap["0700"], resultMap["07_00_00"],
		resultMap["07_00"], resultMap["07"],
	)

	// Validate
	if month < 1 || month > 12 {
		return nil, errors.New("invalid jalaali date input")
	}

	days := 0
	if isLeap(year) {
		days = monthMeta[month-1][1]
	} else {
		days = monthMeta[month-1][0]
	}
	if day < 1 || day > days {
		return nil, errors.New("invalid jalaali date input")
	}

	if hour < 0 || hour > 23 {
		return nil, errors.New("invalid jalaali date input")
	}

	if minute < 0 || minute > 59 {
		return nil, errors.New("invalid jalaali date input")
	}

	if second < 0 || second > 59 {
		return nil, errors.New("invalid jalaali date input")
	}

	// Create date
	return Date(
		year, Month(month), day,
		hour, minute, second, nsec,
		timezone), nil
}

// getLayoutExpression get regex pattern for layout
func getLayoutExpression(layout string) string {
	return "^" + strings.NewReplacer(
		// Year
		"2006", `(?P<2006>\d{4})`,
		"06", `(?P<06>\d{2})`,
		// Hour
		"15", `(?P<15>\d{2})`,
		// Month
		"January", `(?P<January>`+monthsStr()+`)`,
		"Jan", `(?P<Jan>`+shortMonthsStr()+`)`,
		"01", `(?P<01>\d{2})`,
		"1", `(?P<1>\d{1,2})`,
		// Day
		"02", `(?P<02>\d{2})`,
		"_2", `(?P<_2>(\s\d)|\d{2})`,
		"2", `(?P<2>\d{1,2})`,
		// Weekday
		"Monday", `(?P<Monday>`+daysStr()+`)`,
		"Mon", `(?P<Mon>`+shortDaysStr()+`)`,
		// Hour
		"03", `(?P<03>\d{2})`,
		"3", `(?P<3>\d{1,2})`,
		// Minute
		"04", `(?P<04>\d{2})`,
		"4", `(?P<4>\d{1,2})`,
		// Second
		"05", `(?P<05>\d{2})`,
		"5", `(?P<5>\d{1,2})`,
		// Milliseconds
		".999999999", `(?P<999999999>\.\d{1,9})?`,
		".999999", `(?P<999999>\.\d{1,6})?`,
		".999", `(?P<999>\.\d{1,3})?`,
		".000000000", `(?P<000000000>\.\d{9})?`,
		".000000", `(?P<000000>\.\d{6})?`,
		".000", `(?P<000>\.\d{3})?`,
		// Daytime
		"Morning", `(?P<Morning>`+daytimeStr()+`)`,
		"PM", `(?P<PM>`+amPmStr()+`)`,
		"pm", `(?P<pm>`+shortAmPmStr()+`)`,
		// Timezone
		"MST", `(?P<MST>([A-Za-z\/]+)|([-+]\d{4}))`,
		"Z070000", `(?P<Z070000>Z|([+-]\d{6}))`,
		"Z0700", `(?P<Z0700>Z|([+-]\d{4}))`,
		"Z07:00:00", `(?P<Z07_00_00>Z|([+-]\d{2}:\d{2}:\d{2}))`,
		"Z07:00", `(?P<Z07_00>Z|([+-]\d{2}\:\d{2}))`,
		"Z07", `(?P<Z07>Z|([+-]\d{2}))`,
		"-070000", `(?P<070000>[-+]\d{6})`,
		"-0700", `(?P<0700>[-+]\d{4})`,
		"-07:00:00", `(?P<07_00_00>[-+]\d{2}\:\d{2}\:\d{2})`,
		"-07:00", `(?P<07_00>[-+]\d{2}\:\d{2})`,
		"-07", `(?P<07>[-+]\d{2})`,
	).Replace(layout) + "$"
}

func parseNanosec(values ...string) int {
	for _, value := range values {
		if value != "" {
			value = strings.TrimPrefix(value, ".")
			value = value + strings.Repeat("0", 9-len(value))
			value = strings.TrimPrefix(value, ".")
			i, _ := strconv.Atoi(value)
			if i != 0 {
				return i
			}
		}
	}
	return 0
}

func parseTimezone(values ...string) *time.Location {
	rx, _ := regexp.Compile(`^([+-])(\d{2}):?(\d{2})?:?(\d{2})?$`)
	for _, value := range values {
		matches := rx.FindStringSubmatch(value)
		if len(matches) != 5 {
			continue
		}

		sign := 1
		if matches[1] == "-" {
			sign = -1
		}
		hour, _ := strconv.Atoi(matches[2])
		min, _ := strconv.Atoi(matches[3])
		sec, _ := strconv.Atoi(matches[4])

		name := matches[1]
		if hour == 0 {
			name += "00"
		} else if hour > 0 {
			name += fmt.Sprintf("%02d", hour)
		}

		if min == 0 {
			name += ":00"
		} else if min > 0 {
			name += fmt.Sprintf(":%02d", min)
		}

		if sec > 0 {
			name += fmt.Sprintf(":%02d", sec)
		}

		return time.FixedZone(
			name,
			sign*((hour*3600)+(min*60)+(sec)),
		)
	}

	return time.UTC
}
