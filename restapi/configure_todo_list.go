// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/kappa-lab/go-swagger-playground/models"
	"github.com/kappa-lab/go-swagger-playground/restapi/operations"
	"github.com/kappa-lab/go-swagger-playground/restapi/operations/todos"
)

var (
	items     = make(map[int64]*models.Item)
	lastID    int64
	itemsLock = &sync.Mutex{}
)

//go:generate swagger generate server --target ../../go-swagger-playground --name TodoList --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.TodoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams) middleware.Responder {
		mergParam := todos.NewFindTodosParams()
		mergParam.Since = swag.Int64(0)
		if params.Limit != nil {
			mergParam.Limit = params.Limit
		}
		if params.Since != nil {
			mergParam.Since = params.Since
		}
		return todos.NewFindTodosOK().
			WithPayload(
				getItems(*mergParam.Since, *mergParam.Limit))
	})

	api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params todos.AddOneParams) middleware.Responder {
		if params.Body == nil {
			return todos.NewAddOneDefault(500).WithPayload(
				&models.Error{Code: 500, Message: swag.String("item must be present")})
		}
		item, _ := addItem(params.Body)
		return todos.NewAddOneCreated().WithPayload(item)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func getItems(since int64, limit int32) (result []*models.Item) {
	result = make([]*models.Item, 0)
	for id, item := range items {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, item)
		}
	}
	return
}

func addItem(item *models.Item) (*models.Item, error) {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	id := atomic.AddInt64(&lastID, 1)
	newItem := &models.Item{
		Description: item.Description,
		ID:          id,
		Completed:   false,
	}

	items[id] = newItem

	return newItem, nil
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
