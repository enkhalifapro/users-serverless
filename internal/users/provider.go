package users

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"

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
	fmt.Println("save usessssssrrrrrr")
	fmt.Println(user)
	err := p.db.Put("users", user)
	fmt.Println(err)
	return err
}

func (p *Provider) Get(key string) (*User, error) {
	res, err := p.db.Get("users", map[string]interface{}{"Username": key, "Password": "saber1"})
	if err != nil {
		fmt.Println("errrrrrxxxx")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("tttttttttt")
	fmt.Println(res)
	v:=res.(map[string]types.AttributeValue)
	var x string//:= make(map[string]interface{})
	 err = attributevalue.Unmarshal(v["Username"],&x)
	if err != nil{
		fmt.Println("errrror:")
		fmt.Println(err)
		return  nil,err
	}
	fmt.Println("zzzzzzzzz")
	fmt.Println(x)
	return &User{Username: x, Password: ""}, nil
}
