package entity

import (
	"fmt"
)

type userId string

func NewUserId(v string) (*userId, error) {
	userId := userId(v)
	if err := userId.validate(userId); err != nil {
		return nil, err
	}
	return &userId, nil
}

func (userId *userId) validate(id userId) error {
	if id == "" {
		return fmt.Errorf("ユーザーidは必須です")
	}
	return nil
}
