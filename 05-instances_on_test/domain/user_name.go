package domain

import "errors"

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if !isValidUserName(value) {
		return nil, errors.New("invalid user name")
	}
	return &UserName{value: value}, nil
}

func (s *UserName) Value() string {
	return s.value
}

func isValidUserName(value string) bool {
	return len(value) > 0
}
