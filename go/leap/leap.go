// Package leap provides a function to determine if a given year is a leap year.
package leap

// IsLeapYear determines whether a given year is a leap year.
// A leap year is divisible by 4, but not divisible by 100 unless it is also divisible by 400.
//
// Parameters:
//   year (int): The year to be checked.
//
// Returns:
//   bool: True if the year is a leap year, false otherwise.
func IsLeapYear(year int) bool {
	if (year % 100 == 0 && year % 400 != 0) {
		return false
	}

	if year % 4 == 0 {
		return true
	}

	return false
}
