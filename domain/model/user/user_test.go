package user

import (
	"testing"
)

func TestChangeUserName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userId, err := NewUserId("id")
		if err != nil {
			t.Fatal(err)
		}
		name := "name"
		user, err := NewUser(*userId, name)
		if err != nil {
			t.Fatal(err)
		}

		if err := user.ChangeUserName("changedName"); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(user.name, "changedName") {
			t.Errorf("got %v, want changedName", user.name)
		}
	})
}

func TestEquals(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userId1, err := NewUserId("id")
		if err != nil {
			t.Fatal(err)
		}
		name1 := "name"
		user1, err := NewUser(*userId1, name1)
		if err != nil {
			t.Fatal(err)
		}

		user2 := user1
		if !user1.Equals(user2) {
			t.Errorf("%v must be equal to %v", &user1, &user2)
		}

		if err := user1.ChangeUserName("changedName"); err != nil {
			t.Fatal(err)
		}

		if !user1.Equals(user2) {
			t.Errorf("%v must be equal to %v", &user1, &user2)
		}
	})
	t.Run("fail", func(t *testing.T) {
		userId1, err := NewUserId("id1")
		if err != nil {
			t.Fatal(err)
		}
		name1 := "name1"
		user1, err := NewUser(*userId1, name1)
		if err != nil {
			t.Fatal(err)
		}

		userId2, err := NewUserId("id2")
		if err != nil {
			t.Fatal(err)
		}
		name2 := "name2"
		user2, err := NewUser(*userId2, name2)
		if err != nil {
			t.Fatal(err)
		}

		if user1.Equals(user2) {
			t.Errorf("%v must not be equal to %v", &user1, &user2)
		}
	})
}