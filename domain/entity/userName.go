package entity

import "fmt"

type userName string

func NewUserName(v string) (*userName, error) {
	userName := userName(v)
	if err := userName.validate(userName); err != nil {
		return nil, err
	}
	return &userName, nil
}

func (userName *userName) validate(v userName) error {
	if v == "" {
		return fmt.Errorf("氏名は必須です")
	}
	if len(v) < 3 {
		return fmt.Errorf("氏名は3文字以上にしてください")
	}
	return nil
}
