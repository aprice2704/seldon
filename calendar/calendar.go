package calendar

import (
	"fmt"
	"time"
)

const ShortDate = "02-Jan-2006"
const ShortDateTime = "02-Jan-2006 15:04"

// Slot is the number of a time period in a discrete calendar
//type Slot int

type Interval struct {
	Start, End time.Time
}

// CustomUsable defines universal calendar behaviour
type CustomUsable interface {
	SetUnusable(s, e time.Time)
	SetUsable(s, e time.Time)
}

type FindUsable interface {
	IsAllUsable(s, e time.Time) bool
	UsableIntervals(s, e time.Time) []Interval
}

type CalendarFunc func(i Interval) func() []Interval

type Calendar struct {
	Layers []CalendarFunc
}

func (c *CalLayer) GetUsableIntervals(s, e time.Time) []Interval {
	return nil
}

func (c *CalLayer) IsAllUsable(s, e time.Time) bool {
	return s.After(c.Start) && e.Before(c.End)
}

type CalLayer struct {
	Start time.Time
	End   time.Time
}

func (c *CalLayer) String() string {
	return fmt.Sprintf("Layer, Starts: %s Ends: %s", c.Start.Format(ShortDate), c.End.Format(ShortDate))
}

func NewCalLayer(s, e time.Time) *CalLayer {
	d := new(CalLayer)
	d.Start = s
	d.End = e
	return d
}

// Resolution is the minimum length of time which this calendar works on -- its slot length
//type Resolution int

// const (
// 	Second Resolution = iota
// 	Minute
// 	Hour
// 	Day
// 	Week
// 	Fortnite   // 14 days
// 	Month      // Calendar month
// 	Fourweek   // 13 periods/year
// 	Quarter    // 3 calendar months
// 	Shift      // 8 hrs
// 	CustomDays // Any whole number of days
// 	Custom     // Any time.Duration
// )
//
// // CanTimeType is a value indicating whether a calendar is quantized (e.g. deals only in entire days) or continuous,
// // in which case it uses full golang resolution and usable times must be built explicitly rather than by bulk. It is
// // expected that quantized calendars will be faster and more efficient to store, especially when they apply to long periods
// // Procedural calendars have no slot specific storage, they return usable time based entirely on computations.
// //
// type CalTimeType int
//
// const (
// 	Discrete CalTimeType = iota
// 	Procedural
// 	Continuous
// )
//
// func (i *Interval) NewInterval(s, e time.Time) *Interval {
// 	i.Start = s
// 	i.End = e
// }
//
// type Calendar struct {
// 	Start         time.Time
// 	End           time.Time
// }
//
// func NewCalendar(options ...func(*)) *Calendar {
// 	p := new(Calendar)
// 	p.Options(SetResolution(Day), SetDefUsable(true))
// 	return p
// }
//
// // DoWUsable sets all the days of the week given to usable for the duration of the calendar
// func (c *Calendar) DoWUsable(time.Weekday) {
//
// }
//
// func (c *Calendar) SlotLength() time.Duration {
// 	return c.slotlength
// }
//
// // Rob Pike's crazy options-setting pattern: https://commandcenter.blogspot.ca/2014/01/self-referential-functions-and-design.html
// type caloption func(*Calendar)
//
// func (c *Calendar) Options(options ...caloption) {
// 	for _, opt := range options {
// 		opt(c)
// 	}
// }
//
// // Resolution returns a func that sets Resolution, use with c.Options( calendar.SetResolution(calendar.Day))
// func SetResolution(r Resolution) caloption {
// 	return func(c *Calendar) {
// 		c.resolution = r
// 	}
// }
//
// func SetDefUsable(u bool) caloption {
// 	return func(c *Calendar) {
// 		c.DefaultUsable = u
// 	}
// }
