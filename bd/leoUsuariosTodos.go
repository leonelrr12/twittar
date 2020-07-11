package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/leonelrr12/twittar/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos Funcion para leer todos los Usuarios*/
func LeoUsuariosTodos(ID string, page int64, search string, tipoSearch string) ([]*models.Usuario, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCnn.Database("twittar")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOpciones := options.Find()
	findOpciones.SetSkip((page - 1) * 20)
	findOpciones.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOpciones)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err = cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false
		// Esra busqueda no esta trabajando.  Tengo que revisar ConsultaRelacion()
		encontrado, err = ConsultoRelacion(r)
		if tipoSearch == "new" && encontrado == false {
			incluir = true
		}
		if tipoSearch == "follow" && encontrado == true {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.Ubicacion = ""
			s.SitioWeb = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Printf(err.Error())
		return results, false
	}
	cur.Close(ctx)

	return results, true
}
