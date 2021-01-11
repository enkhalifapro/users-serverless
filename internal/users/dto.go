package users

type User struct {
	ID       string `dynamodbav:"userId"`
	Username string
	Password string
}
