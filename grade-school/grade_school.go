package school

import (
	"sort"
)

// Define the Grade and School types here.
type Grade struct {
	level    int
	students []string
}
type School struct {
	grades map[int]*Grade
}

func New() *School {
	return &School{
		grades: make(map[int]*Grade),
	}
}

func (s *School) Add(student string, level int) {
	grade, ok := s.grades[level]
	if ok {
		grade.students = append(grade.students, student)
		return
	}
	s.grades[level] = &Grade{
		level, []string{student},
	}
}

func (s *School) Grade(level int) []string {
	grade, ok := s.grades[level]
	if ok {
		return grade.students
	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	levels := make([]int, 0)
	for k := range s.grades {
		levels = append(levels, k)
	}
	sort.Ints(levels)

	grades := make([]Grade, 0, len(s.grades))
	for _, level := range levels {
		grade := s.grades[level]
		sort.Strings(grade.students)
		grades = append(grades, *grade)
	}

	return grades
}
