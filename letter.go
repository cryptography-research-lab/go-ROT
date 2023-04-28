package rot

func ROTForLetter[T byte | rune](value T, rotN int) T {
	if value >= 'a' && value <= 'z' {
		value = T((int(value-'a') + rotN)%26 + int('a'))
	} else if value >= 'A' && value <= 'Z' {
		value = T((int(value-'A') + rotN)%26 + int('A'))
	}
	return value
}

func ROT5ForLetter[T byte | rune](value T) T {
	return ROTForLetter(value, 5)
}

func ROT13ForLetter[T byte | rune](value T) T {
	return ROTForLetter(value, 13)
}

func ROT18ForLetter[T byte | rune](value T) T {
	return ROTForLetter(value, 18)
}

func ROT47ForLetter[T byte | rune](value T) T {
	return ROTForLetter(value, 47)
}
