package time

// A Time represents an instant in time with nanosecond precision.
//
// Programs using times should typically store and pass them as values,
// not pointers. That is, time variables and struct fields should be of
// type time.Time, not *time.Time.
//
// A Time value can be used by multiple goroutines simultaneously except
// that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
// UnmarshalText are not concurrency-safe.
//
// Time instants can be compared using the Before, After, and Equal methods.
// The Sub method subtracts two instants, producing a Duration.
// The Add method adds a Time and a Duration, producing a Time.
//
// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
// As this time is unlikely to come up in practice, the IsZero method gives
// a simple way of detecting a time that has not been initialized explicitly.
//
// Each time has an associated Location. The methods Local, UTC, and In return a
// Time with a specific Location. Changing the Location of a Time value with
// these methods does not change the actual instant it represents, only the time
// zone in which to interpret it.
//
// Representations of a Time value saved by the GobEncode, MarshalBinary,
// MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
// the location name. They therefore lose information about Daylight Saving Time.
//
// In addition to the required “wall clock” reading, a Time may contain an optional
// reading of the current process's monotonic clock, to provide additional precision
// for comparison or subtraction.
// See the “Monotonic Clocks” section in the package documentation for details.
//
// Note that the Go == operator compares not just the time instant but also the
// Location and the monotonic clock reading. Therefore, Time values should not
// be used as map or database keys without first guaranteeing that the
// identical Location has been set for all values, which can be achieved
// through use of the UTC or Local method, and that the monotonic clock reading
// has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
// to t == u, since t.Equal uses the most accurate comparison available and
// correctly handles the case when only one of its arguments has a monotonic
// clock reading.
type Time struct {

} //md5:a8f1e158c36dbc97


const (
	hasMonotonic = 1 << 63
	maxWall      = wallToInternal + (1<<33 - 1) // year 2157
	minWall      = wallToInternal               // year 1885
	nsecMask     = 1<<30 - 1
	nsecShift    = 30
)// md5:84fc15a2be736955




const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)// md5:147c9abc849d3405




const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)// md5:9efbd1538cfcbf96




const (
	// The unsigned zero year for internal calculations.
	// Must be 1 mod 400, and times before it will not compute correctly,
	// but otherwise can be changed at will.
	absoluteZeroYear = -292277022399

	// The year of the zero Time.
	// Assumed by the unixToInternal computation below.
	internalYear = 1

	// Offsets to convert between internal and absolute or Unix times.
	absoluteToInternal int64 = (absoluteZeroYear - internalYear) * 365.2425 * secondsPerDay
	internalToAbsolute       = -absoluteToInternal

	unixToInternal int64 = (1969*365 + 1969/4 - 1969/100 + 1969/400) * secondsPerDay
	internalToUnix int64 = -unixToInternal

	wallToInternal int64 = (1884*365 + 1884/4 - 1884/100 + 1884/400) * secondsPerDay
)// md5:f6e0b667fa556e68




const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)// md5:7fa4af1b9114d2d1




const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)// md5:a064d6044cbb28fe




const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
	secondsPerDay    = 24 * secondsPerHour
	secondsPerWeek   = 7 * secondsPerDay
	daysPer400Years  = 365*400 + 97
	daysPer100Years  = 365*100 + 24
	daysPer4Years    = 365*4 + 1
)// md5:c9382aa097574f5d




const (
	timeBinaryVersionV1 byte = iota + 1 // For general situation
	timeBinaryVersionV2                 // For LMT only
)// md5:7df46a836c46218d



// After reports whether the time instant t is after u.
func (t Time) After(u Time) bool { //md5:5dfb6db123b2d436

}

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool { //md5:d098b27691e377d5

}

// Compare compares the time instant t with u. If t is before u, it returns -1;
// if t is after u, it returns +1; if they're the same, it returns 0.
func (t Time) Compare(u Time) int { //md5:35c931a4ac1e3dea

}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
func (t Time) Equal(u Time) bool { //md5:02203c2b950ebe2d

}

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string { //md5:014a3380bef17d93

}

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Weekday) String() string { //md5:3aa0ba35fd23d3a0

}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t Time) IsZero() bool { //md5:fd95a0aa1783c81e

}

