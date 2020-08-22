// Determine if a year if a leap year

package leap

// IsLeapYear returns true for a leap year, false otherwise
func IsLeapYear(year int) bool {
	return (year%4 == 0) && ((year%100 != 0) || (year%400 == 0))
}
