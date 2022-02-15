package main

import (
	"fmt"
	"log"

	"github.com/Msksgm/go-itddd-03-entity/domain/entity"
)

func main() {
	var user *entity.User

	rightUserId, err := entity.NewUserId("a")
	if err != nil {
		log.Fatal(err)
	}

	rightUserName, err := entity.NewUserName("right")
	if err != nil {
		log.Fatal(err)
	}

	right, err := user.NewUser(*rightUserId, *rightUserName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(right)

	leftUserId, err := entity.NewUserId("b")
	if err != nil {
		log.Fatal(err)
	}

	left, err := user.NewUser(*leftUserId, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(left)

	fmt.Println(left.Equals(*right))
}
