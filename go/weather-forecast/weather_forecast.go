// Package weather provides tools to weather forecasting.
package weather

// CurrentCondition represents a current condition weather forecast.
var CurrentCondition string

// CurrentLocation represents a current location weather forecast.
var CurrentLocation string

// Forecast returns an string weather condition/forecasting information.
func Forecast(city, condition string) string {
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
