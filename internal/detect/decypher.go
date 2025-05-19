package detect

import (
	"fmt"
	"unicode"
)

// BruteCaesar tries all Caesar shifts from -25 to +25 (excluding 0)
func BruteCaesar(input string) {
	fmt.Println("[+] Caesar Brute-force (ROT-N from -25 to +25):")

	for shift := -25; shift <= 25; shift++ {
		if shift == 0 {
			continue
		}

		decoded := caesarShift(input, shift)
		fmt.Printf("[%+03d] %s\n", shift, decoded)
	}
}

// caesarShift shifts letters by N positions (wraps A-Z and a-z)
func caesarShift(s string, shift int) string {
	shift = (shift%26 + 26) % 26 // normalize shift (handles negative properly)
	runes := []rune(s)

	for i, r := range runes {
		switch {
		case unicode.IsUpper(r):
			runes[i] = 'A' + rune((int(r-'A')+shift)%26)
		case unicode.IsLower(r):
			runes[i] = 'a' + rune((int(r-'a')+shift)%26)
		default:
			// Leave other characters as is (digits, punctuation, etc.)
			runes[i] = r
		}
	}

	return string(runes)
}

// ROT13
func ROT13(input string) string {
	return caesarShift(input, 13)
}

// ROT47
func ROT47(input string) string {
	runes := []rune(input)
	for i, r := range runes {
		if r >= 33 && r <= 126 { // ASCII range ! to ~
			runes[i] = 33 + ((r - 33 + 47) % 94)
		}
	}
	return string(runes)
}

// Atbash (mirror alphabet)
func Atbash(input string) string {
	runes := []rune(input)
	for i, r := range runes {
		switch {
		case unicode.IsUpper(r):
			runes[i] = 'Z' - (r - 'A')
		case unicode.IsLower(r):
			runes[i] = 'z' - (r - 'a')
		}
	}
	return string(runes)
}
