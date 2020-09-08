package moment

import "time"

const (
	// HoursPerDay specifies the number of hours in a day
	HoursPerDay = 24
	// MinutesPerHour specifies the number of minutes in an hour
	MinutesPerHour = 60
	// SecondsPerMinute specifies the number of seconds in a minute
	SecondsPerMinute = 60
)

// Point defines an abstract point in time. It does not include a day, month, or year but simply
// a time of day
type Point struct {
	hour       int
	minute     int
	second     int
	nanoSecond int
	location   *time.Location
}

// NewPoint creates a new time point with the given arguments. If no
// arguments are given, the returned point represents the time 00:00 UTC. You can
// provide up to 4 arguments in the order of "hour", "minute", "second", and "nanosecond".
// Providing a fifth argument results in the default behavior and returns 00:00 UTC.
func NewPoint(args ...int) Point {
	p := Point{
		location: time.Local,
	}

	switch len(args) {
	case 4:
		p.nanoSecond = args[3]
		fallthrough
	case 3:
		p.SetSecond(args[2])
		fallthrough
	case 2:
		p.SetMinute(args[1])
		fallthrough
	case 1:
		p.SetHour(args[0])
	}
	return p
}

// SetLocation sets the point location
func (p Point) SetLocation(loc *time.Location) {
	if loc == nil {
		// do nothing!
		return
	}
	p.location = loc
}

// SetSecond checks to ensure the given value is valid and then sets the "second" parameter
func (p Point) SetSecond(sec int) {
	if sec < 0 || sec >= SecondsPerMinute {
		return
	}
	p.second = sec
}

// SetMinute checks to ensure the given value is valid and then sets the "minute" parameter
func (p Point) SetMinute(min int) {
	if min < 0 || min >= MinutesPerHour {
		return
	}
	p.minute = min
}

// SetHour checks to ensure the given value is valid and then sets the "hour" parameter
func (p Point) SetHour(hr int) {
	if hr < 0 || hr >= HoursPerDay {
		return
	}
	p.hour = hr
}

// On returns the concrete time that the point would occur on the day given
func (p Point) On(day time.Time) time.Time {
	if p.location == nil {
		p.location = time.UTC
	}
	return time.Date(day.Year(), day.Month(), day.Day(), p.hour, p.minute, p.second, p.nanoSecond, p.location)
}

// Span defines a duration of time starting at an abstract moment in time
type Span struct {
	begin  Point
	length time.Duration
}

// NewSpan creates a Span using the given inputs
func NewSpan(begin Point, len time.Duration) Span {
	s := Span{
		begin:  begin,
		length: len,
	}
	return s
}

// Start returns the "real" start time of a Span on the given day
func (s Span) Start(day time.Time) time.Time {
	return s.begin.On(day)
}

// End returns the "real" end time of a Span on the given day
func (s Span) End(day time.Time) time.Time {
	return s.Start(day).Add(s.length)
}
