package domain

type User struct {
	id          *UserId
	name        *UserName
	accessLevel *UserAccessLevel
}

func NewUser(name UserName, accessLevel UserAccessLevel) (*User, error) {
	uid, err := NewUserId()
	if err != nil {
		return nil, err
	}
	return &User{id: uid, name: &name, accessLevel: &accessLevel}, nil
}

func (u *User) CanEditBlog() bool {
	return u.accessLevel.Value() >= MIN_EDITTER_USER_ACCESS_LEVEL
}

func (u *User) Name() string {
	return u.name.Value()
}

func (u *User) ID() string {
	return u.id.Value()
}

func (u *User) AccessLevel() int {
	return u.accessLevel.Value()
}
