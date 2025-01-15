package techpalace

import "strings"

const prefix = `Welcome to the Tech Palace, `

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return prefix + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	panic("Please implement the CleanupMessage() function")
}
