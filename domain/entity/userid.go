package entity

import "fmt"

type userId string

func NewUserId(v string) (*userId, error) {
	if v == "" {
		return nil, fmt.Errorf("newUserId: %v is invalid for newUserID", v)
	}
	userId := userId(v)
	return &userId, nil
}
