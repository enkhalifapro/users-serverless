package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Provider contains dynamoDB functionalities implementation
type Provider struct {
	db *dynamodb.DynamoDB
}

func NewConnector() *Provider {
	mySession := session.Must(session.NewSession())
	db := dynamodb.New(mySession)

	return &Provider{db: db}
}

// Put item
func (p *Provider) Put(tableName string, item interface{}) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	req, _ := p.db.PutItemRequest(&input)

	return req.Send()
}

// Get item
func (p *Provider) Get(tableName string, item interface{}) (interface{}, error) {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	input := dynamodb.GetItemInput{
		Key:       av,
		TableName: aws.String(tableName),
	}

	req, err := p.db.GetItem(&input)
	if err != nil {
		return nil, err
	}

	return req.Item, nil
}
