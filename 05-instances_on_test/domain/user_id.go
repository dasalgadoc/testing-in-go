package domain

import "github.com/google/uuid"

type UserId struct {
	value uuid.UUID
}

func NewUserId() (*UserId, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &UserId{value: id}, nil
}

func NewUserIdFromString(id string) (*UserId, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &UserId{value: uid}, nil
}

func (s *UserId) Value() string {
	return s.value.String()
}
