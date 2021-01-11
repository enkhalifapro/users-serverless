package users

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Provider contains user management functionalities
type Provider struct {
	db dynamoConnector
}

type dynamoConnector interface {
	Put(tableName string, item interface{}) error
	Get(tableName string, item interface{}) (interface{}, error)
}

func NewProvider(db dynamoConnector) *Provider {
	return &Provider{
		db,
	}
}

func (p *Provider) Add(user *User) error {
	user.ID = fmt.Sprintf("%v", time.Now().Unix())
	return p.db.Put("users", user)
}

func (p *Provider) Get(key string) (*User, error) {
	res, err := p.db.Get("users", map[string]interface{}{"userId": key})
	if err != nil {
		return nil, err
	}

	x := res.(map[string]*dynamodb.AttributeValue)
	return &User{ID: x["userId"].String(), Username: x["Username"].String(), Password: x["Password"].String()}, nil
}
