package main

import (
	"ed/cmd"
	"ed/internal/app"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	r, err := app.InitEngine()
	if err != nil {
		log.Panic(err)
	}

	srv := new(cmd.Server)
	if err := srv.Run("4041", r); err != nil {
		log.Panic(err)
	}
}
