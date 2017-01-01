package xbase

import (
	"testing"
	"time"

	c "github.com/smartystreets/goconvey/convey"
)

// TestTimeOfSun will test some basic sunrise/sunset queries
func TestTimeOfYear(t *testing.T) {
	c.Convey("Time of Year helper function", t, func() {
		c.Convey("Season for a couple of dates", func() {
			dt20170101 := time.Date(2017, time.January, 1, 12, 0, 0, 0, time.Local)
			c.Convey("The Season in January should be Winter", func() {
				season := GetSeason(dt20170101)
				c.So(season, c.ShouldEqual, Winter)
			})
			dt20170401 := time.Date(2017, time.April, 1, 12, 0, 0, 0, time.Local)
			c.Convey("The Season in April should be Spring", func() {
				season := GetSeason(dt20170401)
				c.So(season, c.ShouldEqual, Spring)
			})
			dt20170701 := time.Date(2017, time.July, 1, 12, 0, 0, 0, time.Local)
			c.Convey("The Season in July should be Summer", func() {
				season := GetSeason(dt20170701)
				c.So(season, c.ShouldEqual, Summer)
			})
			dt20171001 := time.Date(2017, time.October, 1, 12, 0, 0, 0, time.Local)
			c.Convey("The Season in October should be Autumn", func() {
				season := GetSeason(dt20171001)
				c.So(season, c.ShouldEqual, Autumn)
			})
		})
	})
}
