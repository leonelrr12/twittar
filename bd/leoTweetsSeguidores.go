package bd

import (
	"context"
	"time"

	"github.com/leonelrr12/twittar/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores Funcion para Leer todos los Tweets de Seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	conds := make([]bson.M, 0)
	conds = append(conds, bson.M{"$match": bson.M{"usuarioid": ID}})
	conds = append(conds, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "usuariorelacionid",
			"foreingField": "userid",
			"as":           "tweet",
		}})
	conds = append(conds, bson.M{"$unwind": "$tweet"})
	conds = append(conds, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	conds = append(conds, bson.M{"$skip": skip})
	conds = append(conds, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conds)
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
