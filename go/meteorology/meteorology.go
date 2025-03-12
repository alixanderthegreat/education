package meteorology

import (
	"fmt"
	"strconv"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
	if t != 0 {
		return "°F"
	}
	return "°C"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (s Temperature) String() string {
	return strconv.Itoa(s.degree) + " " + s.unit.String()

}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	if s != 0 {
		return "mph"
	}
	return "km/h"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return strconv.Itoa(s.magnitude) + " " + s.unit.String()

}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (s MeteorologyData) String() string {
	return fmt.Sprintf(
		"%s: %s, Wind %s at %s, %d%% Humidity",
		s.location, s.temperature.String(), s.windDirection, s.windSpeed.String(), s.humidity,
	)
}
