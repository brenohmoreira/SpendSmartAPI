package util

import (
	"os"
	"strconv"
)

func GetOrDefaultPrimitive[T comparable](value T, fallback T) T {
	var zeroValue T
	if value == zeroValue {
		return fallback
	}
	return value
}

func GetEnvInt(key string, fallback int) int {
	v := os.Getenv(key)

	if v == "" {
		return fallback
	}

	i, err := strconv.Atoi(v)

	if err != nil {
		return fallback
	}

	return i
}
