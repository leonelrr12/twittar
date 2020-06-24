package bd

import (
	"context"
	"time"

	"github.com/leonelrr12/twittar/models"
)

/*InsertRelacion Funcion para insertar las relaciones de seguidores*/
func InsertRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
