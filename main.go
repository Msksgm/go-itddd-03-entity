package main

import (
	"fmt"
	"log"

	"github.com/Msksgm/go-itddd-03-entity/domain/entity"
)

func main() {
	var user *entity.User

	userId, err := entity.NewUserId("a")
	if err != nil {
		log.Fatal(err)
	}

	userName, err := entity.NewUserName("name")
	if err != nil {
		log.Fatal(err)
	}

	nerUser, err := user.NewUser(*userId, *userName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nerUser)
}
