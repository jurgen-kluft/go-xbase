package xbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// TimeOfDay is a structure holding a 24-hour time point
type TimeOfDay struct {
	Hours   int8
	Minutes int8
	Seconds int8
}

// Mod will make sure all values are within their ranges
func (t *TimeOfDay) Mod() {
	if t.Seconds > 59 {
		t.Minutes += t.Seconds / 60
		t.Seconds = t.Seconds % 60
	}
	if t.Minutes > 59 {
		t.Hours += t.Minutes / 60
		t.Minutes = t.Minutes % 60
	}
	if t.Hours > 24 {
		t.Hours = t.Hours % 24
	}
}

// Add can do an addition of two TimeOfDay objects and return the result
func (t *TimeOfDay) Add(o *TimeOfDay) {
	t.Hours = t.Hours + o.Hours
	t.Minutes = t.Minutes + o.Minutes
	t.Seconds = t.Seconds + o.Seconds
	t.Mod()
}

// IsBetween checks if this falls between @start and @end
func (t *TimeOfDay) IsBetween(start, end *TimeOfDay) bool {
	selfInSeconds := int64(t.Hours)*3600 + int64(t.Minutes)*60 + int64(t.Seconds)
	startInSeconds := int64(start.Hours)*3600 + int64(start.Minutes)*60 + int64(start.Seconds)
	endInSeconds := int64(end.Hours)*3600 + int64(end.Minutes)*60 + int64(end.Seconds)
	return selfInSeconds >= startInSeconds && selfInSeconds < endInSeconds
}

// String converts TimeOfDay to a string
func (t *TimeOfDay) String() string {
	return fmt.Sprintf("%d:%d:%d", t.Hours, t.Minutes, t.Seconds)
}

// MarshalJSON encodes the TimeOfDay struct into JSON
func (t *TimeOfDay) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	buffer.WriteString(fmt.Sprintf("%d:%d:%d", t.Hours, t.Minutes, t.Seconds))
	return buffer.Bytes(), nil
}

// UnmarshalJSON decodes JSON into a TimeOfDay struct
func (t *TimeOfDay) UnmarshalJSON(b []byte) error {
	var tod string
	err := json.Unmarshal(b, &tod)
	if err != nil {
		return err
	}
	t.Parse(tod)
	return nil
}

func stripSemi(value string) string {
	if strings.HasPrefix(value, ":") {
		return value[1:]
	}
	return value
}

// Parse parses a time of day string value
func (t *TimeOfDay) Parse(str string) {
	durationRegex := regexp.MustCompile("([0-9]{1,2})(:[0-9]{1,2})(:[0-9]{1,2}){0,1}[ ]{0,1}([AP]M){0,1}")
	matches := durationRegex.FindStringSubmatch(str)

	t.Hours = 0
	t.Minutes = 0
	t.Seconds = 0

	lenMatches := len(matches)
	if lenMatches >= 2 {
		t.Hours = ParseInt8(matches[1])

		if matches[lenMatches-1] == "PM" {
			lenMatches--
			if t.Hours < 12 {
				t.Hours += 12
			}
		} else if matches[lenMatches-1] == "AM" {
			lenMatches--
			if t.Hours > 12 {
				t.Hours -= 12
			}
		}

		if lenMatches >= 3 {
			t.Minutes = ParseInt8(stripSemi(matches[2]))
			if lenMatches >= 4 {
				t.Seconds = ParseInt8(stripSemi(matches[3]))
			}
		}
	}
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
