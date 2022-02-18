package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId struct {
	userId string
}

func NewUserId(id string) (*UserId, error) {
	userId := &UserId{userId: id}
	if err := userId.validate(); err != nil {
		return nil, err
	}
	return userId, nil
}

func (UserId *UserId) validate() (err error) {
	defer iterrors.Wrap(&err, "UserId.validate()")
	if UserId.userId == "" {
		return fmt.Errorf("userId is required")
	}
	return nil
}
