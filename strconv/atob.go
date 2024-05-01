package strconv

// ParseBool returns the boolean value represented by the string.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func ParseBool(str string) (bool, error) { //md5:6f1b6c8e55a90306

}

// FormatBool returns "true" or "false" according to the value of b.
func FormatBool(b bool) string { //md5:b07d776a75e22602

}

// AppendBool appends "true" or "false", according to the value of b,
// to dst and returns the extended buffer.
func AppendBool(dst []byte, b bool) []byte { //md5:cfa4028967205e76

}