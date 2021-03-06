// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the usersAPI service.
//
// Command:
// $ goa gen github.com/enkhalifapro/users-serverless/internal/users/api/design
// -o ./api/users

package client

import (
	"fmt"
)

// AddUsersAPIPath returns the URL path to the usersAPI service add HTTP endpoint.
func AddUsersAPIPath() string {
	return "/users"
}

// GetUsersAPIPath returns the URL path to the usersAPI service get HTTP endpoint.
func GetUsersAPIPath(id string) string {
	return fmt.Sprintf("/users/%v", id)
}
