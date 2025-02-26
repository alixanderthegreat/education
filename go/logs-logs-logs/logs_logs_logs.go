package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, that := range log {
		switch that {
		case 10071:
			return "recommendation"
		case 128269:
			return "search"
		case 9728:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
