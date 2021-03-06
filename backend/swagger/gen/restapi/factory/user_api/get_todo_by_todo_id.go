// Code generated by go-swagger; DO NOT EDIT.

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTodoByTodoIDHandlerFunc turns a function with the right signature into a get todo by todo Id handler
type GetTodoByTodoIDHandlerFunc func(GetTodoByTodoIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTodoByTodoIDHandlerFunc) Handle(params GetTodoByTodoIDParams) middleware.Responder {
	return fn(params)
}

// GetTodoByTodoIDHandler interface for that can handle valid get todo by todo Id params
type GetTodoByTodoIDHandler interface {
	Handle(GetTodoByTodoIDParams) middleware.Responder
}

// NewGetTodoByTodoID creates a new http.Handler for the get todo by todo Id operation
func NewGetTodoByTodoID(ctx *middleware.Context, handler GetTodoByTodoIDHandler) *GetTodoByTodoID {
	return &GetTodoByTodoID{Context: ctx, Handler: handler}
}

/* GetTodoByTodoID swagger:route GET /todo/{todo_id} userAPI getTodoByTodoId

TodoIDによって、ユーザー情報を取得する

*/
type GetTodoByTodoID struct {
	Context *middleware.Context
	Handler GetTodoByTodoIDHandler
}

func (o *GetTodoByTodoID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTodoByTodoIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
