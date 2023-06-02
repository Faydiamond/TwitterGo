package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD() //Ejecuto mi conexion

var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/")

/* ConectarBD permite la conexion con la base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions) //contexto
	if err != nil {
		log.Fatal("Fallo la conexion de la base de datos! ", err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Fallo ping conexion de la base de datos! ", err.Error())
		return client
	}
	log.Println("Conexion exitosas")
	//conexion valida
	return client
}

/*ChequeConnection ping a la bd */
func ChequeConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
