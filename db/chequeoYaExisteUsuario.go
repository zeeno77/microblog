package db

import (
	"context"
	"time"

	"microblog/models"

	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya está en la db*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("microblog")
	col := db.Collection("usuarios")

	// M es una función que formatea o mapea a bson lo que recibe como json
	condicion := bson.M{"email": email}

	// en la variable resultado voy a modelar un usuario
	var resultado models.Usuario

	//FindOne me devuelve un sólo registro que cumple con la condición
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
