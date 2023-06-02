package nfcnorm

import (
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// Length returns the rune count of `s`.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// Normalizable checks if `s` can be normalized to NFC, that is, is not a NFC string.
func Normalizable(s string) bool {
	return !norm.NFC.IsNormalString(s)
}

// Normalize returns the normalized NFC string value of `s`.
func Normalize(s string) string {
	return norm.NFC.String(s)
}
