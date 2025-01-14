// Package weather ...
package weather

var (
	// CurrentCondition ...
	CurrentCondition string

	// CurrentLocation ...
	CurrentLocation string
)

// Forecast ...
func Forecast(city, condition string) string {

	CurrentLocation, CurrentCondition =
		city, condition

	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
