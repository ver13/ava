package time

import (
	"fmt"
	"math"
	"strings"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorTimeAVA "github.com/ver13/ava/pkg/common/time/error"
)

// Represents the number of elements in a given period
const (
	secondsPerMinute  = 60
	minutesPerHour    = 60
	hoursPerDay       = 24
	daysPerWeek       = 7
	monthsPerQuarter  = 3
	monthsPerYear     = 12
	yearsPerCenturies = 100
	yearsPerDecade    = 10
	weeksPerLongYear  = 53
	daysInLeapYear    = 366
	daysInNormalYear  = 365
	secondsInWeek     = 691200
)

// Represents the different string formats for dates
const (
	DefaultFormat       = "2006-01-02 15:04:05"
	DateFormat          = "2006-01-02"
	FormattedDateFormat = "Jan 2, 2006"
	TimeFormat          = "15:04:05"
	HourMinuteFormat    = "15:04"
	HourFormat          = "15"
	DayDateTimeFormat   = "Mon, Aug 2, 2006 3:04 PM"
	CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC822Format        = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC2822Format       = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Format       = "2006-01-02T15:04:05-07:00"
	RSSFormat           = "Mon, 02 Jan 2006 15:04:05 -0700"
)

// Provides a simple api extension for Times.
type Time struct {
	time.Time
	weekStartsAt time.Weekday
	weekEndsAt   time.Weekday
	weekendDays  []time.Weekday
	stringFormat string
}

// Used for testing purposes
var (
	isTimeFrozen      bool
	currentFrozenTime time.Time
)

// NewTime returns a pointer to a new Times instance
func NewTime(t time.Time) *Time {
	wds := []time.Weekday{
		time.Saturday,
		time.Sunday,
	}
	return &Time{
		Time:         t,
		weekStartsAt: time.Monday,
		weekEndsAt:   time.Sunday,
		weekendDays:  wds,
		stringFormat: DefaultFormat,
	}
}

// Freeze allows time to be frozen to facilitate testing
func Freeze(time time.Time) {
	currentFrozenTime = time
	isTimeFrozen = true
}

// UnFreeze returns time to normal operation
func UnFreeze() {
	isTimeFrozen = false
}

// IsTimeFrozen allows checking if time has been frozen
func IsTimeFrozen() bool {
	return isTimeFrozen
}

// After will be behave like time.After unless time has been frozen
// If time is frozen it will add the expected delay and immediately send the frozen time on the returned channel
func After(d time.Duration) <-chan time.Time {
	if isTimeFrozen {
		currentFrozenTime = currentFrozenTime.Add(d)
		c := make(chan time.Time, 1)
		c <- currentFrozenTime
		return c
	}

	return time.After(d)
}

// Tick will be behave like time.Tick unless time has been frozen
// If time is frozen it will tick normally but the date will be based on the frozen date
func Tick(d time.Duration) <-chan time.Time {
	if isTimeFrozen {
		c := make(chan time.Time, 1)
		go func() {
			for {
				currentFrozenTime = currentFrozenTime.Add(d)
				c <- currentFrozenTime
			}
		}()
		return c
	}

	return time.Tick(d)
}

// Sleep will be behave like time.Sleep unless time has been frozen
// If time is frozen it will add the expected sleep delay and return immediately
func Sleep(d time.Duration) {
	if isTimeFrozen && d > 0 {
		currentFrozenTime = currentFrozenTime.Add(d)

		return
	}

	time.Sleep(d)
}

// create returns a new Times pointe. It is a helper function to create new dates
func create(y int, mon time.Month, d, h, m, s, ns int, l *time.Location) *Time {
	return NewTime(time.Date(y, mon, d, h, m, s, ns, l))
}

// Create returns a new pointer to Times instance from a specific date and time.
// If the location is invalid, it returns an ErrorAVA instead.
func Create(y int, mon time.Month, d, h, m, s, ns int, location string) (*Time, *errorAVA.Error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, errorTimeAVA.LoadLocation(err, location)
	}
	return create(y, mon, d, h, m, s, ns, l), nil
}

// CreateFromDate returns a new pointer to a Times instance from just a date.
// The time portion is set to now.
// If the location is invalid, it returns an ErrorAVA instead.
func CreateFromDate(y int, mon time.Month, d int, location string) (*Time, *errorAVA.Error) {
	h, m, s := Now().Clock()
	ns := Now().Nanosecond()

	return Create(y, mon, d, h, m, s, ns, location)
}

// CreateFromTime returns a new pointer to a Times instance from just a date.
// The time portion is set to now.
// If the locations is invalid, it returns an ErrorAVA instead.
func CreateFromTime(h, m, s, ns int, location string) (*Time, *errorAVA.Error) {
	y, mon, d := Now().Date()

	return Create(y, mon, d, h, m, s, ns, location)
}

// CreateFromFormat returns a new pointer to a Times instance from a specific format.
// If the location is invalid, it returns an ErrorAVA instead.
func CreateFromFormat(layout, value string, location string) (*Time, *errorAVA.Error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, errorTimeAVA.LoadLocation(err, location)
	}
	t, err := time.ParseInLocation(layout, value, l)
	if err != nil {
		return nil, errorTimeAVA.ParseInLocation(err, fmt.Sprintf("Layout: [%s] Value: [%s] Location: [%s]", layout, value, location))
	}

	return NewTime(t), nil
}

// CreateFromTimestamp returns a new pointer to a Times instance from a timestamp.
// If the location is invalid, it returns an ErrorAVA instead.
func CreateFromTimestamp(timestamp int64, location string) (*Time, *errorAVA.Error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, errorTimeAVA.LoadLocation(err, location)
	}
	t := NewTime(Now().In(l))
	t.SetTimestamp(timestamp)

	return t, nil
}

// CreateFromTimestampUTC returns a new pointer to a Times instance from an UTC timestamp.
// If the location is invalid, it returns an ErrorAVA instead.
func CreateFromTimestampUTC(timestamp int64) (*Time, *errorAVA.Error) {
	return CreateFromTimestamp(timestamp, "UTC")
}

