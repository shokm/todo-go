// Code generated by go-swagger; DO NOT EDIT.

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateTodoByTodoIDOKCode is the HTTP code returned for type UpdateTodoByTodoIDOK
const UpdateTodoByTodoIDOKCode int = 200

/*UpdateTodoByTodoIDOK 成功なレスポンス

swagger:response updateTodoByTodoIdOK
*/
type UpdateTodoByTodoIDOK struct {
}

// NewUpdateTodoByTodoIDOK creates UpdateTodoByTodoIDOK with default headers values
func NewUpdateTodoByTodoIDOK() *UpdateTodoByTodoIDOK {

	return &UpdateTodoByTodoIDOK{}
}

// WriteResponse to the client
func (o *UpdateTodoByTodoIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateTodoByTodoIDBadRequestCode is the HTTP code returned for type UpdateTodoByTodoIDBadRequest
const UpdateTodoByTodoIDBadRequestCode int = 400

/*UpdateTodoByTodoIDBadRequest 無効なTodoID

swagger:response updateTodoByTodoIdBadRequest
*/
type UpdateTodoByTodoIDBadRequest struct {
}

// NewUpdateTodoByTodoIDBadRequest creates UpdateTodoByTodoIDBadRequest with default headers values
func NewUpdateTodoByTodoIDBadRequest() *UpdateTodoByTodoIDBadRequest {

	return &UpdateTodoByTodoIDBadRequest{}
}

// WriteResponse to the client
func (o *UpdateTodoByTodoIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateTodoByTodoIDNotFoundCode is the HTTP code returned for type UpdateTodoByTodoIDNotFound
const UpdateTodoByTodoIDNotFoundCode int = 404

/*UpdateTodoByTodoIDNotFound Todoを見つけていません

swagger:response updateTodoByTodoIdNotFound
*/
type UpdateTodoByTodoIDNotFound struct {
}

// NewUpdateTodoByTodoIDNotFound creates UpdateTodoByTodoIDNotFound with default headers values
func NewUpdateTodoByTodoIDNotFound() *UpdateTodoByTodoIDNotFound {

	return &UpdateTodoByTodoIDNotFound{}
}

// WriteResponse to the client
func (o *UpdateTodoByTodoIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
