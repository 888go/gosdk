package time

// ParseError describes a problem parsing a time string.
type ParseError struct {

} //md5:6c5710d3aad84515


const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
)// md5:0dd08e9a1e224ec7




const (
	_                        = iota
	stdLongMonth             = iota + stdNeedDate  // "January"
	stdMonth                                       // "Jan"
	stdNumMonth                                    // "1"
	stdZeroMonth                                   // "01"
	stdLongWeekDay                                 // "Monday"
	stdWeekDay                                     // "Mon"
	stdDay                                         // "2"
	stdUnderDay                                    // "_2"
	stdZeroDay                                     // "02"
	stdUnderYearDay                                // "__2"
	stdZeroYearDay                                 // "002"
	stdHour                  = iota + stdNeedClock // "15"
	stdHour12                                      // "3"
	stdZeroHour12                                  // "03"
	stdMinute                                      // "4"
	stdZeroMinute                                  // "04"
	stdSecond                                      // "5"
	stdZeroSecond                                  // "05"
	stdLongYear              = iota + stdNeedDate  // "2006"
	stdYear                                        // "06"
	stdPM                    = iota + stdNeedClock // "PM"
	stdpm                                          // "pm"
	stdTZ                    = iota                // "MST"
	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
	stdISO8601SecondsTZ                            // "Z070000"
	stdISO8601ShortTZ                              // "Z07"
	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
	stdNumTZ                                       // "-0700"  // always numeric
	stdNumSecondsTz                                // "-070000"
	stdNumShortTZ                                  // "-07"    // always numeric
	stdNumColonTZ                                  // "-07:00" // always numeric
	stdNumColonSecondsTZ                           // "-07:00:00"
	stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
	stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted

	stdNeedDate       = 1 << 8             // need month, day, year
	stdNeedClock      = 2 << 8             // need hour, minute, second
	stdArgShift       = 16                 // extra argument in high bits, above low stdArgShift
	stdSeparatorShift = 28                 // extra argument in high 4 bits for fractional second separators
	stdMask           = 1<<stdArgShift - 1 // mask out argument
)// md5:30d8e7706ca83b6b




const (
	lowerhex  = "0123456789abcdef"
	runeSelf  = 0x80
	runeError = '\uFFFD'
)// md5:66ebe9c952b8908f



// String returns the time formatted using the format string
//
//	"2006-01-02 15:04:05.999999999 -0700 MST"
//
// If the time has a monotonic clock reading, the returned string
// includes a final field "m=±<value>", where value is the monotonic
// clock reading formatted as a decimal number of seconds.
//
// The returned string is meant for debugging; for a stable serialized
// representation, use t.MarshalText, t.MarshalBinary, or t.Format
// with an explicit format string.
func (t Time) String() string { //md5:244060d240d2fba5

}

// GoString implements fmt.GoStringer and formats t to be printed in Go source
// code.
func (t Time) GoString() string { //md5:bbef06c0bba983b6

}

// Format returns a textual representation of the time value formatted according
// to the layout defined by the argument. See the documentation for the
// constant called Layout to see how to represent the layout format.
//
// The executable example for Time.Format demonstrates the working
// of the layout string in detail and is a good reference.
func (t Time) Format(layout string) string { //md5:cd21a0becf8aba8a

}

// AppendFormat is like Format but appends the textual
// representation to b and returns the extended buffer.
func (t Time) AppendFormat(b []byte, layout string) []byte { //md5:f1300a966a1e3949

}

// Error returns the string representation of a ParseError.
func (e *ParseError) Error() string { //md5:c2cf7f1bf647d37a

}

// Parse parses a formatted string and returns the time value it represents.
// See the documentation for the constant called Layout to see how to
// represent the format. The second argument must be parseable using
// the format string (layout) provided as the first argument.
//
// The example for Time.Format demonstrates the working of the layout string
// in detail and is a good reference.
//
// When parsing (only), the input may contain a fractional second
// field immediately after the seconds field, even if the layout does not
// signify its presence. In that case either a comma or a decimal point
// followed by a maximal series of digits is parsed as a fractional second.
// Fractional seconds are truncated to nanosecond precision.
//
// Elements omitted from the layout are assumed to be zero or, when
// zero is impossible, one, so parsing "3:04pm" returns the time
// corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
// 0, this time is before the zero Time).
// Years must be in the range 0000..9999. The day of the week is checked
// for syntax but it is otherwise ignored.
//
// For layouts specifying the two-digit year 06, a value NN >= 69 will be treated
// as 19NN and a value NN < 69 will be treated as 20NN.
//
// The remainder of this comment describes the handling of time zones.
//
// In the absence of a time zone indicator, Parse returns a time in UTC.
//
// When parsing a time with a zone offset like -0700, if the offset corresponds
// to a time zone used by the current location (Local), then Parse uses that
// location and zone in the returned time. Otherwise it records the time as
// being in a fabricated location with time fixed at the given zone offset.
//
// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
// has a defined offset in the current location, then that offset is used.
// The zone abbreviation "UTC" is recognized as UTC regardless of location.
// If the zone abbreviation is unknown, Parse records the time as being
// in a fabricated location with the given zone abbreviation and a zero offset.
// This choice means that such a time can be parsed and reformatted with the
// same layout losslessly, but the exact instant used in the representation will
// differ by the actual zone offset. To avoid such problems, prefer time layouts
// that use a numeric zone offset, or use ParseInLocation.
func Parse(layout, value string) (Time, error) { //md5:04f0e938ce9833d8

}

// ParseInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets a time as UTC;
// ParseInLocation interprets the time as in the given location.
// Second, when given a zone offset or abbreviation, Parse tries to match it
// against the Local location; ParseInLocation uses the given location.
func ParseInLocation(layout, value string, loc *Location) (Time, error) { //md5:72c26b70b1a8980b

}

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ParseDuration(s string) (Duration, error) { //md5:2817ced6ba396e3a

}