// CreateFromMonthAndYear returns a new pointer to a Times instance from a specific month and year.
// If the location is invalid, it returns an ErrorAVA instead.
func CreateFromMonthAndYear(y int, mon time.Month, location string) (*Time, *errorAVA.Error) {
	_, _, d := Now().Date()
	h, m, s := Now().Clock()
	ns := Now().Nanosecond()

	return Create(y, mon, d, h, m, s, ns, location)
}

// Parser returns a pointer to a new Times instance from a string
// If the location is invalid, it returns an ErrorAVA instead.
func Parse(layout, value, location string) (*Time, *errorAVA.Error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, errorTimeAVA.LoadLocation(err, location)
	}
	t, err := time.ParseInLocation(layout, value, l)
	if err != nil {
		return nil, errorTimeAVA.ParseInLocation(err, fmt.Sprintf("Layout: [%s] Value: [%s] Location: [%s]", layout, value, location))
	}

	return NewTime(t), nil
}

// Today returns a pointer to a new Times instance for today
// If the location is invalid, it returns an ErrorAVA instead.
func Today(location string) (*Time, *errorAVA.Error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, errorTimeAVA.LoadLocation(err, location)
	}

	return NewTime(Now().In(l)), nil
}

// Tomorrow returns a pointer to a new Times instance for tomorrow
// If the location is invalid, it returns an ErrorAVA instead.
func Tomorrow(location string) (*Time, *errorAVA.Error) {
	c, err := Today(location)
	if err != nil {
		return nil, err
	}

	return c.AddDay(), nil
}

// Yesterday returns a pointer to a new Times instance for yesterday
// If the location is invalid, it returns an ErrorAVA instead.
func Yesterday(location string) (*Time, *errorAVA.Error) {
	c, err := Today(location)
	if err != nil {
		return nil, err
	}

	return c.SubDay(), nil
}

// unix*TimenSeconds represents the number of seconds between Year 1 and 1970
const unixTimeSeconds = 62135596801

const maxNSecs = 999999999

// MaxValue returns a pointer to a new Times instance for greatest supported date
func MaxValue() *Time {
	return NewTime(time.Unix(math.MaxInt64-unixTimeSeconds, maxNSecs))
}

// MinValue returns a pointer to a new Times instance for lowest supported date
func MinValue() *Time {
	return NewTime(time.Unix(math.MinInt64+unixTimeSeconds, 0))
}

// Now returns a new Times instance for right now in current localtime
func Now() *Time {
	if isTimeFrozen {
		return NewTime(currentFrozenTime)
	}

	return NewTime(time.Now())
}

// NowInLocation returns a new Times instance for right now in given location.
// The location is in IANA Times Zone database, such as "America/New_York".
func NowInLocation(loc string) (*Time, *errorAVA.Error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, errorTimeAVA.LoadLocation(err, loc)
	}
	return nowIn(l), nil
}

func nowIn(loc *time.Location) *Time {
	return NewTime(Now().In(loc))
}

