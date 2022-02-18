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
}

func TestValidateUserId(t *testing.T) {
	for _, test := range []struct {
		in      string
		wantErr bool
	}{
		{"", true},
		{"id", false},
	} {
		userId := UserId(test.in)
		err := userId.validate(userId)
		if (err != nil) != test.wantErr {
			t.Errorf("ValidateAppVersion(%q) = %v, want error = %t", test.in, err, test.wantErr)
		}
	}
}
