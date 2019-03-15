package space

// Age function to calculate space age
func Age(ageSec float64, planet string) float64 {
	planetAge := make(map[string]float64)

	planetAge["Earth"] = float64(31557600)
	planetAge["Mercury"] = planetAge["Earth"] * float64(0.2408467)
	planetAge["Venus"] = planetAge["Earth"] * float64(0.61519726)
	planetAge["Mars"] = planetAge["Earth"] * float64(1.8808158)
	planetAge["Jupiter"] = planetAge["Earth"] * float64(11.862615)
	planetAge["Saturn"] = planetAge["Earth"] * float64(29.447498)
	planetAge["Uranus"] = planetAge["Earth"] * float64(84.016846)
	planetAge["Neptune"] = planetAge["Earth"] * float64(164.79132)

	return ageSec / planetAge[planet]
}
