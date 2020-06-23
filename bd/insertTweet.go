package bd

import (
	"context"
	"time"

	"github.com/leonelrr12/twittar/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertTweet Funcion para grabar los Tweets en la BD*/
func InsertTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("tweets")

	registro := bson.M{
		"userId":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
