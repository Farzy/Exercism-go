package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	tDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	return tDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	tDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := tDate.Hour()
	return 12 <= hour && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	tDate := Schedule(date)
	return tDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	//anniversary, _ := time.Parse("2006-01-02", fmt.Sprintf("%d-09-15", year))
	anniversary := time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
