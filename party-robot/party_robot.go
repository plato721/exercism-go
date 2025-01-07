package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	msg := "Happy birthday %s! You are now %d years old!"
	return fmt.Sprintf(msg, name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	msg1 := Welcome(name)
	msg2 := tableMessage(table)
	msg3 := directionsMessage(direction, distance)
	msg4 := neighborMessage(neighbor)
	return fmt.Sprintf("%s\n%s %s\n%s", msg1, msg2, msg3, msg4)
}

func directionsMessage(direction string, distance float64) string {
	return fmt.Sprintf("Your table is %s, exactly %.1f meters from here.", direction, distance)
}

func neighborMessage(neighbor string) string {
	return fmt.Sprintf("You will be sitting next to %s.", neighbor)
}

func tableMessage(table int) string {
	return fmt.Sprintf("You have been assigned to table %03d.", table)
}
