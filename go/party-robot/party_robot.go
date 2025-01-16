package partyrobot

import "fmt"

var spf = fmt.Sprintf

// Welcome greets a person by name.
func Welcome(name string) string {
	return spf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return spf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var msg string
	msg += "Welcome to my party, %s!\n"
	msg += "You have been assigned to table %03d. "
	msg += "Your table is %s, "
	msg += "exactly %.1f meters from here.\n"
	msg += "You will be sitting next to %s."
	return spf(msg, name, table, direction, distance, neighbor)
}
