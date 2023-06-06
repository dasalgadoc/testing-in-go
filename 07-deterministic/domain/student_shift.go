package domain

import "errors"

type StudentShift struct {
	value string
}

func NewStudentShift(value string) (*StudentShift, error) {
	if !isShiftRegistered(value) {
		return nil, errors.New("invalid time shift")
	}
	return &StudentShift{value: value}, nil
}

func NewStudentShiftFromHour(hour int) (*StudentShift, error) {
	value := getShiftFromHour(hour)
	if value == "" {
		return nil, errors.New("invalid time shift")
	}
	return &StudentShift{value: value}, nil
}

func (s *StudentShift) Value() string {
	return s.value
}
