package domain

import "errors"

type StudentName struct {
	value string
}

func NewStudentName(value string) (*StudentName, error) {
	if !isValidStudentName(value) {
		return nil, errors.New("invalid student name")
	}
	return &StudentName{value: value}, nil
}

func (s *StudentName) Value() string {
	return s.value
}

func isValidStudentName(value string) bool {
	return len(value) > 0
}
