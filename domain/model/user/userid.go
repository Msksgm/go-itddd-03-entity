package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type userId string

func NewUserId(id string) (*userId, error) {
	userId := userId(id)
	if err := userId.validate(userId); err != nil {
		return nil, err
	}
	return &userId, nil
}

func (userId *userId) validate(id userId) (err error) {
	defer iterrors.Wrap(&err, "userId.validate(%q)", id)
	if id == "" {
		return fmt.Errorf("userId is required")
	}
	return nil
}
