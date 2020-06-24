package bd

import (
	"context"
	"log"
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
		"usuarioid":          t.UsuarioID,
		"usuariorrelacionid": t.UsuarioRelacionID,
	}
	log.Println("cond: ", cond)

	var resultado models.Relacion
	err := col.FindOne(ctx, cond).Decode(&resultado)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}