// Date returns the year, month, and day in which t occurs.
func (t Time) Date() (year int, month Month, day int) { //md5:a5a774f5fec8bb67

}

// Year returns the year in which t occurs.
func (t Time) Year() int { //md5:01bea20ebc40f05e

}

// Month returns the month of the year specified by t.
func (t Time) Month() Month { //md5:12e0562e3bfdd479

}

// Day returns the day of the month specified by t.
func (t Time) Day() int { //md5:0b0c2cb83d99f883

}

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() Weekday { //md5:0eb77d8496ff6452

}

// ISOWeek returns the ISO 8601 year and week number in which t occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (t Time) ISOWeek() (year, week int) { //md5:720dde5fce39726b

}

// Clock returns the hour, minute, and second within the day specified by t.
func (t Time) Clock() (hour, min, sec int) { //md5:b31ce25ccbb0b1b8

}

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (t Time) Hour() int { //md5:efa7745ae50c29d5

}

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (t Time) Minute() int { //md5:c0dd1485a43bcd50

}

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (t Time) Second() int { //md5:e2014d911e6cd34e

}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (t Time) Nanosecond() int { //md5:89292d986d7e3147

}

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (t Time) YearDay() int { //md5:0b29732a13a1b75a

}

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0s.
func (d Duration) String() string { //md5:7e1f3f8f34886861

}

// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64 { //md5:31b7144fa28ceb5b

}

// Microseconds returns the duration as an integer microsecond count.
func (d Duration) Microseconds() int64 { //md5:d4942111d046902a

}

// Milliseconds returns the duration as an integer millisecond count.
func (d Duration) Milliseconds() int64 { //md5:6cc34c02dbb54bc2

}

// Seconds returns the duration as a floating point number of seconds.
func (d Duration) Seconds() float64 { //md5:75def545f7232495

}

// Minutes returns the duration as a floating point number of minutes.
func (d Duration) Minutes() float64 { //md5:900bce301a35981f

}

// Hours returns the duration as a floating point number of hours.
func (d Duration) Hours() float64 { //md5:d5ebe85a747f47bc

}

// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
func (d Duration) Truncate(m Duration) Duration { //md5:fe65fdff50fff8d6

}

// Round returns the result of rounding d to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration,
// Round returns the maximum (or minimum) duration.
// If m <= 0, Round returns d unchanged.
func (d Duration) Round(m Duration) Duration { //md5:2d3f614e4b10efeb

}

// Abs returns the absolute value of d.
// As a special case, math.MinInt64 is converted to math.MaxInt64.
func (d Duration) Abs() Duration { //md5:ce1eb65bdaa80e6d

}

// Add returns the time t+d.
func (t Time) Add(d Duration) Time { //md5:8720c0d0de353e9c

}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
func (t Time) Sub(u Time) Duration { //md5:df77f9b3e6bf5ada

}

// Since returns the time elapsed since t.
// It is shorthand for time.Now().Sub(t).
func Since(t Time) Duration { //md5:810dd8f30988f658

}

// Until returns the duration until t.
// It is shorthand for t.Sub(time.Now()).
func Until(t Time) Duration { //md5:63b7023bdca00f52

}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// Note that dates are fundamentally coupled to timezones, and calendrical
// periods like days don't have fixed durations. AddDate uses the Location of
// the Time value to determine these durations. That means that the same
// AddDate arguments can produce a different shift in absolute time depending on
// the base Time value and its Location. For example, AddDate(0, 0, 1) applied
// to 12:00 on March 27 always returns 12:00 on March 28. At some locations and
// in some years this is a 24 hour shift. In others it's a 23 hour shift due to
// daylight savings time transitions.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
func (t Time) AddDate(years int, months int, days int) Time { //md5:f604377cb448bd9a

}

// Now returns the current local time.
func Now() Time { //md5:b2c3ce5c3a0dddd8

}

// UTC returns t with the location set to UTC.
func (t Time) UTC() Time { //md5:dfc0ea55e77bdd85

}

