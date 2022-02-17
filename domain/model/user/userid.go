package entity

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type userId string

func NewUserId(v string) (*userId, error) {
	userId := userId(v)
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
