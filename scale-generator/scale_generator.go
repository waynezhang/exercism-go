package scale

import (
	"strings"
)

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var steps = map[rune]int{'m': 1, 'M': 2, 'A': 3}

// Scale doc here
func Scale(tonic string, interval string) []string {
	var array = &sharps

	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		array = &flats
	}

	start := 0
	for i, ch := range *array {
		if strings.ToLower(ch) == strings.ToLower(tonic) {
			start = i
			break
		}
	}

	if interval == "" {
		return append((*array)[start:], (*array)[:start]...)
	}

	ret := []string{}
	for _, n := range interval {
		step := steps[n]
		ret = append(ret, (*array)[start])
		start = (start + step) % len(*array)
	}

	return ret
}
