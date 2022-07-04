// Code generated by go-swagger; DO NOT EDIT.

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shokm/todo-go/backend/swagger/gen/models"
)

// GetTodoByTodoIDOKCode is the HTTP code returned for type GetTodoByTodoIDOK
const GetTodoByTodoIDOKCode int = 200

/*GetTodoByTodoIDOK 成功なレスポンス

swagger:response getTodoByTodoIdOK
*/
type GetTodoByTodoIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetTodoByTodoIDOK creates GetTodoByTodoIDOK with default headers values
func NewGetTodoByTodoIDOK() *GetTodoByTodoIDOK {

	return &GetTodoByTodoIDOK{}
}

// WithPayload adds the payload to the get todo by todo Id o k response
func (o *GetTodoByTodoIDOK) WithPayload(payload *models.User) *GetTodoByTodoIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo by todo Id o k response
func (o *GetTodoByTodoIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoByTodoIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTodoByTodoIDBadRequestCode is the HTTP code returned for type GetTodoByTodoIDBadRequest
const GetTodoByTodoIDBadRequestCode int = 400

/*GetTodoByTodoIDBadRequest 無効なTodoID

swagger:response getTodoByTodoIdBadRequest
*/
type GetTodoByTodoIDBadRequest struct {
}

// NewGetTodoByTodoIDBadRequest creates GetTodoByTodoIDBadRequest with default headers values
func NewGetTodoByTodoIDBadRequest() *GetTodoByTodoIDBadRequest {

	return &GetTodoByTodoIDBadRequest{}
}

// WriteResponse to the client
func (o *GetTodoByTodoIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetTodoByTodoIDNotFoundCode is the HTTP code returned for type GetTodoByTodoIDNotFound
const GetTodoByTodoIDNotFoundCode int = 404

/*GetTodoByTodoIDNotFound Todoを見つけていません

swagger:response getTodoByTodoIdNotFound
*/
type GetTodoByTodoIDNotFound struct {
}

// NewGetTodoByTodoIDNotFound creates GetTodoByTodoIDNotFound with default headers values
func NewGetTodoByTodoIDNotFound() *GetTodoByTodoIDNotFound {

	return &GetTodoByTodoIDNotFound{}
}

// WriteResponse to the client
func (o *GetTodoByTodoIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
