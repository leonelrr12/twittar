package bd

import (
	"context"
	"time"

	"github.com/leonelrr12/twittar/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion Funcion para consultar Relacion en 2 usuarios*/
func ConsultoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("relacion")

	cond := bson.M{
		"usuarioId":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, cond).Decode(&resultado)
	if err != nil {
		return false, err
	}

	return true, nil
}