// Copy returns a new copy of the current Times instance
func (c *Time) Copy() *Time {
	return create(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// WeekStartsAt get the starting day of the week
func (c *Time) WeekStartsAt() time.Weekday {
	return c.weekStartsAt
}

// WeekEndsAt gets the ending day of the week
func (c *Time) WeekEndsAt() time.Weekday {
	return c.weekEndsAt
}

// WeekendDays gets the weekend days of the week
func (c *Time) WeekendDays() []time.Weekday {
	return c.weekendDays
}

// Quarter gets the current quarter
func (c *Time) Quarter() int {
	month := c.Month()
	switch {
	case month < 4:
		return 1
	case month >= 4 && month < 7:
		return 2
	case month >= 7 && month < 10:
		return 3
	}
	return 4
}

// Age gets the age from the current instance time to now
func (c *Time) Age() int {
	return int(c.DiffInYears(Now(), true))
}

// DaysInMonth returns the number of days in the month
func (c *Time) DaysInMonth() int {
	return c.EndOfMonth().Day()
}

// DaysInYear returns the number of days in the year
func (c *Time) DaysInYear() int {
	if c.IsLeapYear() {
		return daysInLeapYear
	}

	return daysInNormalYear
}

// WeekOfMonth returns the week of the month
func (c *Time) WeekOfMonth() int {
	w := math.Ceil(float64(c.Day() / daysPerWeek))
	return int(w + 1)
}

// WeekOfYear returns the week of the current year.
// This is an alias for time.ISOWeek
func (c *Time) WeekOfYear() (int, int) {
	return c.ISOWeek()
}

// TimeZone gets the current timezone
func (c *Time) TimeZone() string {
	return c.Location().String()
}

// Timestamp gets the current time since January 1, 1970 UTC
func (c *Time) Timestamp() int64 {
	return c.Unix()
}

// String gets the current date using the previously set format
func (c *Time) String() string {
	return c.Format(c.stringFormat)
}

// AddYears adds a year to the current time.
// Positive values travel forward while negative values travel into the past
func (c *Time) AddYears(y int) *Time {
	return NewTime(c.AddDate(y, 0, 0))
}

// AddYear adds a year to the current time
func (c *Time) AddYear() *Time {
	return c.AddYears(1)
}

// AddQuarters adds quarters to the current time.
// Positive values travel forward while negative values travel into the past
func (c *Time) AddQuarters(q int) *Time {
	return NewTime(c.AddDate(0, monthsPerQuarter*q, 0))
}

// AddQuarter adds a quarter to the current time
func (c *Time) AddQuarter() *Time {
	return c.AddQuarters(1)
}

// AddCenturies adds centuries to the time.
// Positive values travels forward while negative values travels into the past
func (c *Time) AddCenturies(cent int) *Time {
	return NewTime(c.AddDate(yearsPerCenturies*cent, 0, 0))
}

// AddCentury adds a century to the current time
func (c *Time) AddCentury() *Time {
	return c.AddCenturies(1)
}

// AddMonths adds months to the current time.
// Positive value travels forward while negative values travels into the past
func (c *Time) AddMonths(m int) *Time {
	return NewTime(c.AddDate(0, m, 0))
}

// AddMonth adds a month to the current time
func (c *Time) AddMonth() *Time {
	return c.AddMonths(1)
}

// AddSeconds adds seconds to the current time.
// Positive values travels forward while negative values travels into the past.
func (c *Time) AddSeconds(s int) *Time {
	d := time.Duration(s) * time.Second
	return NewTime(c.Add(d))
}

// AddSecond adds a second to the time
func (c *Time) AddSecond() *Time {
	return c.AddSeconds(1)
}

// AddDays adds a day to the current time.
// Positive value travels forward while negative value travels into the past
func (c *Time) AddDays(d int) *Time {
	return NewTime(c.AddDate(0, 0, d))
}

// AddDay adds a day to the current time
func (c *Time) AddDay() *Time {
	return c.AddDays(1)
}

// AddWeekdays adds a weekday to the current time.
// Positive value travels forward while negative value travels into the past
func (c *Time) AddWeekdays(wd int) *Time {
	d := 1
	if wd < 0 {
		wd, d = -wd, -d
	}
	t := c.Copy()
	for wd > 0 {
		t = t.AddDays(d)
		if t.IsWeekday() {
			wd--
		}
	}

	return t
}

// AddWeekday adds a weekday to the current time
func (c *Time) AddWeekday() *Time {
	return c.AddWeekdays(1)
}

// AddWeeks adds a week to the current time.
// Positive value travels forward while negative value travels into the past.
func (c *Time) AddWeeks(w int) *Time {
	return NewTime(c.AddDate(0, 0, daysPerWeek*w))
}

// AddWeek adds a week to the current time
func (c *Time) AddWeek() *Time {
	return c.AddWeeks(1)
}

// AddHours adds an hour to the current time.
// Positive value travels forward while negative value travels into the past
func (c *Time) AddHours(h int) *Time {
	d := time.Duration(h) * time.Hour

	return NewTime(c.Add(d))
}

// AddHour adds an hour to the current time
func (c *Time) AddHour() *Time {
	return c.AddHours(1)
}

// AddMonthsNoOverflow adds a month to the current time, not overflowing in case the
// destination month has less days than the current one.
// Positive value travels forward while negative value travels into the past.
func (c *Time) AddMonthsNoOverflow(m int) *Time {
	addedDate := NewTime(c.AddDate(0, m, 0))
	if c.Day() != addedDate.Day() {
		return addedDate.PreviousMonthLastDay()
	}

	return addedDate
}

// PreviousMonthLastDay returns the last day of the previous month
func (c *Time) PreviousMonthLastDay() *Time {
	return NewTime(c.AddDate(0, 0, -c.Day()))
}

// AddMonthNoOverflow adds a month with no overflow to the current time
func (c *Time) AddMonthNoOverflow() *Time {
	return c.AddMonthsNoOverflow(1)
}

// AddMinutes adds minutes to the current time.
// Positive value travels forward while negative value travels into the past.
func (c *Time) AddMinutes(m int) *Time {
	d := time.Duration(m) * time.Minute

	return NewTime(c.Add(d))
}

// AddMinute adds a minute to the current time
func (c *Time) AddMinute() *Time {
	return c.AddMinutes(1)
}

// SubYear removes a year from the current time
func (c *Time) SubYear() *Time {
	return c.SubYears(1)
}

// SubYears removes years from current time
func (c *Time) SubYears(y int) *Time {
	return c.AddYears(-1 * y)
}

// SubQuarter removes a quarter from the current time
func (c *Time) SubQuarter() *Time {
	return c.SubQuarters(1)
}

// SubQuarters removes quarters from current time
func (c *Time) SubQuarters(q int) *Time {
	return c.AddQuarters(-q)
}

// SubCentury removes a century from the current time
func (c *Time) SubCentury() *Time {
	return c.SubCenturies(1)
}

// SubCenturies removes centuries from the current time
func (c *Time) SubCenturies(cent int) *Time {
	return c.AddCenturies(-cent)
}

// SubMonth removes a month from the current time
func (c *Time) SubMonth() *Time {
	return c.SubMonths(1)
}

// SubMonths removes months from the current time
func (c *Time) SubMonths(m int) *Time {
	return c.AddMonths(-m)
}

// SubMonthNoOverflow remove a month with no overflow from the current time
func (c *Time) SubMonthNoOverflow() *Time {
	return c.SubMonthsNoOverflow(1)
}

// SubMonthsNoOverflow removes months with no overflow from the current time
func (c *Time) SubMonthsNoOverflow(m int) *Time {
	return c.AddMonthsNoOverflow(-m)
}

// SubDay removes a day from the current instance
func (c *Time) SubDay() *Time {
	return c.SubDays(1)
}

// SubDays removes days from the current time
func (c *Time) SubDays(d int) *Time {
	return c.AddDays(-d)
}

// SubWeekday removes a weekday from the current time
func (c *Time) SubWeekday() *Time {
	return c.SubWeekdays(1)
}

// SubWeekdays removes a weekday from the current time
func (c *Time) SubWeekdays(wd int) *Time {
	return c.AddWeekdays(-wd)
}

// SubWeek removes a week from the current time
func (c *Time) SubWeek() *Time {
	return c.SubWeeks(1)
}

// SubWeeks removes weeks to the current time
func (c *Time) SubWeeks(w int) *Time {
	return c.AddWeeks(-w)
}

// SubHour removes an hour from the current time
func (c *Time) SubHour() *Time {
	return c.SubHours(1)
}

// SubHours removes hours from the current time
func (c *Time) SubHours(h int) *Time {
	return c.AddHours(-h)
}

// SubMinute removes a minute from the current time
func (c *Time) SubMinute() *Time {
	return c.SubMinutes(1)
}

// SubMinutes removes minutes from the current time
func (c *Time) SubMinutes(m int) *Time {
	return c.AddMinutes(-m)
}

// SubSecond removes a second from the current time
func (c *Time) SubSecond() *Time {
	return c.SubSeconds(1)
}

// SubSeconds removes seconds from the current time
func (c *Time) SubSeconds(s int) *Time {
	return c.AddSeconds(-s)
}

// SetYear sets the year of the current time
func (c *Time) SetYear(y int) {
	c.Time = time.Date(y, c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetMonth sets the month of the current time
func (c *Time) SetMonth(m time.Month) {
	c.Time = time.Date(c.Year(), m, c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetDay sets the day of the current time
func (c *Time) SetDay(d int) {
	c.Time = time.Date(c.Year(), c.Month(), d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetHour sets the hour of the current time
func (c *Time) SetHour(h int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), h, c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetMinute sets the minute of the current time
func (c *Time) SetMinute(m int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), m, c.Second(), c.Nanosecond(), c.Location())
}

// SetSecond sets the second of the current time
func (c *Time) SetSecond(s int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), s, c.Nanosecond(), c.Location())
}

// SetDate sets only the date of the current time
func (c *Time) SetDate(y int, m time.Month, d int) {
	c.Time = time.Date(y, m, d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetDateTime sets the date and the time
func (c *Time) SetDateTime(y int, mon time.Month, d, h, m, s int) {
	c.Time = time.Date(y, mon, d, h, m, s, c.Nanosecond(), c.Location())
}

// SetTimeFromTimeString receives a string and sets the current time
// It accepts the following formats: "hh:mm:ss", "hh:mm" and "hh"
func (c *Time) SetTimeFromTimeString(timeString string) *errorAVA.Error {
	layouts := []string{
		TimeFormat,
		HourMinuteFormat,
		HourFormat,
	}

	var t time.Time
	var err error
	for i, layout := range layouts {
		t, err = time.Parse(layout, timeString)
		if err == nil {
			h, m, s := t.Clock()
			switch i {
			case 1:
				s = c.Second()
			case 2:
				m, s = c.Minute(), c.Second()
			}
			c.SetHour(h)
			c.SetMinute(m)
			c.SetSecond(s)
			return nil
		}
	}

	return errorTimeAVA.TimeFormat(err, "Only supports hh:mm:ss, hh:mm and hh formats")
}

// SetWeekEndsAt sets the last day of week
func (c *Time) SetWeekEndsAt(wd time.Weekday) {
	c.weekEndsAt = wd
}

// SetWeekStartsAt sets the first day of week
func (c *Time) SetWeekStartsAt(wd time.Weekday) {
	c.weekStartsAt = wd
}

// SetWeekendDays sets the weekend days
func (c *Time) SetWeekendDays(wds []time.Weekday) {
	c.weekendDays = wds
}

// SetTimestamp sets the current time given a timestamp
func (c *Time) SetTimestamp(sec int64) {
	c.Time = time.Unix(sec, 0).In(c.Location())
}

// SetTimeZone sets the location from a string
// If the location is invalid, it returns an ErrorAVA instead.
func (c *Time) SetTimeZone(name string) *errorAVA.Error {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return errorTimeAVA.LoadLocation(err, name)
	}
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), loc)

	return nil
}

// ResetStringFormat changes the format to the DefaultFormat
func (c *Time) ResetStringFormat() {
	c.stringFormat = DefaultFormat
}

// SetStringFormat formats the current time with the set format string
func (c *Time) SetStringFormat(format string) {
	c.stringFormat = format
}

// DateString return the current time in Y-m-d format
func (c *Time) DateString() string {
	return c.Format(DateFormat)
}

// FormattedDateString returns the current time as a readable date
func (c *Time) FormattedDateString() string {
	return c.Format(FormattedDateFormat)
}

// TimeString returns the current time in hh:mm:ss format
func (c *Time) TimeString() string {
	return c.Format(TimeFormat)
}

// DateTimeString returns the current time in Y-m-d hh:mm:ss format
func (c *Time) DateTimeString() string {
	return c.Format(DefaultFormat)
}

// DayDateTimeString returns the current time with a day, date and time format
func (c *Time) DayDateTimeString() string {
	return c.Format(DayDateTimeFormat)
}

// AtomString formats the current time to a Atom date format
func (c *Time) AtomString() string {
	return c.Format(RFC3339Format)
}

// CookieString formats the current time to a Cookie date format
func (c *Time) CookieString() string {
	return c.Format(CookieFormat)
}

// ISO8601String returns the current time in ISO8601 format
func (c *Time) ISO8601String() string {
	return c.Format(RFC3339Format)
}

// RFC822String returns the current time in RFC 822 format
func (c *Time) RFC822String() string {
	return c.Format(RFC822Format)
}

// RFC850String returns the current time in RFC 850 format
func (c *Time) RFC850String() string {
	return c.Format(time.RFC850)
}

// RFC1036String returns the current time in RFC 1036 format
func (c *Time) RFC1036String() string {
	return c.Format(RFC1036Format)
}

// RFC1123String returns the current time in RFC 1123 format
func (c *Time) RFC1123String() string {
	return c.Format(time.RFC1123Z)
}

// RFC2822String returns the current time in RFC 2822 format
func (c *Time) RFC2822String() string {
	return c.Format(RFC2822Format)
}

// RFC3339String returns the current time in RFC 3339 format
func (c *Time) RFC3339String() string {
	return c.Format(RFC3339Format)
}

// RSSString returns the current time for RSS format
func (c *Time) RSSString() string {
	return c.Format(RSSFormat)
}

// W3CString returns the current time for WWW Consortium format
func (c *Time) W3CString() string {
	return c.Format(RFC3339Format)
}

// IsWeekday determines if the current time is a weekday
func (c *Time) IsWeekday() bool {
	return !c.IsWeekend()
}

// IsWeekend determines if the current time is a weekend day
func (c *Time) IsWeekend() bool {
	d := c.Weekday()
	for _, wd := range c.WeekendDays() {
		if d == wd {
			return true
		}
	}

	return false
}

// IsYesterday determines if the current time is yesterday
func (c *Time) IsYesterday() bool {
	n := Now().SubDay()

	return c.IsSameDay(n)
}

// IsToday determines if the current time is today
func (c *Time) IsToday() bool {
	return c.IsSameDay(Now())
}

// IsTomorrow determines if the current time is tomorrow
func (c *Time) IsTomorrow() bool {
	n := Now().AddDay()

	return c.IsSameDay(n)
}

// IsFuture determines if the current time is in the future, ie. greater (after) than now
func (c *Time) IsFuture() bool {
	return c.After(Now().GetTime())
}

func (c *Time) After(u time.Time) bool {
	return c.Time.After(u)
}

// IsPast determines if the current time is in the past, ie. less (before) than now
func (c *Time) IsPast() bool {
	return c.Before(Now().GetTime())
}

func (c *Time) GetTime() time.Time {
	return c.Time
}

// IsLeapYear determines if current current time is a leap year
func (c *Time) IsLeapYear() bool {
	y := c.Year()
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		return true
	}

	return false
}

// IsLongYear determines if the instance is a long year
func (c *Time) IsLongYear() bool {
	timeAVA := create(c.Year(), time.December, 31, 0, 0, 0, 0, c.Location())
	_, w := timeAVA.WeekOfYear()

	return w == weeksPerLongYear
}

// IsSameAs compares the formatted values of the two dates.
// If passed date is nil, compares against today
func (c *Time) IsSameAs(format string, timeAVA *Time) bool {
	if timeAVA == nil {
		return c.Format(DefaultFormat) == Now().Format(format)
	}

	return c.Format(DefaultFormat) == timeAVA.Format(format)
}

// IsCurrentYear determines if the current time is in the current year
func (c *Time) IsCurrentYear() bool {
	return c.Year() == Now().Year()
}

// IsSameYear checks if the passed in date is in the same year as the current time year.
// If passed date is nil, compares against today
func (c *Time) IsSameYear(timeAVA *Time) bool {
	if timeAVA == nil {
		return c.Year() == nowIn(c.Location()).Year()
	}

	return c.Year() == timeAVA.Year()
}

// IsCurrentMonth determines if the current time is in the current month
func (c *Time) IsCurrentMonth() bool {
	return c.Month() == Now().Month()
}

// IsSameMonth checks if the passed in date is in the same month as the current month
// If passed date is nil, compares against today
func (c *Time) IsSameMonth(timeAVA *Time, sameYear bool) bool {
	m := nowIn(c.Location()).Month()
	if timeAVA != nil {
		m = timeAVA.Month()
	}
	if sameYear {
		return c.IsSameYear(timeAVA) && c.Month() == m
	}

	return c.Month() == m
}

// IsSameDay checks if the passed in date is the same day as the current day.
// If passed date is nil, compares against today
func (c *Time) IsSameDay(timeAVA *Time) bool {
	n := nowIn(c.Location())
	if timeAVA != nil {
		n = timeAVA
	}

	return c.Year() == n.Year() && c.Month() == n.Month() && c.Day() == n.Day()
}

// IsSunday checks if this day is a Sunday.
func (c *Time) IsSunday() bool {
	return c.Weekday() == time.Sunday
}

// IsMonday checks if this day is a Monday.
func (c *Time) IsMonday() bool {
	return c.Weekday() == time.Monday
}

// IsTuesday checks if this day is a Tuesday.
func (c *Time) IsTuesday() bool {
	return c.Weekday() == time.Tuesday
}

// IsWednesday checks if this day is a Wednesday.
func (c *Time) IsWednesday() bool {
	return c.Weekday() == time.Wednesday
}

// IsThursday checks if this day is a Thursday.
func (c *Time) IsThursday() bool {
	return c.Weekday() == time.Thursday
}

// IsFriday checks if this day is a Friday.
func (c *Time) IsFriday() bool {
	return c.Weekday() == time.Friday
}

// IsSaturday checks if this day is a Saturday.
func (c *Time) IsSaturday() bool {
	return c.Weekday() == time.Saturday
}

// IsLastWeek returns true is the date is within last week
func (c *Time) IsLastWeek() bool {
	secondsInWeek := float64(secondsInWeek)
	difference := Now().Sub(c.Time)
	if difference.Seconds() > 0 && difference.Seconds() < secondsInWeek {
		return true
	}

	return false
}

// IsLastMonth returns true is the date is within last month
func (c *Time) IsLastMonth() bool {
	now := Now()

	monthDifference := now.Month() - c.Month()

	if absValue(true, int64(monthDifference)) != 1 {
		return false
	}

	if now.UnixNano() > c.UnixNano() && monthDifference == 1 {
		return true
	}

	return false
}

// Eq determines if the current Times is equal to another
func (c *Time) Eq(timeAVA *Time) bool {
	return c.Equal(timeAVA.GetTime())
}

// EqualTo determines if the current Times is equal to another
func (c *Time) EqualTo(timeAVA *Time) bool {
	return c.Eq(timeAVA)
}

// Ne determines if the current Times is not equal to another
func (c *Time) Ne(timeAVA *Time) bool {
	return !c.Eq(timeAVA)
}

// NotEqualTo determines if the current Times is not equal to another
func (c *Time) NotEqualTo(timeAVA *Time) bool {
	return c.Ne(timeAVA)
}

// Gt determines if the current Times is greater (after) than another
func (c *Time) Gt(timeAVA *Time) bool {
	return c.After(timeAVA.GetTime())
}

// GreaterThan determines if the current Times is greater (after) than another
func (c *Time) GreaterThan(timeAVA *Time) bool {
	return c.Gt(timeAVA)
}

// Gte determines if the instance is greater (after) than or equal to another
func (c *Time) Gte(timeAVA *Time) bool {
	return c.Gt(timeAVA) || c.Eq(timeAVA)
}

// GreaterThanOrEqualTo determines if the instance is greater (after) than or equal to another
func (c *Time) GreaterThanOrEqualTo(timeAVA *Time) bool {
	return c.Gte(timeAVA) || c.Eq(timeAVA)
}

// Lt determines if the instance is less (before) than another
func (c *Time) Lt(timeAVA *Time) bool {
	return c.Before(timeAVA.GetTime())
}

// LessThan determines if the instance is less (before) than another
func (c *Time) LessThan(timeAVA *Time) bool {
	return c.Lt(timeAVA)
}

// Lte determines if the instance is less (before) or equal to another
func (c *Time) Lte(timeAVA *Time) bool {
	return c.Lt(timeAVA) || c.Eq(timeAVA)
}

// LessThanOrEqualTo determines if the instance is less (before) or equal to another
func (c *Time) LessThanOrEqualTo(timeAVA *Time) bool {
	return c.Lte(timeAVA)
}

// Between determines if the current instance is between two others
// eq Indicates if a > and < comparison should be used or <= or >=
func (c *Time) Between(a, b *Time, eq bool) bool {
	if a.Gt(b) {
		a, b = swap(a, b)
	}
	if eq {
		return c.Gte(a) && c.Lte(b)
	}

	return c.Gt(a) && c.Lt(b)
}

// Closest returns the closest date from the current time
func (c *Time) Closest(a, b *Time) *Time {
	if c.DiffInSeconds(a, true) < c.DiffInSeconds(b, true) {
		return a
	}

	return b
}

// Farthest returns the farthest date from the current time
func (c *Time) Farthest(a, b *Time) *Time {
	if c.DiffInSeconds(a, true) > c.DiffInSeconds(b, true) {
		return a
	}

	return b
}

// Min returns the minimum instance between a given instance and the current instance
func (c *Time) Min(timeAVA *Time) *Time {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}
	if c.Lt(timeAVA) {
		return c
	}

	return timeAVA
}

// Minimum returns the minimum instance between a given instance and the current instance
func (c *Time) Minimum(timeAVA *Time) *Time {
	return c.Min(timeAVA)
}

// Max returns the maximum instance between a given instance and the current instance
func (c *Time) Max(timeAVA *Time) *Time {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}

	if c.Gt(timeAVA) {
		return c
	}

	return timeAVA
}