// Local returns t with the location set to local time.
func (t Time) Local() Time { //md5:773d045186c2d5e6

}

// In returns a copy of t representing the same time instant, but
// with the copy's location information set to loc for display
// purposes.
//
// In panics if loc is nil.
func (t Time) In(loc *Location) Time { //md5:cdf90b2adf199e29

}

// Location returns the time zone information associated with t.
func (t Time) Location() *Location { //md5:320904942ebf3a70

}

// Zone computes the time zone in effect at time t, returning the abbreviated
// name of the zone (such as "CET") and its offset in seconds east of UTC.
func (t Time) Zone() (name string, offset int) { //md5:f6bf762564bdb7bd

}

// ZoneBounds returns the bounds of the time zone in effect at time t.
// The zone begins at start and the next zone begins at end.
// If the zone begins at the beginning of time, start will be returned as a zero Time.
// If the zone goes on forever, end will be returned as a zero Time.
// The Location of the returned times will be the same as t.
func (t Time) ZoneBounds() (start, end Time) { //md5:764bd854389bfd7e

}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
// Unix-like operating systems often record time as a 32-bit
// count of seconds, but since the method here returns a 64-bit
// value it is valid for billions of years into the past or future.
func (t Time) Unix() int64 { //md5:708a59c5afafb133

}

// UnixMilli returns t as a Unix time, the number of milliseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// milliseconds cannot be represented by an int64 (a date more than 292 million
// years before or after 1970). The result does not depend on the
// location associated with t.
func (t Time) UnixMilli() int64 { //md5:8ad355b089164e81

}

// UnixMicro returns t as a Unix time, the number of microseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// microseconds cannot be represented by an int64 (a date before year -290307 or
// after year 294246). The result does not depend on the location associated
// with t.
func (t Time) UnixMicro() int64 { //md5:df74c29f1d20ba8a

}

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC. The result is undefined if the Unix time
// in nanoseconds cannot be represented by an int64 (a date before the year
// 1678 or after 2262). Note that this means the result of calling UnixNano
// on the zero Time is undefined. The result does not depend on the
// location associated with t.
func (t Time) UnixNano() int64 { //md5:e37c334610d19b98

}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t Time) MarshalBinary() ([]byte, error) { //md5:da3dea82d6dc9326

}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (t *Time) UnmarshalBinary(data []byte) error { //md5:a4a2ca872058057e

}

// GobEncode implements the gob.GobEncoder interface.
func (t Time) GobEncode() ([]byte, error) { //md5:bd28fa55c2574953

}

// GobDecode implements the gob.GobDecoder interface.
func (t *Time) GobDecode(data []byte) error { //md5:f1c31f760b2eb0ef

}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in the RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
func (t Time) MarshalJSON() ([]byte, error) { //md5:45b1e08736011f81

}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time must be a quoted string in the RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error { //md5:471a9f0fcd628553

}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
func (t Time) MarshalText() ([]byte, error) { //md5:53833c0a151d9869

}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time must be in the RFC 3339 format.
func (t *Time) UnmarshalText(data []byte) error { //md5:ba75915cc097f6f0

}

// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
func Unix(sec int64, nsec int64) Time { //md5:8f372b5d16d3f416

}

// UnixMilli returns the local Time corresponding to the given Unix time,
// msec milliseconds since January 1, 1970 UTC.
func UnixMilli(msec int64) Time { //md5:94c00e0069987ce8

}

// UnixMicro returns the local Time corresponding to the given Unix time,
// usec microseconds since January 1, 1970 UTC.
func UnixMicro(usec int64) Time { //md5:a519cb1431da9d9d

}

// IsDST reports whether the time in the configured location is in Daylight Savings Time.
func (t Time) IsDST() bool { //md5:6a98901b623dee23

}

// Date returns the Time corresponding to
//
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time { //md5:fc73fa423236be91

}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t Time) Truncate(d Duration) Time { //md5:6c9853dfe8f0ce4e

}

// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t Time) Round(d Duration) Time { //md5:514f2ad77ae8743a

}