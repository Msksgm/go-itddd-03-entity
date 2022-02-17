package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type User struct {
	id   userId
	name string
}

func NewUser(userId userId, name string) (*User, error) {
	user := &User{id: userId, name: name}

	if err := userId.validate(userId); err != nil {
		return nil, err
	}

	if err := user.ChangeUserName(name); err != nil {
		return nil, err
	}
	return user, nil
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