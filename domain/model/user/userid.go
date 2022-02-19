package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId struct {
	id string
}

func NewUserId(id string) (*UserId, error) {
	userId := new(UserId)
	if err := userId.setId(id); err != nil {
		return nil, err
	}
	return userId, nil
}

func (userId *UserId) setId(id string) (err error) {
	defer iterrors.Wrap(&err, "userId.setId()")
	if id == "" {
		return fmt.Errorf("userId is required")
	}
	userId.id = id
	return nil
}
