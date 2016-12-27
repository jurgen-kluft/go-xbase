package xbase

import (
	"regexp"
	"strconv"
)

// TimeOfDay is a structure holding a 24-hour time point
type TimeOfDay struct {
	Hours   int8
	Minutes int8
	Seconds int8
}

// Mod will make sure all values are within their ranges
func (t TimeOfDay) Mod() (r TimeOfDay) {
	r.Seconds = t.Seconds
	r.Minutes = t.Minutes
	r.Hours = t.Hours
	if r.Seconds > 59 {
		r.Minutes += r.Seconds / 60
		r.Seconds = r.Seconds % 60
	}
	if r.Minutes > 59 {
		r.Hours += r.Minutes / 60
		r.Minutes = r.Minutes % 60
	}
	if r.Hours > 24 {
		r.Hours = r.Hours % 24
	}
	return
}

// Add can do an addition of two TimeOfDay objects and return the result
func (t TimeOfDay) Add(o TimeOfDay) (r TimeOfDay) {
	r.Hours = t.Hours + o.Hours
	r.Minutes = t.Minutes + o.Minutes
	r.Seconds = t.Seconds + o.Seconds
	r = r.Mod()
	return
}

// ParseTimeOfDay parses a time of day string value and returns an instance of TimeOfDay
func ParseTimeOfDay(str string) TimeOfDay {
	durationRegex := regexp.MustCompile("([0-9]{1,2}):([0-9]{1,2}):([0-9]{1,2})[ ]{0,1}([AP]M){0,1}")
	matches := durationRegex.FindStringSubmatch(str)

	hour := ParseInt8(matches[1])
	minute := ParseInt8(matches[2])
	second := ParseInt8(matches[3])

	if len(matches) == 5 {
		if matches[4] == "PM" {
			if hour < 12 {
				hour += 12
			}
		} else if matches[4] == "AM" {
			if hour > 12 {
				hour -= 12
			}
		}
	}

	return TimeOfDay{Hours: hour, Minutes: minute, Seconds: second}
}

// ParseInt8 parses a string containing a number into an 8 bit value
func ParseInt8(value string) int8 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)])
	if err != nil {
		return 0
	}
	return int8(parsed)
}
