package bd

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/leonelrr12/twittar/models"
)

/*BorroRelacion Funcion para eliminar seguidores*/
func BorroRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("relacion")

	res, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	if res.DeletedCount == 0 {
		log.Println("DeleteOne() document not found:", res)
		return false, errors.New("no se han eliminado registros")
	}

	return true, nil
}
