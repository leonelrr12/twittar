package models

/*msgTweet es el modelo para los Tweets en la DB*/
type msgTweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
