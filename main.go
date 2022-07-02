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

	db, err := bootstrap.DBConnect(c)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := bootstrap.NewRedisClient()
	if err != nil {
		log.Println("Unable to connect to a cache service")
	}

	repo := repository.New(db, client)

	serv := service.New(repo)

	r := api.NewServer(serv, c.HTTPPort)

	err = r.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
