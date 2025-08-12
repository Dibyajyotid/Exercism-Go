//Package weather provides tools to 
//forcast weather conditions for various cities.
package weather 

//CurrentCondition stores the current weather condition in string format.
var CurrentCondition string

//CurrentLocation stores the name of the city to be forcast.
var CurrentLocation string

//Forecast function takes city and condition as parameters
//returns a string output with location and weather condition of that particular location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
