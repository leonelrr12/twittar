package models

/*Relacion Modelo para crear las relaciones de seguidores*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId`
}
