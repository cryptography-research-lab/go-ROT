package rot

import (
	"github.com/golang-infrastructure/go-gtypes"
)

func ROT[T gtypes.Integer](value, rotN T) T {
	return value + rotN
}

func ROT5[T gtypes.Integer](value T) T {
	return ROT(value, 5)
}

func ROT13[T gtypes.Integer](value T) T {
	return ROT(value, 13)
}

func ROT18[T gtypes.Integer](value T) T {
	return ROT(value, 18)
}

func ROT47[T gtypes.Integer](value T) T {
	return ROT(value, 47)
}
