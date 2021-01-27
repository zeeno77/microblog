package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Tweet captura del Body, el mensaje que nos llega*/
type Tweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Userid  primitive.ObjectID `bson:"_id,omitempty" json:"userId"`
	Mensaje string             `bson:"mensaje" json:"mensaje"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
