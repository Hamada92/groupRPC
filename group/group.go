package group

import (
	"fmt"
)

type Group struct {
	title    string
	students []int
}

func NewGroup(title string) *Group {
	g := Group{title: title}
	return &g
}

func (g *Group) is_attending(student_id int) bool {
	for i := range g.students {
		if student_id == g.students[i] {
			return true
		}
	}
	return false
}

func (g *Group) AddStudent(student_id int) error {
	if g.is_attending(student_id) {
		return ErrDupStudent
	}

	g.students = append(g.students, student_id)
	return nil
}

var (
	ErrDupStudent = fmt.Errorf("student already attending")
)