// Maximum returns the maximum instance between a given instance and the current instance
func (c *Time) Maximum(timeAVA *Time) *Time {
	return c.Max(timeAVA)
}

// DiffInYears returns the difference in years
func (c *Time) DiffInYears(timeAVA *Time, abs bool) int64 {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}

	if c.Year() == timeAVA.Year() {
		return 0
	}

	start := NewTime(c.Time)
	end := NewTime(timeAVA.GetTime())
	if end.UnixNano() < start.UnixNano() {
		aux := start
		start = end
		end = aux
	}

	yearsAmmount := int64(end.Year()-start.Year()) - 1

	start.SetYear(end.Year())

	if start.UnixNano() <= end.UnixNano() {
		yearsAmmount++
	}

	return absValue(abs, yearsAmmount)
}

// DiffInMonths returns the difference in months
func (c *Time) DiffInMonths(timeAVA *Time, abs bool) int64 {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}

	cAux := c.Copy()
	timeAVAAux := timeAVA.Copy()
	if cAux.Location() != timeAVAAux.Location() {
		cAux = NewTime(cAux.In(time.UTC))
		timeAVAAux = NewTime(timeAVAAux.In(time.UTC))
	}

	return calculateDiffInMonths(cAux, timeAVAAux, abs)
}

func calculateDiffInMonths(c, timeAVA *Time, abs bool) int64 {
	if c.Month() == timeAVA.Month() && c.Year() == timeAVA.Year() {
		return 0
	}

	if c.Month() != timeAVA.Month() && c.Year() == timeAVA.Year() {
		diffInMonths := int64(timeAVA.Month() - c.Month())
		remainingTime := int(timeAVA.DiffInHours(c, true))

		if remainingTime < c.DaysInMonth()*hoursPerDay {
			return 0
		}

		return absValue(abs, diffInMonths)
	}

	m := monthsPerYear - c.Month() + timeAVA.Month() - 1
	if c.Year() < timeAVA.Year() && c.hasRemainingHours(timeAVA) {
		m = m + 1
	}

	if c.Year() > timeAVA.Year() {
		m = monthsPerYear - timeAVA.Month() + c.Month() - 1

		if timeAVA.hasRemainingHours(c) {
			m = m + 1
		}
	}

	diffYr := c.Year() - timeAVA.Year()
	if math.Abs(float64(diffYr)) > 1 {
		dateWithoutMonths := c.AddMonths(int(m))
		diff := dateWithoutMonths.DiffInYears(timeAVA, abs)*monthsPerYear + int64(m)

		return absValue(abs, diff)
	}

	diff := int64(m)

	if c.GreaterThan(timeAVA) {
		diff = -diff
	}

	return absValue(abs, diff)
}

