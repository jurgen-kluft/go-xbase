package xbase

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

// TestTimeOfDayParsing will go over all possible formats that should be handled correctly
func TestTimeOfDayParsing(t *testing.T) {
	c.Convey("Given a time value", t, func() {
		c.Convey("When all digits are set", func() {
			clock := &TimeOfDay{}
			clock.Parse("12:13:14")
			c.Convey("The timeofday should be 12:13:14", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 12, Minutes: 13, Seconds: 14})
			})
		})

		c.Convey("When single digits are set", func() {
			clock := &TimeOfDay{}
			clock.Parse("2:3:4")
			c.Convey("The timeofday should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing AM", func() {
			clock := &TimeOfDay{}
			clock.Parse("2:3:4 AM")
			c.Convey("The timeofday should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing AM correction", func() {
			clock := &TimeOfDay{}
			clock.Parse("14:3:4 AM")
			c.Convey("The timeofday should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing PM", func() {
			clock := &TimeOfDay{}
			clock.Parse("14:3:4 PM")
			c.Convey("The timeofday should be 14:3:4", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 14, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing PM correction", func() {
			clock := &TimeOfDay{}
			clock.Parse("2:3:4 PM")
			c.Convey("The timeofday should be 14:3:4", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 14, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing addition, no mod", func() {
			clock := &TimeOfDay{}
			clock.Parse("2:0:0")
			clock2 := &TimeOfDay{}
			clock2.Parse("0:3:4")
			clock.Add(clock2)
			c.Convey("The result should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing addition, mod", func() {
			clock := &TimeOfDay{}
			clock.Parse("18:57:55")
			clock2 := &TimeOfDay{}
			clock2.Parse("7:5:9")
			clock.Add(clock2)
			c.Convey("The result should be 2:3:4", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 2, Minutes: 3, Seconds: 4})
			})
		})

		c.Convey("Testing HH:MM", func() {
			clock := &TimeOfDay{}
			clock.Parse("18:57")
			c.Convey("The result should be 18:57:00", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 18, Minutes: 57, Seconds: 0})
			})
		})

		c.Convey("Testing HH:MM AM", func() {
			clock := &TimeOfDay{}
			clock.Parse("8:57 AM")
			c.Convey("The result should be 8:57:00", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 8, Minutes: 57, Seconds: 0})
			})
		})

		c.Convey("Testing HH:MM PM", func() {
			clock := &TimeOfDay{}
			clock.Parse("18:57 PM")
			c.Convey("The result should be 18:57:00", func() {
				c.So(clock, c.ShouldResemble, &TimeOfDay{Hours: 18, Minutes: 57, Seconds: 0})
			})
		})

		c.Convey("Testing IsBetween(start,end)", func() {
			clock := &TimeOfDay{}
			start := &TimeOfDay{}
			end := &TimeOfDay{}

			clock.Parse("18:57 PM")
			start.Parse("16:00 PM")
			end.Parse("19:00PM")

			c.Convey("(18:57 PM).IsBetween(16:00, 19:00) == true", func() {
				c.So(clock.IsBetween(start, end), c.ShouldEqual, true)
			})

			clock.Parse("19:01 PM")
			c.Convey("(19:01 PM).IsBetween(16:00, 19:00) == false", func() {
				c.So(clock.IsBetween(start, end), c.ShouldEqual, false)
			})

			clock.Parse("7:30")
			start.Parse("6:00")
			end.Parse("11:00")

			c.Convey("(7:30).IsBetween(6:00, 11:00) == true", func() {
				c.So(clock.IsBetween(start, end), c.ShouldEqual, true)
			})
		})

	})
}
