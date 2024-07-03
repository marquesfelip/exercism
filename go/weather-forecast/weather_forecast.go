// Package weather package provides a function to get the current weather condition of a city.
package weather

// CurrentCondition is package level variable to store the current weather condition.
var CurrentCondition string
// CurrentLocation is package level variable to store the current location.
var CurrentLocation string

// Forecast receives a city and a weather condition and returns a string with the city and the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