func (c *Time) hasRemainingHours(timeAVA *Time) bool {
	totalHr := int64(c.DaysInMonth() * hoursPerDay)
	cHr := c.StartOfMonth().DiffInHours(c, false)
	remainHr := totalHr - cHr
	spentInHr := timeAVA.StartOfMonth().DiffInHours(timeAVA, false)

	return remainHr+spentInHr >= totalHr
}

// DiffDurationInString returns the duration difference in string format
func (c *Time) DiffDurationInString(timeAVA *Time) string {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}

	return strings.Replace(timeAVA.Sub(c.Time).String(), "-", "", 1)
}

// DiffInWeeks returns the difference in weeks
func (c *Time) DiffInWeeks(timeAVA *Time, abs bool) int64 {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}
	return c.DiffInDays(timeAVA, abs) / daysPerWeek
}

// DiffInDays returns the difference in days
func (c *Time) DiffInDays(timeAVA *Time, abs bool) int64 {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}
	return c.DiffInHours(timeAVA, abs) / hoursPerDay
}

// DiffInNights returns the difference in nights
func (c *Time) DiffInNights(timeAVA *Time, abs bool) int64 {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}
	return c.DiffInDays(timeAVA, abs)
}

// Filter represents a predicate used for filtering diffs
type Filter func(*Time) bool

