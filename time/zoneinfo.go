package time

// A Location maps time instants to the zone in use at that time.
// Typically, the Location represents the collection of time offsets
// in use in a geographical area. For many Locations the time offset varies
// depending on whether daylight savings time is in use at the time instant.
//
// Location is used to provide a time zone in a printed Time value and for
// calculations involving intervals that may cross daylight savings time
// boundaries.
type Location struct {

} //md5:b2741259c471a509


const (
	alpha = -1 << 63  // math.MinInt64
	omega = 1<<63 - 1 // math.MaxInt64
)// md5:7214a6052ed65a5f




const (
	ruleJulian ruleKind = iota
	ruleDOY
	ruleMonthWeekDay
)// md5:cd28ad5922a22ce1



// UTC represents Universal Coordinated Time (UTC).
var UTC *Location = &utcLoc //md5:cb0833e2130b743d



// Local represents the system's local time zone.
// On Unix systems, Local consults the TZ environment
// variable to find the time zone to use. No TZ means
// use the system default /etc/localtime.
// TZ="" means use UTC.
// TZ="foo" means use file foo in the system timezone directory.
var Local *Location = &localLoc //md5:57aa689da8ee76b4



// String returns a descriptive name for the time zone information,
// corresponding to the name argument to LoadLocation or FixedZone.
func (l *Location) String() string { //md5:c44612032555be76

}

// FixedZone returns a Location that always uses
// the given zone name and offset (seconds east of UTC).
func FixedZone(name string, offset int) *Location { //md5:839d95c9de0383bc

}

// LoadLocation returns the Location with the given name.
//
// If the name is "" or "UTC", LoadLocation returns UTC.
// If the name is "Local", LoadLocation returns Local.
//
// Otherwise, the name is taken to be a location name corresponding to a file
// in the IANA Time Zone database, such as "America/New_York".
//
// LoadLocation looks for the IANA Time Zone database in the following
// locations in order:
//
//   - the directory or uncompressed zip file named by the ZONEINFO environment variable
//   - on a Unix system, the system standard installation location
//   - $GOROOT/lib/time/zoneinfo.zip
//   - the time/tzdata package, if it was imported
func LoadLocation(name string) (*Location, error) { //md5:506ad5b64122238f

}