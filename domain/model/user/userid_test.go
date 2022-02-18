package user

import (
	"reflect"
	"testing"
)

func TestNewUserId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := "id"
		userId, err := NewUserId(id)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(userId.userId, id) {
			t.Errorf("userId.userId %v should be %v", userId.userId, id)
		}
	})
	t.Run("fail userId is empty", func(t *testing.T) {
		userId, err := NewUserId("")
		if (err != nil) != true || userId != nil {
			t.Errorf("empty is not perimetted for userId")
		}
	})
}