// dayDuration reprensets a day in time.Duration format
const dayDuration = time.Hour * hoursPerDay

// DiffInDaysFiltered returns the difference in days using a filter
func (c *Time) DiffInDaysFiltered(f Filter, timeAVA *Time, abs bool) int64 {
	return c.DiffFiltered(dayDuration, f, timeAVA, abs)
}

// DiffInHoursFiltered returns the difference in hours using a filter
func (c *Time) DiffInHoursFiltered(f Filter, timeAVA *Time, abs bool) int64 {
	return c.DiffFiltered(time.Hour, f, timeAVA, abs)
}

// DiffInWeekdays returns the difference in weekdays
func (c *Time) DiffInWeekdays(timeAVA *Time, abs bool) int64 {
	f := func(t *Time) bool {
		return t.IsWeekday()
	}

	return c.DiffFiltered(dayDuration, f, timeAVA, abs)
}

// DiffInWeekendDays returns the difference in weekend days using a filter
func (c *Time) DiffInWeekendDays(timeAVA *Time, abs bool) int64 {
	f := func(t *Time) bool {
		return t.IsWeekend()
	}

	return c.DiffFiltered(dayDuration, f, timeAVA, abs)
}

// DiffFiltered returns the difference by the given duration using a filter
func (c *Time) DiffFiltered(duration time.Duration, f Filter, timeAVA *Time, abs bool) int64 {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}
	if c.IsSameDay(timeAVA) {
		return 0
	}

	inverse := false
	var counter int64
	s := int64(duration.Seconds())
	start, end := c.Copy(), timeAVA.Copy()
	if start.Gt(end) {
		start, end = swap(start, end)
		inverse = true
	}
	for start.DiffInSeconds(end, true)/s > 0 {
		if f(end) {
			counter++
		}
		end = NewTime(end.Add(-duration))
	}
	if inverse {
		counter = -counter
	}

	return absValue(abs, counter)
}

