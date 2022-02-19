package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewUserId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := "id"
		got, err := NewUserId(id)
		if err != nil {
			t.Fatal(err)
		}
		want := &UserId{id: "id"}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail userId is empty", func(t *testing.T) {
		userId, err := NewUserId("")

		want := "userId.setId(): userId is required"
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}

		if userId != nil {
			t.Errorf("userId must be nil, but %v", userId)
		}
	})
}
