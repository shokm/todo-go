// Code generated by go-swagger; DO NOT EDIT.

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateTodoByTodoIDHandlerFunc turns a function with the right signature into a update todo by todo Id handler
type UpdateTodoByTodoIDHandlerFunc func(UpdateTodoByTodoIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateTodoByTodoIDHandlerFunc) Handle(params UpdateTodoByTodoIDParams) middleware.Responder {
	return fn(params)
}

// UpdateTodoByTodoIDHandler interface for that can handle valid update todo by todo Id params
type UpdateTodoByTodoIDHandler interface {
	Handle(UpdateTodoByTodoIDParams) middleware.Responder
}

// NewUpdateTodoByTodoID creates a new http.Handler for the update todo by todo Id operation
func NewUpdateTodoByTodoID(ctx *middleware.Context, handler UpdateTodoByTodoIDHandler) *UpdateTodoByTodoID {
	return &UpdateTodoByTodoID{Context: ctx, Handler: handler}
}

/* UpdateTodoByTodoID swagger:route PUT /todo/{todo_id} userAPI updateTodoByTodoId

TodoIDによって、ユーザー情報を更新する

*/
type UpdateTodoByTodoID struct {
	Context *middleware.Context
	Handler UpdateTodoByTodoIDHandler
}

func (o *UpdateTodoByTodoID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateTodoByTodoIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
