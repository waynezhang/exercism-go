package kindergarten

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// Garden doc here
type Garden struct {
	names []string
	cups  []string
}

// NewGarden doc here
func NewGarden(diagram string, children []string) (*Garden, error) {
	if regexp.MustCompile("[^VCRG\n]+").MatchString(diagram) {
		return nil, errors.New("")
	}
	splited := strings.Split(diagram, "\n")[1:]
	if len(splited) != 2 || len(splited[0]) != len(children)*2 {
		return nil, errors.New("")
	}

	garden := new(Garden)
	garden.names = append([]string{}, children...)
	sort.Sort(sort.StringSlice(garden.names))
	// check duplicated name
	for i := 0; i < len(garden.names)-1; i++ {
		if garden.names[i] == garden.names[i+1] {
			return nil, errors.New("")
		}
	}

	garden.cups = splited
	return garden, nil
}

// Plants doc here
func (g *Garden) Plants(child string) ([]string, bool) {
	plants := map[byte]string{
		'V': "violets",
		'C': "clover",
		'R': "radishes",
		'G': "grass",
	}

	idx := -1
	for i, n := range g.names {
		if n == child {
			idx = i
			break
		}
	}

	if idx == -1 {
		return nil, false
	}

	ret := make([]string, 4)
	ret[0] = plants[g.cups[0][idx*2]]
	ret[1] = plants[g.cups[0][idx*2+1]]
	ret[2] = plants[g.cups[1][idx*2]]
	ret[3] = plants[g.cups[1][idx*2+1]]
	return ret, true
}
