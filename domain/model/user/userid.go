package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId string

func NewUserId(id string) (*UserId, error) {
	UserId := UserId(id)
	if err := UserId.validate(UserId); err != nil {
		return nil, err
	}
	return &UserId, nil
}

func (UserId *UserId) validate(userId UserId) (err error) {
	defer iterrors.Wrap(&err, "UserId.validate(%q)", userId)
	if userId == "" {
		return fmt.Errorf("userId is required")
	}
	return nil
}
