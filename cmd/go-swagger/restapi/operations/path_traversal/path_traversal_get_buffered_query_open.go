// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PathTraversalGetBufferedQueryOpenHandlerFunc turns a function with the right signature into a path traversal get buffered query open handler
type PathTraversalGetBufferedQueryOpenHandlerFunc func(PathTraversalGetBufferedQueryOpenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PathTraversalGetBufferedQueryOpenHandlerFunc) Handle(params PathTraversalGetBufferedQueryOpenParams) middleware.Responder {
	return fn(params)
}

// PathTraversalGetBufferedQueryOpenHandler interface for that can handle valid path traversal get buffered query open params
type PathTraversalGetBufferedQueryOpenHandler interface {
	Handle(PathTraversalGetBufferedQueryOpenParams) middleware.Responder
}

// NewPathTraversalGetBufferedQueryOpen creates a new http.Handler for the path traversal get buffered query open operation
func NewPathTraversalGetBufferedQueryOpen(ctx *middleware.Context, handler PathTraversalGetBufferedQueryOpenHandler) *PathTraversalGetBufferedQueryOpen {
	return &PathTraversalGetBufferedQueryOpen{Context: ctx, Handler: handler}
}

/* PathTraversalGetBufferedQueryOpen swagger:route GET /pathTraversal/os.Open/buffered-query/{safety} path_traversal pathTraversalGetBufferedQueryOpen

demonstrates Path Traversal via buffered-query, with vulnerable function os.Open

*/
type PathTraversalGetBufferedQueryOpen struct {
	Context *middleware.Context
	Handler PathTraversalGetBufferedQueryOpenHandler
}

func (o *PathTraversalGetBufferedQueryOpen) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPathTraversalGetBufferedQueryOpenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
