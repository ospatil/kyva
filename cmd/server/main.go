package main

import (
	"log"

	"github.com/ospatil/kyva/internal/server"
)

func main() {
	srv := server.NewHttpServer(":7070")
	log.Fatal(srv.ListenAndServe())
}
