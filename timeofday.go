package xbase

import (
	"regexp"
	"strconv"
)

// TimeOfDay is a structure holding a 24-hour time point
type TimeOfDay struct {
	hours   int8
	minutes int8
	seconds int8
}

// Mod will make sure all values are within their ranges
func (t TimeOfDay) Mod() (r TimeOfDay) {
	r.seconds = t.seconds
	r.minutes = t.minutes
	r.hours = t.hours
	if r.seconds > 59 {
		r.minutes += r.seconds / 60
		r.seconds = r.seconds % 60
	}
	if r.minutes > 59 {
		r.hours += r.minutes / 60
		r.minutes = r.minutes % 60
	}
	if r.hours > 24 {
		r.hours = r.hours % 24
	}
	return
}

// Add can do an addition of two TimeOfDay objects and return the result
func (t TimeOfDay) Add(o TimeOfDay) (r TimeOfDay) {
	r.hours = t.hours + o.hours
	r.minutes = t.minutes + o.minutes
	r.seconds = t.seconds + o.seconds
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

	return TimeOfDay{hours: hour, minutes: minute, seconds: second}
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
