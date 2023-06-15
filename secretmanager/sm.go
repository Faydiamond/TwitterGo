package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/Faydiamond/TwitterGo/awsgo"
	"github.com/Faydiamond/TwitterGo/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println(">Pido secreto: " + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg) //servicio
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("Error en la funcion GetSecret, error ", err)
		return datosSecret, err
	}
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(">Lectura de secreto perfecta: " + secretName)
	return datosSecret, nil
}
