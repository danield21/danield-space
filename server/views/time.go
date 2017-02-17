package views

import (
	"time"
)

//HumanTime takes the time and display in a human readable form
func HumanTime(date time.Time) string {
	return date.Format(time.UnixDate)
}

//MachineTime takes the time and display in a Machine readable form
func MachineTime(date time.Time) string {
	return date.UTC().String()
}
