package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId string

func NewUserId(v string) (*UserId, error) {
	userId := UserId(v)
	if err := userId.validate(userId); err != nil {
		return nil, err
	}
	return &userId, nil
}

func (userId *UserId) validate(id UserId) (err error) {
	defer iterrors.Wrap(&err, "userId.validate(%q)", id)
	if id == "" {
		return fmt.Errorf("userId is required")
	}
	return nil
}
