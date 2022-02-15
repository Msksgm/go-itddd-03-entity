package entity

import "fmt"

type userName string

func NewUserName(v string) (*userName, error) {
	if v == "" {
		return nil, fmt.Errorf("氏名は必須です")
	}
	if len(v) < 3 {
		return nil, fmt.Errorf("氏名は3文字以上にしてください")
	}
	userName := userName(v)
	return &userName, nil
}
