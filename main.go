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

	result := c.Validate()
	if result != nil {
		log.Fatalln(result)
	}

	db, err := bootstrap.DBConnect(c)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := bootstrap.NewRedisClient(c.RedisPort)
	if err != nil {
		log.Println("Unable to connect to a cache service")
	}

	repo := repository.New(db, client)

	serv := service.New(repo)

	server := api.NewServer(c.HTTPPort)

	err = server.Start(serv)
	if err != nil {
		log.Fatal(err)
	}
}
