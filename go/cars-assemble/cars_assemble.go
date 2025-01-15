package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var (
		cost       uint
		groups     int
		cogs           = 10000
		batch_cost int = 95000
	)
	if carsCount == 1 {
		cost = uint(cogs)
	} else {
		remainder := carsCount % 10
		if remainder >= 0 {
			groups = carsCount / 10
		}
		cost = uint((groups * batch_cost) + ((remainder) * (cogs)))
	}
	return cost
}
