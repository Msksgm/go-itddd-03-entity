package entity

type User struct {
	id   userId
	name userName
}

func (user *User) NewUser(userId userId, userName userName) (*User, error) {
	if err := userId.validate(userId); err != nil {
		return nil, err
	}
	if err := userName.validate(userName); err != nil {
		return nil, err
	}
	return &User{id: userId, name: userName}, nil
}

func (user *User) ChangeUserName(name string) (err error) {
	defer iterrors.Wrap(&err, "user.ChangeUserName(%q)", name)
	if name == "" {
		return fmt.Errorf("userName is required")
	}
	if len(name) < 3 {
		return fmt.Errorf("userName %v is less than three characters long", name)
	}
	user.name = name
	return nil
}

func (user *User) Equals(other User) bool {
	if user == &other {
		return true
	}
	if user.id == other.id {
		return true
	}
	return false
}
