package bd

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BorroTweet Funcion para borrar un Tweet*/
func BorroTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	cond := bson.M{
		"_id":    objID,
		"userId": UserID,
	}

	res, err := col.DeleteOne(ctx, cond)
	if res.DeletedCount == 0 {
		log.Println("DeleteOne() document not found:", res)
		return errors.New("no se han eliminado registros")
	}

	return err
}
