// Code generated by goa v3.2.6, DO NOT EDIT.
//
// usersAPI service
//
// Command:
// $ goa gen github.com/enkhalifapro/users-serverless/internal/users/api/design
// -o ./api/users

package usersapi

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Users API
type Service interface {
	// Add a new user.
	Add(context.Context, *User) (err error)
	// Get user by id.
	Get(context.Context, *GetPayload) (res *User, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "usersAPI"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"add", "get"}

// User is the payload type of the usersAPI service add method.
type User struct {
	Username string
	Password string
}

// GetPayload is the payload type of the usersAPI service get method.
type GetPayload struct {
	// user id
	ID *string
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "InternalError",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
