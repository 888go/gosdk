package strconv

// FormatUint returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatUint(i uint64, base int) string { //md5:25e76e61cf965490

}

// FormatInt returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatInt(i int64, base int) string { //md5:c8ffdd9b6b31204e

}

// Itoa is equivalent to FormatInt(int64(i), 10).
func Itoa(i int) string { //md5:855a975d9bf4f4e5

}

// AppendInt appends the string form of the integer i,
// as generated by FormatInt, to dst and returns the extended buffer.
func AppendInt(dst []byte, i int64, base int) []byte { //md5:f8f82dfdc15b8e82

}

// AppendUint appends the string form of the unsigned integer i,
// as generated by FormatUint, to dst and returns the extended buffer.
func AppendUint(dst []byte, i uint64, base int) []byte { //md5:7faa1c64ec12eba6

}