package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/enkhalifapro/users-serverless/pkg/aws/dynamo"

	"github.com/enkhalifapro/users-serverless/internal/users"

	"goa.design/goa/v3/middleware"

	"github.com/enkhalifapro/users-serverless/api/users/gen/http/users_api/server"
	usersapi "github.com/enkhalifapro/users-serverless/api/users/gen/users_api"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
)

func BuildHTTPHandler(logger *zap.Logger) http.Handler {

	// Initialize service dependencies such as databases.
	db := dynamo.NewConnector()
	usrProvider := users.NewProvider(db)

	// Initialize services.
	userService := NewHandler(usrProvider)

	// Initialize Service EndPoints.
	endPoints := usersapi.NewEndpoints(userService)

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	mux := goahttp.NewMuxer()

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	eh := errorHandler(logger)

	userServer := server.New(endPoints, mux, dec, enc, eh, nil)

	// Configure the mux.
	server.Mount(mux, userServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		//handler = httpmdlwr.RequestID()(handler)
	}

	return handler
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *zap.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Error("ERROR:", zap.String("ID", id),
			zap.String("Message", err.Error()))
	}
}

// Handler contains users API handling functionalities
type Handler struct {
	provider provider
}

type provider interface {
	Add(user *users.User) error
	Get(key string) (*users.User, error)
}

func NewHandler(provider provider) *Handler {
	return &Handler{provider: provider}
}

// Add a new user.
func (h *Handler) Add(_ context.Context, usr *usersapi.User) (err error) {
	return h.provider.Add(&users.User{
		Username: usr.Username,
		Password: usr.Password,
	})
}

// Get a user by id.
func (h *Handler) Get(_ context.Context, in *usersapi.GetPayload) (res *usersapi.User, err error) {
	fmt.Println("xxxxxxxx22222222333333Kader")
	fmt.Printf("user isssss %s\n", *in.ID)
	/*usr, err := h.provider.Get(*in.ID)
	if err != nil {
		return nil, err
	}*/

	return &usersapi.User{Username: "usr.Username", Password: "usr.Password"}, nil
}
