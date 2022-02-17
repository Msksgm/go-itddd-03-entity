package main

import (
	"fmt"
	"log"

	"github.com/Msksgm/go-itddd-03-entity/domain/entity"
)

func main() {
	userId, err := entity.NewUserId("a")
	if err != nil {
		log.Fatal(err)
	}

	newUser, err := entity.NewUser(*userId, "name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUser)
}
