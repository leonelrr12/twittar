package bd

import (
	"context"
	"time"

	"github.com/leonelrr12/twittar/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ExisteUsuario Funcion que valida la duplicidad del Usuario por Email*/
func ExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("usuarios")

	cond := bson.M{"email": email}

	var result models.Usuario
	err := col.FindOne(ctx, cond).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
