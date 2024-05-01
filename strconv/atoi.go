package strconv

// A NumError records a failed conversion.
type NumError struct {

} //md5:0cd5b5c6c9f477b9

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize //md5:19007de8c1f03835



// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range") //md5:879a72936976b39e



// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax") //md5:5371aebaa9b3d94e




func (e *NumError) Error() string { //md5:d92759861dd6caa4

}


func (e *NumError) Unwrap() error { //md5:35635a0deb05e0a6

}

// ParseUint is like ParseInt but for unsigned numbers.
//
// A sign prefix is not permitted.
func ParseUint(s string, base int, bitSize int) (uint64, error) { //md5:f4c2829256142922

}

// ParseInt interprets a string s in the given base (0, 2 to 36) and
// bit size (0 to 64) and returns the corresponding value i.
//
// The string may begin with a leading sign: "+" or "-".
//
// If the base argument is 0, the true base is implied by the string's
// prefix following the sign (if present): 2 for "0b", 8 for "0" or "0o",
// 16 for "0x", and 10 otherwise. Also, for argument base 0 only,
// underscore characters are permitted as defined by the Go syntax for
// [integer literals].
//
// The bitSize argument specifies the integer type
// that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
// correspond to int, int8, int16, int32, and int64.
// If bitSize is below 0 or above 64, an error is returned.
//
// The errors that ParseInt returns have concrete type *NumError
// and include err.Num = s. If s is empty or contains invalid
// digits, err.Err = ErrSyntax and the returned value is 0;
// if the value corresponding to s cannot be represented by a
// signed integer of the given size, err.Err = ErrRange and the
// returned value is the maximum magnitude integer of the
// appropriate bitSize and sign.
//
// [integer literals]: https://go.dev/ref/spec#Integer_literals
func ParseInt(s string, base int, bitSize int) (i int64, err error) { //md5:8bc39378cce23cd2

}

// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
func Atoi(s string) (int, error) { //md5:576d8377bb1d7532

}