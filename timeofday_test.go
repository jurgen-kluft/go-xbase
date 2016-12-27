package xbase

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

// TestTimeOfDayParsing will go over all possible formats that should be handled correctly
func TestTimeOfDayParsing(t *testing.T) {
	c.Convey("Given a time value", t, func() {
		c.Convey("When all digits are set", func() {
			clock := ParseTimeOfDay("12:13:14")
			c.Convey("The timeofday should be 12:13:14", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 12, Minutes: 13, Seconds: 14})
			})
		})

		c.Convey("When single digits are set", func() {
			clock := ParseTimeOfDay("2:3:4")
			c.Convey("The timeofday should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing AM", func() {
			clock := ParseTimeOfDay("2:3:4 AM")
			c.Convey("The timeofday should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing AM correction", func() {
			clock := ParseTimeOfDay("14:3:4 AM")
			c.Convey("The timeofday should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing PM", func() {
			clock := ParseTimeOfDay("14:3:4 PM")
			c.Convey("The timeofday should be 14:3:4", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 14, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing PM correction", func() {
			clock := ParseTimeOfDay("2:3:4 PM")
			c.Convey("The timeofday should be 14:3:4", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 14, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing addition, no mod", func() {
			clock1 := ParseTimeOfDay("2:0:0")
			clock2 := ParseTimeOfDay("0:3:4")
			clock := clock1.Add(clock2)
			c.Convey("The result should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing addition, mod", func() {
			clock1 := ParseTimeOfDay("18:57:55")
			clock2 := ParseTimeOfDay("7:5:9")
			clock := clock1.Add(clock2)
			c.Convey("The result should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})
		
		c.Convey("Testing HH:MM, mod", func() {
			clock := ParseTimeOfDay("18:57")
			c.Convey("The result should be 18:57:00", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 18, Minutes: 57, Seconds: 0})
			})
		})
		
		c.Convey("Testing HH:MM PM, mod", func() {
			clock := ParseTimeOfDay("18:57 PM")
			c.Convey("The result should be 18:57:00", func() {
				c.So(clock, c.ShouldResemble, TimeOfDay{Hours: 18, Minutes: 57, Seconds: 0})
			})
		})

	})
}
