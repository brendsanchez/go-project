package secret

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awsgo "github.com/brendsanchez/go-project/aws"
	"github.com/brendsanchez/go-project/models"
)

func GetSecret(name string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson

	fmt.Print("> get secret ", name, dataSecret)

	svg := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svg.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})

	if err != nil {
		fmt.Print(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)
	fmt.Print("> secret OK ", name)
	return dataSecret, nil
}
