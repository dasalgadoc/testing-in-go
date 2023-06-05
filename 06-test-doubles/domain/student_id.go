package domain

import "github.com/google/uuid"

type StudentId struct {
	value uuid.UUID
}

func NewStudentId() (*StudentId, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &StudentId{value: id}, nil
}

func NewStudentIdFromString(id string) (*StudentId, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &StudentId{value: uid}, nil
}

func (s *StudentId) Value() string {
	return s.value.String()
}
