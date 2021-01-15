package main

import (
	"log"

	"github.com/zeeno77/microblog/db"
	"github.com/zeeno77/microblog/handlers"
	//"github.com/zeeno77/microblog/db"
	//"github.com/zeeno77/microblog/handlers"
)

func main() {
	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
