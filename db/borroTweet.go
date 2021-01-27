package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeeno77/microblog/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ModificoRegistro permite modificar el perfil del tweet*/
func BorroTweet(twt models.Tweet, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if twt.Userid.Hex() != ID {
		println("Ese tweet no le pertenece")
		return false, errors.New("not allowed")
	}
	db := MongoCN.Database(DBName)
	col := db.Collection("tweet")

	filtro := bson.M{"_id": bson.M{"$eq": twt.ID}}
	//Do not delete, just print on screen
	//deletion is working
	_, err := col.DeleteOne(ctx, filtro)
	if err != nil {
		return false, err
	}
	fmt.Println("Tweet borrado con exito")

	return true, nil

}
