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

		want := "UserId.validate(): userId is required"
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}

		if userId != nil {
			t.Errorf("userId must be nil, but %v", userId)
		}
	})
}
