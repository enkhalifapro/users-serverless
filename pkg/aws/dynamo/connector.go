package dynamo

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Provider contains dynamoDB functionalities implementation
type Provider struct {
	db *dynamodb.Client
}

func NewConnector() *Provider {
	awsRegion := "us-east-1"
	awsEndpoint := os.Getenv("LOCALSTACK_HOSTNAME")
	fmt.Println("jjjjjjjjj")
	fmt.Println(awsEndpoint)

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				URL:           fmt.Sprintf("http://%s:4566", awsEndpoint),
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	db := dynamodb.NewFromConfig(awsCfg)

	return &Provider{db: db}
}

// Put item
func (p *Provider) Put(tableName string, item interface{}) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	input := dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	res, err := p.db.PutItem(context.Background(), &input)
	if err != nil {
		return err
	}

	fmt.Println("dynamo ressssssssss")
	fmt.Println(res)

	return nil
}

// Get item
func (p *Provider) Get(tableName string, item interface{}) (interface{}, error) {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		fmt.Println("errr11111")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(av)
	input := dynamodb.GetItemInput{
		Key:       av,
		TableName: aws.String(tableName),
	}

	req, err := p.db.GetItem(context.Background(), &input)
	if err != nil {
		return nil, err
	}

	return req.Item, nil
}
