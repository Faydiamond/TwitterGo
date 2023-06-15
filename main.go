package main

import (
	"context"
	"os"
	"strings"

	"github.com/Faydiamond/TwitterGo/awsgo"
	"github.com/Faydiamond/TwitterGo/bd"
	"github.com/Faydiamond/TwitterGo/handlers"
	"github.com/Faydiamond/TwitterGo/models"
	"github.com/Faydiamond/TwitterGo/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//aqui desarrollo mi funcion lambda
	lambda.Start(EjecutarLambda)
}

/*   */
func EjecutarLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.IniciarAws()
	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entrono, deben incluir: SecretName, BucketName, UrlPrefix  ",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Erroen la lectura se Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	path := strings.Replace(request.PathParameters["Twiter"], os.Getenv("UrlPrefix"), "", -1) //palabra truncar
	//context
	//establezco los valores de mi secret
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)                          //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)          //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)          //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)      //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)              //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)      //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)        //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)                  //
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName")) //

	//Chqqueo conexion bd
	err = bd.ConexionBD(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Erro conectando en la base de datos " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	respApi := handlers.Manejadores(awsgo.Ctx, request)
	if respApi.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       string(respApi.Message),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return respApi.CustomResp, nil
	}

}

func ValidoParametros() bool {
	//lambda recibe tres variables de entorno
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
