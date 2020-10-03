package group

import (
	"testing"
)

func TestNewGroup(t *testing.T) {
	group := NewGroup("test group")
	if group.title != "test group" {
		t.Errorf("group title is wrong ")
	}
}

func TestAddStudent(t *testing.T) {
	student_id := 10
	group := NewGroup("test group")
	group.AddStudent(student_id)
	student := group.students[0]
	if student != student_id {
		t.Errorf("student not added ")
	}
}

func TestAddDupStudent(t *testing.T) {
	student_id := 10
	group := NewGroup("test group")
	group.AddStudent(student_id)

	err := group.AddStudent(student_id)
	if err == nil {
		t.Errorf("student added twice")
	}
}

func TestIsAttending(t *testing.T) {
	student_id := 10
	group := NewGroup("test group")
	group.AddStudent(student_id)
	if !group.is_attending(10) {
		t.Errorf("student is not attending, but should be attending")
	}
}
