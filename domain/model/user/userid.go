package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId struct {
	userId string
}

func NewUserId(id string) (*UserId, error) {
	userId := new(UserId)
	if err := userId.setUserId(id); err != nil {
		return nil, err
	}
	return userId, nil
}

func (userId *UserId) setUserId(id string) (err error) {
	defer iterrors.Wrap(&err, "userId.setUserId()")
	if id == "" {
		return fmt.Errorf("userId is required")
	}
	userId.userId = id
	return nil
}
