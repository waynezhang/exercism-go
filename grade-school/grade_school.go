package school

import "sort"

// Grade doc here
type Grade struct {
	grade int
	names []string
}

// School doc here
type School struct {
	grades map[int]Grade
}

// New doc here
func New() *School {
	return &School{grades: make(map[int]Grade)}
}

// Add doc here
func (s *School) Add(name string, grade int) {
	var g *Grade
	if val, ok := s.grades[grade]; ok {
		g = &val
	} else {
		g = &Grade{grade: grade}
	}
	g.names = append(g.names, name)
	s.grades[grade] = *g
}

// Grade doc here
func (s *School) Grade(grade int) []string {
	return s.grades[grade].names
}

// Enrollment doc here
func (s *School) Enrollment() []Grade {
	grades := make([]Grade, 0)
	for _, v := range s.grades {
		sort.Strings(v.names)
		grades = append(grades, v)
	}
	sort.Slice(grades, func(i, j int) bool { return grades[i].grade < grades[j].grade })
	return grades
}
