package booking

import (
	"strconv"
	"time"
)

func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"

	t, _ := time.Parse(layout, date)

	return t
}

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t.Hour() >= 12 && t.Hour() < 18
}

func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return "You have an appointment on " + t.Weekday().String() + ", " + t.Month().String() + " " + strconv.Itoa(t.Day()) + ", " + strconv.Itoa(t.Year()) + ", at " + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + "."
}

func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
