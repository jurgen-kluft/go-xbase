package xbase

import (
	"testing"
	"time"

	c "github.com/smartystreets/goconvey/convey"
)

// TestTimeOfSun will test some basic sunrise/sunset queries
func TestTimeOfSun(t *testing.T) {
	c.Convey("Given a date-time and latitude,longtitude", t, func() {
		latitude := 31.2222200
		longtitude := -121.4580600
		c.Convey("Shanghai, 1st January 2017", func() {
			loc, _ := time.LoadLocation("Asia/Shanghai")
			dt20170101 := time.Date(2017, 1, 1, 6, 0, 0, 0, loc)
			c.Convey("The sunrise on that day should be 6:52:59", func() {
				sunrise := CalcSunrise(dt20170101, latitude, longtitude)
				c.So(sunrise, c.ShouldResemble, time.Date(2017, 1, 1, 6, 52, 59, 0, loc))
			})
			c.Convey("The sunset on that day should be 17:03:01", func() {
				sunset := CalcSunset(dt20170101, latitude, longtitude)
				//fmt.Println(sunset)
				c.So(sunset, c.ShouldResemble, time.Date(2017, 1, 1, 17, 3, 1, 0, loc))
			})
		})
	})
}
