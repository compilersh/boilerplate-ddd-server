package main

import (
	"log"
	"time"

	"github.com/compilersh/boilerplate-ddd-server/repository"
	"github.com/compilersh/boilerplate-ddd-server/server"
	"github.com/compilersh/boilerplate-ddd-server/user"
)

func main() {

	cfg := Config{
		Port: "9292",
		UserConfig: user.Config{
			SomeValue: "some value",
		},
	}

	repo := &repository.InMemDB{}

	userService := user.NewUserService(cfg.UserConfig, repo)

	userHandler := user.NewHandler(userService)

	router := server.NewRouter(userHandler)

	srv := server.NewServer(":"+cfg.Port, router, server.WithTimeout(10 * time.Second))

	log.Fatal(srv.ListenAndServe())
}

// This is the config struct that initializes the app
// If the config grows, we can move it to its own file
// and consider using some form of config file (yaml, json, etc)
type Config struct {
	Port       string
	UserConfig user.Config
}


