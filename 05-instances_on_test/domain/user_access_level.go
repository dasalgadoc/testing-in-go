package domain

import "errors"

const MIN_EDITTER_USER_ACCESS_LEVEL = 2

var accessLevelMap = map[string]int{
	"USER":   1,
	"AUTHOR": 2,
	"ADMIN":  3,
}

type UserAccessLevel struct {
	accessLevel int
}

func NewUserAccessLevelFromInt(accessLevel int) (*UserAccessLevel, error) {
	if !isValidUserAccessLevel(accessLevel) {
		return nil, errors.New("invalid access level")
	}
	return &UserAccessLevel{accessLevel: accessLevel}, nil
}

func NewUserAccessLevelFromString(accessLevel string) (*UserAccessLevel, error) {
	level, ok := accessLevelMap[accessLevel]
	if !ok {
		return nil, errors.New("invalid access level")
	}
	return &UserAccessLevel{accessLevel: level}, nil
}

func isValidUserAccessLevel(accessLevel int) bool {
	for _, v := range accessLevelMap {
		if v == accessLevel {
			return true
		}
	}
	return false
}

func (s *UserAccessLevel) Value() int {
	return s.accessLevel
}
