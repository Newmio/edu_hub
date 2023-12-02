package main

import (
	"ed/internal/configs"
	"ed/pkg/server"
	"log"
	_"github.com/lib/pq"
)

func main() {
	r, err :=  configs.InitEngine()
	if err != nil{
		log.Panic(err)
	}

	srv := new(server.Server)
	if err := srv.Run("4041", r); err != nil {
		log.Panic(err)
	}
}
