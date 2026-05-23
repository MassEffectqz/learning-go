package booking

import "time"

func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	h := t.Hour()
	return h >= 12 && h < 18
}
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04.")
}
func AnniversaryDate() time.Time {
	year := time.Now().UTC().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
