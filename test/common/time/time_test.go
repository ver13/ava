package time_test

import (
	"testing"
	"time"
	
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	
	. "github.com/ver13/ava/pkg/common/time"
)

type timeSuite struct {
	suite.Suite

	time TimeI
}

func TestTimeInit(t *testing.T) {
	suite.Run(t, new(timeSuite))
}

func (t *timeSuite) BeforeTest() {
	t.T().Log("BeforeTest")
}

func (t *timeSuite) AfterTest() {
	t.T().Log("AfterTest")
}

func (t *timeSuite) SetupSuite() {
	t.T().Log("SetupSuite")
	t.time = NewTime(time.Now())
}

func (t *timeSuite) SetupTest() {
	t.T().Log("SetupTest")
}

func (t *timeSuite) TearDownSuite() {
	t.T().Log("TearDownSuite")
}

func (t *timeSuite) TearDownTest() {
	t.T().Log("TearDownTest")
}

func (t *timeSuite) TestTime_Copy() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_WeekStartsAt() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_WeekEndsAt() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_WeekendDays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Quarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Age() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DaysInMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DaysInYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_WeekOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_WeekOfYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_TimeZone() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Timestamp() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_AddYears() {
	Convey("Given a Time, add 10 years ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddYears(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddYears(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddYear() {
	Convey("Given a Time, add 1 year ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddYear()

			expected, err := Create(2010, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddYear()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddQuarters() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddQuarters(1)

			expected, err := Create(2010, time.February, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddQuarters(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddQuarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddQuarter()

			expected, err := Create(2010, time.February, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddQuarter()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddCenturies() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddCenturies(1)
			
			expected, err := Create(2109, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddCenturies(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddCentury() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddCentury()

			expected, err := Create(2109, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddCentury()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddMonths() {
	Convey("Given a Time, add 10 months ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonths(1)

			expected, err := Create(2009, time.December, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonths(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddMonth() {
	Convey("Given a Time, add 1 month ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonth()

			expected, err := Create(2009, time.December, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonth()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddSeconds() {
	Convey("Given a Time, add 10 seconds ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddSeconds(10)

			expected, err := Create(2009, time.November, 10, 23, 0, 10, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddSeconds(10)

			expected, err := Create(2009, time.November, 10, 23, 0, 10, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddSecond() {
	Convey("Given a Time, add 1 second ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddSecond()

			expected, err := Create(2009, time.November, 10, 23, 0, 1, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddSecond()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddDays() {
	Convey("Given a Time, add 10 days ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddDays(10)

			expected, err := Create(2009, time.November, 20, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddDays(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddDay() {
	Convey("Given a Time, add 1 day ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddDay()

			expected, err := Create(2009, time.November, 11, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddDay()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddWeekdays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeekdays(1)

			expected, err := Create(2009, time.November, 11, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeekdays(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddWeekday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeekday()

			expected, err := Create(2009, time.November, 11, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeekday()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddWeeks() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeeks(1)

			expected, err := Create(2009, time.November, 17, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeeks(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddWeek() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeek()

			expected, err := Create(2009, time.November, 17, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddWeek()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddHours() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 2, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddHours(10)

			expected, err := Create(2009, time.November, 10, 12, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddHours(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddHour() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 20, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddHour()

			expected, err := Create(2009, time.November, 10, 21, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddHour()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddMonthsNoOverflow() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonthsNoOverflow(1)

			expected, err := Create(2009, time.December, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonthsNoOverflow(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_PreviousMonthLastDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_AddMonthNoOverflow() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonthNoOverflow()

			expected, err := Create(2009, time.December, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMonthNoOverflow()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddMinutes() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMinutes(10)

			expected, err := Create(2009, time.November, 10, 23, 10, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMinutes(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_AddMinute() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMinute()

			expected, err := Create(2009, time.November, 10, 23, 1, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.AddMinute()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubYear()

			expected, err := Create(2008, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubYear()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubYears() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubYears(1)

			expected, err := Create(2008, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubYears(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubQuarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubQuarter()

			expected, err := Create(2009, time.August, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubQuarter()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubQuarters() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubQuarters(1)

			expected, err := Create(2009, time.August, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubQuarters(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubCentury() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubCentury()

			expected, err := Create(1909, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubCentury()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubCenturies() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubCenturies(1)

			expected, err := Create(1909, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubCenturies(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonth()

			expected, err := Create(2009, time.October, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonth()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubMonths() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonths(2)

			expected, err := Create(2009, time.September, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonths(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubMonthNoOverflow() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonthNoOverflow()

			expected, err := Create(2009, time.October, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonthNoOverflow()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubMonthsNoOverflow() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonthsNoOverflow(1)

			expected, err := Create(2009, time.October, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMonthsNoOverflow(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubDay()

			expected, err := Create(2009, time.November, 9, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubDay()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubDays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubDays(2)

			expected, err := Create(2009, time.November, 8, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubDays(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubWeekday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeekday()

			expected, err := Create(2009, time.November, 9, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeekday()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubWeekdays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeekdays(1)

			expected, err := Create(2009, time.November, 9, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeekdays(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubWeek() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeek()

			expected, err := Create(2009, time.November, 3, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeek()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubWeeks() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeeks(1)

			expected, err := Create(2009, time.November, 3, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubWeeks(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubHour() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubHour()

			expected, err := Create(2009, time.November, 10, 22, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubHour()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubHours() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubHours(2)

			expected, err := Create(2009, time.November, 10, 21, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubHours(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubMinute() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMinute()

			expected, err := Create(2009, time.November, 10, 22, 59, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMinute()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubMinutes() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMinutes(10)

			expected, err := Create(2009, time.November, 10, 22, 50, 0, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubMinutes(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubSecond() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubSecond()

			expected, err := Create(2009, time.November, 10, 22, 59, 59, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubSecond()

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SubSeconds() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubSeconds(10)

			expected, err := Create(2009, time.November, 10, 22, 59, 50, 0, "UTC")
			So(err, ShouldBeNil)
			So(expected, ShouldResemble, actual)
		})
		Convey("Went it's wrong ", func() {
			c, err := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
			So(err, ShouldBeNil)

			actual := c.SubSeconds(10)

			expected, err := Create(2019, time.November, 10, 23, 0, 0, 0, "America/New_York")
			So(err, ShouldBeNil)
			So(expected.GetTime(), ShouldNotResemble, actual.GetTime())
		})
	})
}

func (t *timeSuite) TestTime_SetYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetHour() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetMinute() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetSecond() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetDate() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetDateTime() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetTimeFromTimeString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetWeekEndsAt() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetWeekStartsAt() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetWeekendDays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetTimestamp() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetTimeZone() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_ResetStringFormat() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SetStringFormat() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DateString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_FormattedDateString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_TimeString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DateTimeString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DayDateTimeString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_AtomString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_CookieString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_ISO8601String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_RFC822String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_RFC850String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_RFC1036String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_RFC1123String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_RFC2822String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_RFC3339String() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_RSSString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_W3CString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsWeekday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsWeekend() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsYesterday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsToday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsTomorrow() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsFuture() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsPast() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsLeapYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsLongYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsSameAs() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsCurrentYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsSameYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsCurrentMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsSameMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsSameDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsSunday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsMonday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsTuesday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsWednesday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsThursday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsFriday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsSaturday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsLastWeek() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_IsLastMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Eq() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EqualTo() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Ne() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_NotEqualTo() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Gt() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_GreaterThan() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Gte() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_GreaterThanOrEqualTo() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Lt() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_LessThan() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Lte() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_LessThanOrEqualTo() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Between() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Closest() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Farthest() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Min() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Minimum() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Max() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Maximum() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInYears() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInMonths() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_hasRemainingHours() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffDurationInString() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInWeeks() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInDays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInNights() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInDaysFiltered() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInHoursFiltered() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInWeekdays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInWeekendDays() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffFiltered() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInHours() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInMinutes() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_DiffInSeconds() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SecondsSinceMidnight() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_SecondsUntilEndOfDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_StartOfDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EndOfDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_StartOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EndOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_StartOfQuarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EndOfQuarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_StartOfYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EndOfYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_StartOfDecade() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EndOfDecade() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_StartOfCentury() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EndOfCentury() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_StartOfWeek() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_EndOfWeek() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Next() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_NextWeekday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_PreviousWeekday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_NextWeekendDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_PreviousWeekendDay() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Previous() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_FirstOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_LastOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_LastDayOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_FirstDayOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_NthOfMonth() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_FirstOfQuarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_LastOfQuarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_NthOfQuarter() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_FirstOfYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_LastOfYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_NthOfYear() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Average() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
func (t *timeSuite) TestTime_Clock() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Nanosecond() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Date() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_In() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Day() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Month() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Year() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_After() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_GetTime() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Format() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Sub() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_UnixNano() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Location() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Add() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (t *timeSuite) TestTime_Weekday() {
	Convey("Given a ", t.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
