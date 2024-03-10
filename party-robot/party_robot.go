package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %[1]s!\n"+
		"You have been assigned to table %03[2]d. Your table is %[4]s, "+
		"exactly %.1[5]f meters from here.\n"+
		"You will be sitting next to %[3]s.",
		name, table, neighbor, direction, distance)
}
