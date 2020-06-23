package models

/*MsgTweet es el modelo para los Tweets en la DB*/
type MsgTweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
