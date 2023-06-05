package domain

import "errors"

type StudentAge struct {
	value int
}

func NewStudentAge(value int) (*StudentAge, error) {
	if value < 0 {
		return nil, errors.New("age can't be negative")
	}
	return &StudentAge{value: value}, nil
}

func (s *StudentAge) Value() int {
	return s.value
}
