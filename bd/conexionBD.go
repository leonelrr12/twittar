package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCnn es el objeto de conexion a la DB*/
var MongoCnn = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:root@cluster0-uqrhw.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConnectDB es la funcion que me permite conectar al la DB */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la DB")
	return client
}

/*CheckConnect es la funcion que me permite validar si la DB esta arriba */
func CheckConnect() bool {
	err := MongoCnn.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
