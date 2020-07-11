package models

import (
	"time"
)

/*Tweet es el modelo para los Tweets en la DB*/
type Tweet struct {
	UserID  string    `bson:"userId" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
