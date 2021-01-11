package users

import (
	"fmt"
	"time"
)

// Provider contains user management functionalities
type Provider struct {
	db dynamoConnector
}

type dynamoConnector interface {
	Put(tableName string, item interface{}) error
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