// DiffInHours returns the difference in hours
func (c *Time) DiffInHours(d *Time, abs bool) int64 {
	return c.DiffInMinutes(d, abs) / minutesPerHour
}

// DiffInMinutes returns the difference in minutes
func (c *Time) DiffInMinutes(d *Time, abs bool) int64 {
	return c.DiffInSeconds(d, abs) / secondsPerMinute
}

// DiffInSeconds returns the difference in seconds
func (c *Time) DiffInSeconds(timeAVA *Time, abs bool) int64 {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}
	diff := timeAVA.Timestamp() - c.Timestamp()

	return absValue(abs, diff)
}

// SecondsSinceMidnight returns the number of seconds since midnight.
func (c *Time) SecondsSinceMidnight() int {
	startOfDay := c.StartOfDay()

	return int(c.DiffInSeconds(startOfDay, true))
}

// SecondsUntilEndOfDay returns the number of seconds until 23:59:59.
func (c *Time) SecondsUntilEndOfDay() int {
	dayEnd := c.EndOfDay()

	return int(c.DiffInSeconds(dayEnd, true))
}

// absValue returns the abs value if needed
func absValue(needsAbs bool, value int64) int64 {
	if needsAbs && value < 0 {
		return -value
	}

	return value
}

func swap(a, b *Time) (*Time, *Time) {
	return b, a
}

// StartOfDay returns the time at 00:00:00 of the same day
func (c *Time) StartOfDay() *Time {
	return create(c.Year(), c.Month(), c.Day(), 0, 0, 0, 0, c.Location())
}

// EndOfDay returns the time at 23:59:59 of the same day
func (c *Time) EndOfDay() *Time {
	return create(c.Year(), c.Month(), c.Day(), 23, 59, 59, maxNSecs, c.Location())
}

// StartOfMonth returns the date on the first day of the month and the time to 00:00:00
func (c *Time) StartOfMonth() *Time {
	return create(c.Year(), c.Month(), 1, 0, 0, 0, 0, c.Location())
}

