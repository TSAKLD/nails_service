package main

import (
	_ "github.com/lib/pq"
	"log"
	"nails/api"
	"nails/bootstrap"
	"nails/repository"
	"nails/service"
)

func main() {
	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.DBConnect(c)
	if err != nil {
		log.Fatalln(err)
	}

	client := repository.NewRedisClient()

	repo := repository.New(db, client)

	serv := service.New(repo)

	r := api.NewServer(serv, c.HTTPPort)

	err = r.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
