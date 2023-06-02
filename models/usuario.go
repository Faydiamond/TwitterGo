package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario en la base de datos TwitetterrGo de Mongodb */
type Usuario struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"  json:"id"`        //porque es un tipo de dato d emongo,_id convencion mongo
	Nombre    string             `bson:"nombre"  json:"nombre,omitempty"` //salida al navegador
	Apellidos string             `bson:"apellidos"  json:"apellidos,omitempty"`
	Fecha     time.Time          `bson:"fechaNacimiento"  json:"fechaNacimiento,omitempty"`
	Email     string             `bson:"email"  json:"email"`                 //siempre lo retorna
	Password  string             `bson:"password"  json:"password,omitempty"` //nunca lo retorno
	Avatar    string             `bson:"avatar"  json:"avatar,omitempty"`
	Banner    string             `bson:"banner"  json:"banner,omitempty"`
	Biografia string             `bson:"ubicacion"  json:"ubicacion,omitempty"`
	SitioWeb  string             `bson:"sitioweb"  json:"sitioWeb,omitempty"`
}
