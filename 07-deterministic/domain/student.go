package domain

import "time"

var CurrentHour = func() int {
	return time.Now().Hour()
}

type Student struct {
	ID    *StudentId
	Shift *StudentShift
}

func NewStudent(shift StudentShift) (*Student, error) {
	id, err := NewStudentId()
	if err != nil {
		return nil, err
	}
	return &Student{ID: id, Shift: &shift}, nil
}

func NewStudentWithId(id StudentId, shift StudentShift) (*Student, error) {
	return &Student{ID: &id, Shift: &shift}, nil
}

func (s *Student) CanTakeClassNow() bool {
	hour := CurrentHour()
	return s.Shift.Value() == getShiftFromHour(hour)
}
