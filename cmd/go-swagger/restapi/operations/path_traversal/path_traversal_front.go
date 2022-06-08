// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PathTraversalFrontHandlerFunc turns a function with the right signature into a path traversal front handler
type PathTraversalFrontHandlerFunc func(PathTraversalFrontParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PathTraversalFrontHandlerFunc) Handle(params PathTraversalFrontParams) middleware.Responder {
	return fn(params)
}

// PathTraversalFrontHandler interface for that can handle valid path traversal front params
type PathTraversalFrontHandler interface {
	Handle(PathTraversalFrontParams) middleware.Responder
}

// NewPathTraversalFront creates a new http.Handler for the path traversal front operation
func NewPathTraversalFront(ctx *middleware.Context, handler PathTraversalFrontHandler) *PathTraversalFront {
	return &PathTraversalFront{Context: ctx, Handler: handler}
}

/* PathTraversalFront swagger:route GET /pathTraversal path_traversal pathTraversalFront

front page of the Path Traversal vulnerability

*/
type PathTraversalFront struct {
	Context *middleware.Context
	Handler PathTraversalFrontHandler
}

func (o *PathTraversalFront) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPathTraversalFrontParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
