package domain

type UserBuilder struct {
	User User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) WithId(id string) *UserBuilder {
	uid, err := NewUserIdFromString(id)
	if err != nil {
		panic(err)
	}

	b.User.id = uid
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	un, err := NewUserName(name)
	if err != nil {
		panic(err)
	}

	b.User.name = un
	return b
}

func (b *UserBuilder) WithAccessLevel(accessLevel int) *UserBuilder {
	ual, err := NewUserAccessLevelFromInt(accessLevel)
	if err != nil {
		panic(err)
	}

	b.User.accessLevel = ual
	return b
}

func (b *UserBuilder) Build() User {
	return b.User
}
