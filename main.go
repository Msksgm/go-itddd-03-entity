package main

import (
	"fmt"
	"log"

	"github.com/Msksgm/go-itddd-03-entity/domain/model/user"
)

func main() {
	userId, err := user.NewUserId("a")
	if err != nil {
		log.Fatal(err)
	}

	newUser, err := user.NewUser(*userId, "name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUser)
}
