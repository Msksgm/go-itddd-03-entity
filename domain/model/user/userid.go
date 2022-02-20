package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId struct {
	id string
}

func NewUserId(id string) (_ *UserId, err error) {
	defer iterrors.Wrap(&err, "NewUserId(%s)", id)
	userId := new(UserId)
	if id == "" {
		return nil, fmt.Errorf("userId is required")
	}
	userId.id = id
	return userId, nil
}
