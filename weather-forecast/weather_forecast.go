// Package weather transforms information about a forecast and location
// into a string description.
package weather

// CurrentCondition describes the weather conditions in a location.
var CurrentCondition string

// CurrentLocation is the location being described by the conditions.
var CurrentLocation string

// Forecast takes a city and condition string and returns another string
// describing the weather in that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
