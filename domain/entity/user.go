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

func (user *User) ChangeUserName(name userName) error {
	if err := name.validate(name); err != nil {
		return err
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
