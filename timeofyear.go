package xbase

import "time"

// Season is any of Summer, Autumn, Winter or Spring
type Season uint8

const (
	// Summer falls in month 6, 7 and 8
	Summer Season = 1
	// Autumn falls in month 9, 10, 11
	Autumn Season = 2
	// Winter falls in month 12, 1, 2
	Winter Season = 3
	// Spring falls in month 3, 4, 5
	Spring Season = 4
)

func getSeasonForMonth(m time.Month) Season {
	switch {
	case m == time.December || m == time.January || m == time.February:
		return Winter
	case m == time.March || m == time.April || m == time.May:
		return Spring
	case m == time.June || m == time.July || m == time.August:
		return Summer
	case m == time.September || m == time.October || m == time.November:
		return Autumn
	}
	return Summer
}

// GetSeason will return the const value of the Season defined by the given month
func GetSeason(dt time.Time) Season {
	return getSeasonForMonth(dt.Month())
}

// TimeOfYear contains information about the Seasonsu
type TimeOfYear struct {
	Season Season
}
