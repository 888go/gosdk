package time


const (
	seekStart   = 0
	seekCurrent = 1
	seekEnd     = 2
)// md5:489a8240bba6ec1a



// LoadLocationFromTZData returns a Location with the given name
// initialized from the IANA Time Zone database-formatted data.
// The data should be in the format of a standard IANA time zone file
// (for example, the content of /etc/localtime on Unix systems).
func LoadLocationFromTZData(name string, data []byte) (*Location, error) { //md5:31b2d35bb24a50e9

}