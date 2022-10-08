package main

import (
	"flag"

	"twojsomsiad/config"
	"twojsomsiad/database"
	"twojsomsiad/router"
	"twojsomsiad/utils"
)

func main() {
	err := config.Setup()
	if err != nil {
		panic(err)
	}

	postman := flag.Bool("postman", false, "generate postman collection")
	flag.Parse()
	if *postman {
		utils.GenerateRequestCollections()
		return
	}

	err = database.Setup()
	if err != nil {
		panic(err)
	}

	r := router.Setup(database.GetDB())

	r.Run()
}
