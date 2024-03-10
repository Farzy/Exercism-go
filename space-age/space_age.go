package space

type Planet string

// Number of seconds in an Earth year
const SecondsInEarthYear = 31557600

// Ratio by which to divide the age on Earth to find the age on "planet"
var yearRatio = map[Planet]float64{
	Planet("Mercury"): 0.2408467,
	Planet("Venus"):   0.61519726,
	Planet("Earth"):   1.0,
	Planet("Mars"):    1.8808158,
	Planet("Jupiter"): 11.862615,
	Planet("Saturn"):  29.447498,
	Planet("Uranus"):  84.016846,
	Planet("Neptune"): 164.79132,
}

func Age(seconds float64, planet Planet) (age float64) {
	yr, ok := yearRatio[planet]
	if ok {
		age = seconds / SecondsInEarthYear / yr
	} else {
		age = -1
	}
	return
}
