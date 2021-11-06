// Package weather helps forecast the weather based on the location and condition.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the city name.
var CurrentLocation string

// Forecast computes the weather forecast based on the city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
