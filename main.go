package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Msksgm/go-itddd-03-entity/domain/model/user"
)

var (
	id   = flag.String("id", "", "id of user")
	name = flag.String("name", "", "name of user")
)

func main() {
	flag.Parse()
	userId, err := user.NewUserId(*id)
	if err != nil {
		log.Fatal(err)
	}

	newUser, err := user.NewUser(*userId, *name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUser)
}
