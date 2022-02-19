package user

import (
	"fmt"
	"reflect"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type User struct {
	userId UserId
	name   string
}

func NewUser(userId UserId, name string) (*User, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	user := &User{userId: userId, name: name}

	return user, nil
}

func (user *User) ChangeUserName(name string) (err error) {
	defer iterrors.Wrap(&err, "user.ChangeUserName(%q)", name)
	if err := validateName(name); err != nil {
		return err
	}
	user.name = name
	return nil
}

func validateName(name string) (err error) {
	defer iterrors.Wrap(&err, "user.go validateUserName(%q)", name)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if len(name) < 3 {
		return fmt.Errorf("name %v is less than three characters long", name)
	}
	return nil
}

func (user *User) Equals(other *User) bool {
	return reflect.DeepEqual(user.userId, other.userId)
}
