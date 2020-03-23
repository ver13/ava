//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package time

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type TimeI interface {
	Copy() *Time
	WeekStartsAt() time.Weekday
	WeekEndsAt() time.Weekday
	WeekendDays() []time.Weekday
	Quarter() int
	Age() int
	DaysInMonth() int
	DaysInYear() int
	WeekOfMonth() int
	WeekOfYear() (int, int)
	TimeZone() string
	Timestamp() int64
	String() string
	AddYears(y int) *Time
	AddYear() *Time
	AddQuarters(q int) *Time
	AddQuarter() *Time
	AddCenturies(cent int) *Time
	AddCentury() *Time
	AddMonths(m int) *Time
	AddMonth() *Time
	AddSeconds(s int) *Time
	AddSecond() *Time
	AddDays(d int) *Time
	AddDay() *Time
	AddWeekdays(wd int) *Time
	AddWeekday() *Time
	AddWeeks(w int) *Time
	AddWeek() *Time
	AddHours(h int) *Time
	AddHour() *Time
	AddMonthsNoOverflow(m int) *Time
	PreviousMonthLastDay() *Time
	AddMonthNoOverflow() *Time
	AddMinutes(m int) *Time
	AddMinute() *Time
	SubYear() *Time
	SubYears(y int) *Time
	SubQuarter() *Time
	SubQuarters(q int) *Time
	SubCentury() *Time
	SubCenturies(cent int) *Time
	SubMonth() *Time
	SubMonths(m int) *Time
	SubMonthNoOverflow() *Time
	SubMonthsNoOverflow(m int) *Time
	SubDay() *Time
	SubDays(d int) *Time
	SubWeekday() *Time
	SubWeekdays(wd int) *Time
	SubWeek() *Time
	SubWeeks(w int) *Time
	SubHour() *Time
	SubHours(h int) *Time
	SubMinute() *Time
	SubMinutes(m int) *Time
	SubSecond() *Time
	SubSeconds(s int) *Time
	SetYear(y int)
	SetMonth(m time.Month)
	SetDay(d int)
	SetHour(h int)
	SetMinute(m int)
	SetSecond(s int)
	SetDate(y int, m time.Month, d int)
	SetDateTime(y int, mon time.Month, d, h, m, s int)
	SetTimeFromTimeString(timeString string) *errorAVA.Error
	SetWeekEndsAt(wd time.Weekday)
	SetWeekStartsAt(wd time.Weekday)
	SetWeekendDays(wds []time.Weekday)
	SetTimestamp(sec int64)
	SetTimeZone(name string) *errorAVA.Error
	ResetStringFormat()
	SetStringFormat(format string)
	DateString() string
	FormattedDateString() string
	TimeString() string
	DateTimeString() string
	DayDateTimeString() string
	AtomString() string
	CookieString() string
	ISO8601String() string
	RFC822String() string
	RFC850String() string
	RFC1036String() string
	RFC1123String() string
	RFC2822String() string
	RFC3339String() string
	RSSString() string
	W3CString() string
	IsWeekday() bool
	IsWeekend() bool
	IsYesterday() bool
	IsToday() bool
	IsTomorrow() bool
	IsFuture() bool
	IsPast() bool
	IsLeapYear() bool
	IsLongYear() bool
	IsSameAs(format string, time *Time) bool
	IsCurrentYear() bool
	IsSameYear(time *Time) bool
	IsCurrentMonth() bool
	IsSameMonth(time *Time, sameYear bool) bool
	IsSameDay(time *Time) bool
	IsSunday() bool
	IsMonday() bool
	IsTuesday() bool
	IsWednesday() bool
	IsThursday() bool
	IsFriday() bool
	IsSaturday() bool
	IsLastWeek() bool
	IsLastMonth() bool
	Eq(time *Time) bool
	EqualTo(time *Time) bool
	Ne(time *Time) bool
	NotEqualTo(time *Time) bool
	Gt(time *Time) bool
	GreaterThan(time *Time) bool
	Gte(time *Time) bool
	GreaterThanOrEqualTo(time *Time) bool
	Lt(time *Time) bool
	LessThan(time *Time) bool
	Lte(time *Time) bool
	LessThanOrEqualTo(time *Time) bool
	Between(a, b *Time, eq bool) bool
	Closest(a, b *Time) *Time
	Farthest(a, b *Time) *Time
	Min(time *Time) *Time
	Minimum(time *Time) *Time
	Max(time *Time) *Time
	Maximum(time *Time) *Time
	DiffInYears(time *Time, abs bool) int64
	DiffInMonths(time *Time, abs bool) int64
	hasRemainingHours(time *Time) bool
	DiffDurationInString(time *Time) string
	DiffInWeeks(time *Time, abs bool) int64
	DiffInDays(time *Time, abs bool) int64
	DiffInNights(time *Time, abs bool) int64
	DiffInDaysFiltered(f Filter, time *Time, abs bool) int64
	DiffInHoursFiltered(f Filter, time *Time, abs bool) int64
	DiffInWeekdays(time *Time, abs bool) int64
	DiffInWeekendDays(time *Time, abs bool) int64
	DiffFiltered(duration time.Duration, f Filter, time *Time, abs bool) int64
	DiffInHours(d *Time, abs bool) int64
	DiffInMinutes(d *Time, abs bool) int64
	DiffInSeconds(time *Time, abs bool) int64
	SecondsSinceMidnight() int
	SecondsUntilEndOfDay() int
	StartOfDay() *Time
	EndOfDay() *Time
	StartOfMonth() *Time
	EndOfMonth() *Time
	StartOfQuarter() *Time
	EndOfQuarter() *Time
	StartOfYear() *Time
	EndOfYear() *Time
	StartOfDecade() *Time
	EndOfDecade() *Time
	StartOfCentury() *Time
	EndOfCentury() *Time
	StartOfWeek() *Time
	EndOfWeek() *Time
	Next(wd time.Weekday) *Time
	NextWeekday() *Time
	PreviousWeekday() *Time
	NextWeekendDay() *Time
	PreviousWeekendDay() *Time
	Previous(wd time.Weekday) *Time
	FirstOfMonth(wd time.Weekday) *Time
	LastOfMonth(wd time.Weekday) *Time
	LastDayOfMonth() *Time
	FirstDayOfMonth() *Time
	NthOfMonth(nth int, wd time.Weekday) *Time
	FirstOfQuarter(wd time.Weekday) *Time
	LastOfQuarter(wd time.Weekday) *Time
	NthOfQuarter(nth int, wd time.Weekday) *Time
	FirstOfYear(wd time.Weekday) *Time
	LastOfYear(wd time.Weekday) *Time
	NthOfYear(nth int, wd time.Weekday) *Time
	Average(time *Time) *Time
	Clock() (hour, min, sec int)
	Nanosecond() int
	Date() (year int, month time.Month, day int)
	In(loc *time.Location) time.Time
	Day() int
	Month() time.Month
	Year() int
	After(u time.Time) bool
	GetTime() time.Time
	Format(layout string) string
	Sub(u time.Time) time.Duration
	UnixNano() int64
	Location() *time.Location
	Add(d time.Duration) time.Time
	Weekday() time.Weekday
}
