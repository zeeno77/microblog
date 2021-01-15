package main

import (
	"log"

	"github.com/zeeno77/microblog/db"
	"github.com/zeeno77/microblog/handlers"
	//"microblog/db"
	//"microblog/handlers"
)

func main() {
	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
