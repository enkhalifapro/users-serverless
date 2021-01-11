package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

// DataTypes

// Index type
var User = Type("User", func() {
	Description("Index describes a collection of indices to be stored.")
	Field(1, "username", String)
	Field(2, "password", String)
	Required("username", "password")
})

// Services

var _ = Service("usersAPI", func() {
	Description("Users API")

	HTTP(func() {
		Path("/users")
	})

	Method("add", func() {
		Description("Add a new user.")

		Payload(User)

		Error("InternalError")

		HTTP(func() {
			POST("")
			Response("InternalError", StatusInternalServerError)
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get user by id.")

		Payload(func() {
			Attribute("id", String, "user id")
		})

		Error("InternalError")
		Result(User)

		HTTP(func() {
			GET("/{id}")
			Response("InternalError", StatusInternalServerError)
			Response(StatusOK)
		})
	})
})

// Server

// API describes the global properties of the API server.
var _ = API("usersAPIServer", func() {
	Title("Users API server")
	Description("Users restful API")
	Server("usersAPIServer", func() {
		Description("user hosts the User Service.")

		// List the services hosted by this server.
		Services("usersAPI")

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})

		Host("production", func() {
			Description("Production hosts.")
			// URIs can be parameterized using {param} notation.
			URI("http://0.0.0.0:8000")
			URI("grpc://0.0.0.0:8080")
		})
	})
})
