package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewUserId(t *testing.T) {
	got, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}

	want := UserId("id")
	if !cmp.Equal(*got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
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
