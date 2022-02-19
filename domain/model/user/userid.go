package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId string

func NewUserId(id string) (_ *UserId, err error) {
	defer iterrors.Wrap(&err, "userId.NewUserId(%s)", id)
	if id == "" {
		return nil, fmt.Errorf("userId is required")
	}
	userId := UserId(id)
	return &userId, nil
}
