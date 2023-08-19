package main

import (
	"fmt"
	"log"
	"net"

	"KVADO-library/internal/api"
	"KVADO-library/internal/app"
	"KVADO-library/internal/repository"
	"KVADO-library/internal/service"
)

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		log.Fatalf("init config: %s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		err := lis.Close()
		if err != nil {
			log.Printf("close listener: %s/n", err)
		}
	}()

	db, err := app.ConnectToMySQL(cfg.MySQLdsn)
	if err != nil {
		log.Panicf("connect to database: %s", err)
	}
	repo := repository.NewRepository(db)
	s := service.NewBookService(repo)
	h := api.NewHandler(s)
	srv := api.NewServer(h)
	err = srv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
