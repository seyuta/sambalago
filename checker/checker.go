package checker

import "strings"

// Function to check if the given string is a valid Roman numeral
func IsValidRoman(s string) bool {
	s = strings.ToUpper(s)
	validChars := map[byte]bool{
		'M': true,
		'D': true,
		'C': true,
		'L': true,
		'X': true,
		'V': true,
		'I': true,
	}

	romanMap := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	n := len(s)
	for i := 0; i < n; i++ {
		if !validChars[s[i]] {
			return false
		}

		if i < n-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			if !(s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X')) &&
				!(s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C')) &&
				!(s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M')) {
				return false
			}
		}
	}

	return true
}
