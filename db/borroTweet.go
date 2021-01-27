package db

import (
	"context"
	"time"

	"github.com/zeeno77/microblog/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ModificoRegistro permite modificar el perfil del tweet*/
//func BorroTweet(twt models.Tweet, ID string) (bool, error) {
func BorroTweet(twt models.Tweet) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	col := db.Collection("tweet")

	/* supongo que me van a enviar de a un campo a modificar a la vez,
	   por eso me fijo si lo que viene tiene valor (largo mayor a cero)
	   Creamos un mapa de interfaces para armar el registro de actualización a la bd
	   Le pongo la información que hay que modificar*/
	//	registro := make(map[string]interface{})

	//	if len(twt.Nombre) > 0 {
	//		registro["nombre"] = twt.Nombre
	//	}
	//	if len(twt.Apellidos) > 0 {
	//		registro["apellidos"] = twt.Apellidos
	//	}
	//	registro["fechaNacimiento"] = twt.FechaNacimiento
	//	if len(twt.Avatar) > 0 {
	//		registro["avatar"] = twt.Avatar
	//	}
	//	if len(twt.Banner) > 0 {
	//		registro["banner"] = twt.Banner
	//	}
	//	if len(twt.Biografia) > 0 {
	//		registro["biografia"] = twt.Biografia
	//	}
	//	if len(twt.Ubicacion) > 0 {
	//		registro["ubicacion"] = twt.Ubicacion
	//	}
	//	if len(twt.SitioWeb) > 0 {
	//		registro["sitioWeb"] = twt.SitioWeb
	//	}
	//	updtString := bson.M{
	//		"$set": registro,
	//	}

	//	if len(twt.userId) > 0 {
	//		if ID == twt.userId {
	//			println("Este tweet es mio")
	//		} //si el tweet es mio
	//	}

	//objID, _ := primitive.ObjectIDFromHex(twt.ID)
	// ahora tengo que realizar un filtro con el ID

	println("******************")
	println(twt.ID.Hex())
	println("******************")
	filtro := bson.M{"_id": bson.M{"$eq": twt.ID}}

	_, err := col.DeleteOne(ctx, filtro)
	if err != nil {
		return false, err
	}
	return true, nil

}
