package entity

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type userName string

func NewUserName(v string) (*userName, error) {
	userName := userName(v)
	if err := userName.validate(userName); err != nil {
		return nil, err
	}
	return &userName, nil
}

func (userName *userName) validate(v userName) (err error) {
	defer iterrors.Wrap(&err, "userName.validate(%q)", v)
	if v == "" {
		return fmt.Errorf("userName is required")
	}
	if len(v) < 3 {
		return fmt.Errorf("userName %v is less than three characters long", *userName)
	}
	return nil
}