// EndOfMonth returns the date at the end of the month and time at 23:59:59
func (c *Time) EndOfMonth() *Time {
	return create(c.Year(), c.Month()+1, 0, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfQuarter returns the date at the first day of the quarter and time at 00:00:00
func (c *Time) StartOfQuarter() *Time {
	month := time.Month((c.Quarter()-1)*monthsPerQuarter + 1)

	return create(c.Year(), time.Month(month), 1, 0, 0, 0, 0, c.Location())
}

// EndOfQuarter returns the date at end of the quarter and time at 23:59:59
func (c *Time) EndOfQuarter() *Time {
	return c.StartOfQuarter().AddMonths(monthsPerQuarter - 1).EndOfMonth()
}

// StartOfYear returns the date at the first day of the year and the time at 00:00:00
func (c *Time) StartOfYear() *Time {
	return create(c.Year(), time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfYear returns the date at end of the year and time to 23:59:59
func (c *Time) EndOfYear() *Time {
	return create(c.Year(), time.December, 31, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfDecade returns the date at the first day of the decade and time at 00:00:00
func (c *Time) StartOfDecade() *Time {
	year := c.Year() - c.Year()%yearsPerDecade

	return create(year, time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfDecade returns the date at the end of the decade and time at 23:59:59
func (c *Time) EndOfDecade() *Time {
	year := c.Year() - c.Year()%yearsPerDecade + yearsPerDecade - 1

	return create(year, time.December, 31, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfCentury returns the date of the first day of the century at 00:00:00
func (c *Time) StartOfCentury() *Time {
	year := c.Year() - c.Year()%yearsPerCenturies

	return create(year, time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfCentury returns the date of the end of the century at 23:59:59
func (c *Time) EndOfCentury() *Time {
	year := c.Year() - 1 - c.Year()%yearsPerCenturies + yearsPerCenturies

	return create(year, time.December, 31, 23, 59, 59, maxNSecs, c.Location())
}

// StartOfWeek returns the date of the first day of week at 00:00:00
func (c *Time) StartOfWeek() *Time {
	if c.Weekday() == c.WeekStartsAt() {
		return c.StartOfDay()
	}

	return c.Previous(c.WeekStartsAt())
}

// EndOfWeek returns the date of the last day of the week at 23:59:59
func (c *Time) EndOfWeek() *Time {
	return c.Next(c.WeekEndsAt()).EndOfDay()
}

// Results changes the time to the next occurrence of a given day of the week
func (c *Time) Next(wd time.Weekday) *Time {
	var temp = c
	temp = temp.AddDay()
	for temp.Weekday() != wd {
		temp = temp.AddDay()
	}
	c.Time = temp.GetTime()
	return c.StartOfDay()
}

// NextWeekday goes forward to the next weekday
func (c *Time) NextWeekday() *Time {
	return c.AddWeekday()
}

// PreviousWeekday goes back to the previous weekday
func (c *Time) PreviousWeekday() *Time {
	return c.SubWeekday()
}

// NextWeekendDay goes forward to the next weekend day
func (c *Time) NextWeekendDay() *Time {
	var temp = c
	temp = temp.AddDay()
	for !temp.IsWeekend() {
		temp = temp.AddDay()
	}
	c.Time = temp.GetTime()
	return c
}

// PreviousWeekendDay goes back to the previous weekend day
func (c *Time) PreviousWeekendDay() *Time {
	var temp = c
	temp = temp.SubDay()
	for !temp.IsWeekend() {
		temp = temp.SubDay()
	}
	c.Time = temp.GetTime()
	return c
}

// Previous changes the time to the previous occurrence of a given day of the week
func (c *Time) Previous(wd time.Weekday) *Time {
	var temp = c
	temp = temp.SubDay()
	for temp.Weekday() != wd {
		temp = temp.SubDay()
	}
	c.Time = temp.GetTime()
	return c.StartOfDay()
}

// FirstOfMonth returns the first occurrence of a given day of the week in the current month
func (c *Time) FirstOfMonth(wd time.Weekday) *Time {
	d := c.StartOfMonth()
	if d.Weekday() != wd {
		return d.Next(wd)
	}

	return d
}

// LastOfMonth returns the last occurrence of a given day of the week in the current month
func (c *Time) LastOfMonth(wd time.Weekday) *Time {
	d := c.EndOfMonth()
	if d.Weekday() != wd {
		return d.Previous(wd)
	}

	return d.StartOfDay()
}

// LastDayOfMonth returns a new Times instance with the last day of current month
func (c *Time) LastDayOfMonth() *Time {
	return NewTime(time.Date(c.Year(), c.Month(), c.DaysInMonth(), 0, 0, 0, 0, time.UTC))
}

// FirstDayOfMonth returns a new Times instance with the first day of current month
func (c *Time) FirstDayOfMonth() *Time {
	return NewTime(time.Date(c.Year(), c.Month(), 1, 0, 0, 0, 0, time.UTC))
}

// NthOfMonth returns the given occurrence of a given day of the week in the current month
// If the calculated occurrence is outside the scope of current month, no modifications are made
func (c *Time) NthOfMonth(nth int, wd time.Weekday) *Time {
	cp := c.Copy().StartOfMonth()
	i := 0
	if cp.Weekday() == wd {
		i++
	}
	for i < nth {
		cp = cp.Next(wd)
		i++
	}
	if cp.Gt(c.EndOfMonth()) {
		return c
	}

	return cp
}

// FirstOfQuarter returns the first occurrence of a given day of the week in the current quarter
func (c *Time) FirstOfQuarter(wd time.Weekday) *Time {
	d := c.StartOfQuarter()
	if d.Weekday() != wd {
		return d.Next(wd)
	}

	return d
}

// LastOfQuarter returns the last occurrence of a given day of the week in the current quarter
func (c *Time) LastOfQuarter(wd time.Weekday) *Time {
	d := c.EndOfQuarter()
	if d.Weekday() != wd {
		return d.Previous(wd)
	}

	return d.StartOfDay()
}

// NthOfQuarter returns the given occurrence of a given day of the week in the current quarter
// If the calculated occurrence is outside the scope of current quarter, no modifications are made
func (c *Time) NthOfQuarter(nth int, wd time.Weekday) *Time {
	cp := c.Copy().StartOfQuarter()
	i := 0
	if cp.Weekday() == wd {
		i++
	}
	for i < nth {
		cp = cp.Next(wd)
		i++
	}
	if cp.Gt(c.EndOfQuarter()) {
		return c
	}

	return cp
}

// FirstOfYear returns the first occurrence of a given day of the week in the current year
func (c *Time) FirstOfYear(wd time.Weekday) *Time {
	d := c.StartOfYear()
	if d.Weekday() != wd {
		return d.Next(wd)
	}

	return d
}

// LastOfYear returns the last occurrence of a given day of the week in the current year
func (c *Time) LastOfYear(wd time.Weekday) *Time {
	d := c.EndOfYear()
	if d.Weekday() != wd {
		return d.Previous(wd)
	}

	return d.StartOfDay()
}

// NthOfYear returns the given occurrence of a given day of the week in the current year
// If the calculated occurrence is outside the scope of current year, no modifications are made
func (c *Time) NthOfYear(nth int, wd time.Weekday) *Time {
	cp := c.Copy().StartOfYear()
	i := 0
	if cp.Weekday() == wd {
		i++
	}
	for i < nth {
		cp = cp.Next(wd)
		i++
	}
	if cp.Gt(c.EndOfYear()) {
		return c
	}

	return cp
}

// Average returns the average between a given Times date and the current date
func (c *Time) Average(timeAVA *Time) *Time {
	if timeAVA == nil {
		timeAVA = nowIn(c.Location())
	}
	if c.Eq(timeAVA) {
		return c.Copy()
	}
	average := int(c.DiffInSeconds(timeAVA, false) / 2)

	return c.AddSeconds(average)
}

// Clock returns the hour, minute, and second within the day specified by t.
func (c *Time) Clock() (hour, min, sec int) {
	return c.Time.Clock()
}

func (c *Time) Nanosecond() int {
	return c.Time.Nanosecond()
}

func (c *Time) Date() (year int, month time.Month, day int) {
	return c.Time.Date()
}

func (c *Time) In(loc *time.Location) time.Time {
	return c.Time.In(loc)
}

// GetMillis is a convenience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetMillisForTime is a convenience method to get milliseconds since epoch for provided Times.
func GetMillisForTime(thisTime time.Time) int64 {
	return thisTime.UnixNano() / int64(time.Millisecond)
}

// PadDateStringZeros is a convenience method to pad 2 digit date parts with zeros to meet ISO 8601 format
func PadDateStringZeros(dateString string) string {
	parts := strings.Split(dateString, "-")
	for index, part := range parts {
		if len(part) == 1 {
			parts[index] = "0" + part
		}
	}
	dateString = strings.Join(parts[:], "-")
	return dateString
}

// GetStartOfDayMillis is a convenience method to get milliseconds since epoch for provided date's start of day
func GetStartOfDayMillis(thisTime time.Time, timeZoneOffset int) int64 {
	localSearchTimeZone := time.FixedZone("Local Search Times Zone", timeZoneOffset)
	resultTime := time.Date(thisTime.Year(), thisTime.Month(), thisTime.Day(), 0, 0, 0, 0, localSearchTimeZone)
	return GetMillisForTime(resultTime)
}

// GetEndOfDayMillis is a convenience method to get milliseconds since epoch for provided date's end of day
func GetEndOfDayMillis(thisTime time.Time, timeZoneOffset int) int64 {
	localSearchTimeZone := time.FixedZone("Local Search Times Zone", timeZoneOffset)
	resultTime := time.Date(thisTime.Year(), thisTime.Month(), thisTime.Day(), 23, 59, 59, 999999999, localSearchTimeZone)
	return GetMillisForTime(resultTime)
}

func ParseDuration(v string) time.Duration {
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0
	}
	return d
}

func NowStr() string {
	return time.Now().Format("2006-01-02T15:04:05.000")
}
