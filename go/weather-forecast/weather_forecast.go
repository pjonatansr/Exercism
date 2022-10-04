// Package weather can be used to get the weather forecast for a given city.
package weather

// CurrentCondition variable current condition is used to store the current weather condition.
var CurrentCondition string
// CurrentLocation variable current location is used to store the current location.
var CurrentLocation string

// Forecast function is used to get the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
