package entity

import "fmt"

type User struct {
	id   userId
	name userName
}

func (user *User) NewUser(userId userId, userName userName) (*User, error) {
	return &User{id: userId, name: userName}, nil
}

func (user *User) ChangeUserName(name userName) error {
	if name == "" {
		return fmt.Errorf("氏名は必須です")
	}
	if len(name) < 3 {
		return fmt.Errorf("氏名は3文字以上にしてください")
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
