package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int
	for _, bird := range birdsPerDay {
		total += bird
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// :)
	return TotalBirdCount(birdsPerDay[(7*week)-7 : (7 * week)])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 { // we are assuming ever other day
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
