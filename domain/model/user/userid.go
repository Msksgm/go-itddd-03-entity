package user

import (
	"fmt"

	"github.com/Msksgm/go-itddd-03-entity/iterrors"
)

type UserId struct {
	userId string
}

func NewUserId(id string) (*UserId, error) {
	if err := validateId(id); err != nil {
		return nil, err
	}
	userId := &UserId{userId: id}
	return userId, nil
}

func validateId(id string) (err error) {
	defer iterrors.Wrap(&err, "userid.go validateId()")
	if id == "" {
		return fmt.Errorf("userId is required")
	}
	return nil
}
