package space

type Planet string

func Age(age float64, planet Planet) float64 {
	age /= 31557600

	switch planet {
	case "Mercury":
		age /= 0.2408467
	case "Venus":
		age /= 0.61519726
	case "Mars":
		age /= 1.8808158
	case "Jupiter":
		age /= 11.862615
	case "Saturn":
		age /= 29.447498
	case "Uranus":
		age /= 84.016846
	case "Neptune":
		age /= 164.79132
	}
	return age
}
