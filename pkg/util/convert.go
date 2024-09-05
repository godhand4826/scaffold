package util

import (
	"strconv"
)

// Atob returns `true` only if the value is truthy.
func Atob(value string) bool {
	return Deref(AtobPtr(value))
}

func AtobPtr(value string) *bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}
	return &b
}
