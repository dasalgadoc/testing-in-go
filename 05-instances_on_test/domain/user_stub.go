package domain

import "math/rand"

var names = []string{
	"John Doe",
	"Jane Doe",
	"Foo Bar",
	"Bar Foo",
	"Lorem Ipsum",
	"Ipsum Lorem",
}

func UserWithName(name string) *User {
	id, err := NewUserId()
	if err != nil {
		panic(err)
	}
	userName, err := NewUserName(name)
	if err != nil {
		panic(err)
	}
	userAccessLevel, err := NewUserAccessLevelFromInt(randomAccessLevel())
	if err != nil {
		panic(err)
	}

	return &User{
		id:          id,
		name:        userName,
		accessLevel: userAccessLevel,
	}
}

func UserWithAccessLevel(accessLevel int) *User {
	id, err := NewUserId()
	if err != nil {
		panic(err)
	}
	userName, err := NewUserName(randomName())
	if err != nil {
		panic(err)
	}
	userAccessLevel, err := NewUserAccessLevelFromInt(accessLevel)
	if err != nil {
		panic(err)
	}

	return &User{
		id:          id,
		name:        userName,
		accessLevel: userAccessLevel,
	}
}

func UserWithID(id string) *User {
	userId, err := NewUserIdFromString(id)
	if err != nil {
		panic(err)
	}
	userName, err := NewUserName(randomName())
	if err != nil {
		panic(err)
	}
	userAccessLevel, err := NewUserAccessLevelFromInt(randomAccessLevel())
	if err != nil {
		panic(err)
	}

	return &User{
		id:          userId,
		name:        userName,
		accessLevel: userAccessLevel,
	}
}

func UserWithParameters(data map[string]any) *User {
	dataId, ok := data["id"]
	if !ok {
		id, err := NewUserId()
		if err != nil {
			panic(err)
		}
		dataId = id.Value()
	}
	userId, err := NewUserIdFromString(dataId.(string))
	if err != nil {
		panic(err)
	}

	dataName, ok := data["name"]
	if !ok {
		dataName = randomName()
	}
	userName, err := NewUserName(dataName.(string))
	if err != nil {
		panic(err)
	}

	dataAccessLevel, ok := data["accessLevel"]
	if !ok {
		dataAccessLevel = randomAccessLevel()
	}
	userAccessLevel, err := NewUserAccessLevelFromInt(dataAccessLevel.(int))
	if err != nil {
		panic(err)
	}

	return &User{
		id:          userId,
		name:        userName,
		accessLevel: userAccessLevel,
	}
}

func randomAccessLevel() int {
	keys := make([]string, 0, len(accessLevelMap))
	for k := range accessLevelMap {
		keys = append(keys, k)
	}

	randomIndex := rand.Intn(len(keys))
	randomKey := keys[randomIndex]
	randomValue := accessLevelMap[randomKey]

	return randomValue
}

func randomName() string {
	randomIndex := rand.Intn(len(names))
	return names[randomIndex]
}
