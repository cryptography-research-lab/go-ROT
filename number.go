package rot

func ROTForNumber[T byte | rune](value T, rotN int) T {
	if value >= '0' && value <= '9' {
		value = T((int(value-'0')+rotN)%10 + int('0'))
	}
	return value
}
