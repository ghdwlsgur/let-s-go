package lib

// 외부에서 접근 가능
func IsDigit(c int32) bool {
	return '0' <= c && c <= '9'
}

// 외부에서 접근 불가
func isSpace(c int32) bool {
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
