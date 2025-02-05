package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	for _, f := range []string{
		"1/02/2006 15:04:05", "1/2/2006 15:04:05",
		"January 2, 2006 15:04:05", "Monday, January 2, 2006 15:04:05",
	} {
		t, err := time.Parse(f, date)
		if err == nil {
			return t
		}
	}
	panic("unsupported format")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	h, _, _ := Schedule(date).Clock()
	return h >= 12 && h <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return `You have an appointment on ` + Schedule(date).Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
