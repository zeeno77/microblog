package main

import (
	"log"

	"github.com/zeeno77/microblog/bd"
	"github.com/zeeno77/microblog/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
