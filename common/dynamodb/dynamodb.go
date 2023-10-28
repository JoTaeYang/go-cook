package dynamodb

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type Config struct {
	Region    string
	AccessKey string
	SecretKey string
	EndPoint  string
	Table     string
}

var (
	boscoTableName string
	dynamodbconn   *dynamodb.Client
)

func InitDynamoDB(cfg Config) {
	log.Println("dynamodb init")
	boscoTableName = cfg.Table

	if cfg.AccessKey != "" {
		awsConfig, err := config.LoadDefaultConfig(
			context.Background(),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")),
			config.WithRegion(cfg.Region),
		)

		if err != nil {
			log.Println(err)
			return
		}
		dynamodbconn = dynamodb.NewFromConfig(awsConfig)
	}
}

func GetItemPKSK(pk, sk string) interface{} {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(boscoTableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pk},
			"SK": &types.AttributeValueMemberS{Value: sk},
		},
	}
	result, err := dynamodbconn.GetItem(context.Background(), params)
	if err != nil {
		log.Println(err)
		return nil
	}

	if result.Item != nil {
		log.Println(result.Item)
	}

	return nil
}
