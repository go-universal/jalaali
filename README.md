# Jalaali Library

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/jalaali?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/jalaali.svg)](https://pkg.go.dev/github.com/go-universal/jalaali)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/jalaali/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/jalaali)](https://goreportcard.com/report/github.com/go-universal/jalaali)
![Contributors](https://img.shields.io/github/contributors/go-universal/jalaali)
![Issues](https://img.shields.io/github/issues/go-universal/jalaali)

The Jalaali library provides an interface and implementation for manipulating Jalaali (Persian) calendar dates and times. It supports standard Go time package formats and includes functions for setting and getting various components of Jalaali dates and times, as well as converting between Jalaali and Gregorian dates. The library also includes utility functions for working with time zones specific to Tehran and Kabul.

## Features

- Full support for Jalaali (Persian) calendar.
- Conversion between Jalaali and Gregorian dates.
- Utility functions for date and time manipulation.
- Support for custom time zones (e.g., Tehran and Kabul).
- Parsing and formatting Jalaali dates using Go's time package layout.

## Installation

To install the library, use the following command:

```bash
go get github.com/go-universal/jalaali
```

## API Documentation

### Constructors

#### New

```go
New(t time.Time) Jalaali
```

Creates a new Jalaali instance from a `time.Time` object. If the year is less than 1097, it returns a zero Jalaali instance.

```go
jalaaliDate := jalaali.New(time.Now())
fmt.Println(jalaaliDate.String()) // Outputs the current Jalaali date in RFC3339 format.
```

#### Date

```go
Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Jalaali
```

Creates a new Jalaali instance from the specified Jalaali date and time components.

```go
jalaaliDate := jalaali.Date(1403, jalaali.Farvardin, 15, 20, 14, 0, 0, jalaali.TehranTz())
fmt.Println(jalaaliDate.String()) // Outputs: "1403-01-15T20:14:00+03:30"
```

#### Unix

```go
Unix(sec, nsec int64) Jalaali
```

Creates a new Jalaali instance from a Unix timestamp.

```go
jalaaliDate := jalaali.Unix(1672531200, 0)
fmt.Println(jalaaliDate.String()) // Outputs the Jalaali date corresponding to the Unix timestamp.
```

#### Now

```go
Now() Jalaali
```

Creates a new Jalaali instance representing the current time.

```go
jalaaliDate := jalaali.Now()
fmt.Println(jalaaliDate.String()) // Outputs the current Jalaali date and time.
```

#### TehranTz

```go
TehranTz() *time.Location
```

Returns the Tehran time zone (UTC+03:30).

#### KabulTz

```go
KabulTz() *time.Location
```

Returns the Kabul time zone (UTC+04:30).

### Date and Time Manipulation

#### `Add(d time.Duration) Jalaali`

Adds a duration to the Jalaali instance and returns a new instance.

#### `AddDate(year, month, day int) Jalaali`

Adds the specified number of years, months, and days to the Jalaali instance.

#### `AddTime(hour, min, sec, nsec int) Jalaali`

Adds the specified time components to the Jalaali instance.

#### `Yesterday() Jalaali`

Returns a new instance representing the previous day.

#### `Tomorrow() Jalaali`

Returns a new instance representing the next day.

#### `BeginningOfDay() Jalaali`

Returns a new instance representing the start of the day (00:00:00).

#### `EndOfDay() Jalaali`

Returns a new instance representing the end of the day (23:59:59.999999999).

#### `BeginningOfMonth() Jalaali`

Returns a new instance representing the first day of the month.

#### `EndOfMonth() Jalaali`

Returns a new instance representing the last day of the month.

#### `BeginningOfYear() Jalaali`

Returns a new instance representing the first day of the year.

#### `EndOfYear() Jalaali`

Returns a new instance representing the last day of the year.

### Date and Time Components

#### `Year() int`

Returns the year of the Jalaali instance.

#### `Month() Month`

Returns the month of the Jalaali instance.

#### `Day() int`

Returns the day of the Jalaali instance.

#### `Hour() int`

Returns the hour of the Jalaali instance.

#### `Minute() int`

Returns the minute of the Jalaali instance.

#### `Second() int`

Returns the second of the Jalaali instance.

#### `Nanosecond() int`

Returns the nanosecond of the Jalaali instance.

#### `Weekday() Weekday`

Returns the weekday of the Jalaali instance.

#### `IsLeap() bool`

Returns whether the year of the Jalaali instance is a leap year.

### Formatting and Parsing

#### `Format(layout string) string`

Formats the Jalaali date using Go's time package layout.

```go
jalaaliDate := jalaali.Date(1403, jalaali.Farvardin, 15, 20, 14, 0, 0, jalaali.TehranTz())
formatted := jalaaliDate.Format("2006-01-02 15:04:05")
fmt.Println(formatted) // Outputs: "1403-01-15 20:14:00"
```

#### `Parse(layout, datetime string) (Jalaali, error)`

Parses a Jalaali date from a string using the specified layout.

```go
jalaaliDate, err := jalaali.Parse("2006-01-02", "1403-01-15")
if err != nil {
    fmt.Println("Error parsing date:", err)
} else {
    fmt.Println(jalaaliDate.String()) // Outputs: "1403-01-15T00:00:00+03:30"
}
```

### Time Zones

#### `Zone() (string, int)`

Returns the time zone name and offset in seconds.

#### `In(loc *time.Location) Jalaali`

Sets the location of the Jalaali instance and returns a new instance.

## Examples

### Create a Jalaali Date

```go
date := jalaali.Date(1403, jalaali.Ordibehesht, 10, 14, 30, 0, 0, jalaali.TehranTz())
fmt.Println(date.String()) // Outputs: "1403-02-10T14:30:00+03:30"
```

### Add Time to a Jalaali Date

```go
date := jalaali.Date(1403, jalaali.Farvardin, 15, 20, 14, 0, 0, jalaali.TehranTz())
newDate := date.AddDate(1, 1, 1)
fmt.Println(newDate.String()) // Outputs: "1404-02-16T20:14:00+03:30"
```

### Format a Jalaali Date

```go
date := jalaali.Date(1403, jalaali.Farvardin, 15, 20, 14, 0, 0, jalaali.TehranTz())
formatted := date.Format("Monday, 02 January 2006")
fmt.Println(formatted) // Outputs: "شنبه, 15 فروردین 1403"
```

### Parse a Jalaali Date

```go
parsedDate, err := jalaali.Parse("2006-01-02", "1403-01-15")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println(parsedDate.String()) // Outputs: "1403-01-15T00:00:00+03:30"
}
```

## License

This library is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
