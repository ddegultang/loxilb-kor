// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigBgpNeighAllOKCode is the HTTP code returned for type GetConfigBgpNeighAllOK
const GetConfigBgpNeighAllOKCode int = 200

/*
GetConfigBgpNeighAllOK OK

swagger:response getConfigBgpNeighAllOK
*/
type GetConfigBgpNeighAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigBgpNeighAllOKBody `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllOK creates GetConfigBgpNeighAllOK with default headers values
func NewGetConfigBgpNeighAllOK() *GetConfigBgpNeighAllOK {

	return &GetConfigBgpNeighAllOK{}
}

// WithPayload adds the payload to the get config bgp neigh all o k response
func (o *GetConfigBgpNeighAllOK) WithPayload(payload *GetConfigBgpNeighAllOKBody) *GetConfigBgpNeighAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all o k response
func (o *GetConfigBgpNeighAllOK) SetPayload(payload *GetConfigBgpNeighAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpNeighAllNoContentCode is the HTTP code returned for type GetConfigBgpNeighAllNoContent
const GetConfigBgpNeighAllNoContentCode int = 204

/*
GetConfigBgpNeighAllNoContent OK

swagger:response getConfigBgpNeighAllNoContent
*/
type GetConfigBgpNeighAllNoContent struct {
}

// NewGetConfigBgpNeighAllNoContent creates GetConfigBgpNeighAllNoContent with default headers values
func NewGetConfigBgpNeighAllNoContent() *GetConfigBgpNeighAllNoContent {

	return &GetConfigBgpNeighAllNoContent{}
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetConfigBgpNeighAllBadRequestCode is the HTTP code returned for type GetConfigBgpNeighAllBadRequest
const GetConfigBgpNeighAllBadRequestCode int = 400

/*
GetConfigBgpNeighAllBadRequest Malformed arguments for API call

swagger:response getConfigBgpNeighAllBadRequest
*/
type GetConfigBgpNeighAllBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllBadRequest creates GetConfigBgpNeighAllBadRequest with default headers values
func NewGetConfigBgpNeighAllBadRequest() *GetConfigBgpNeighAllBadRequest {

	return &GetConfigBgpNeighAllBadRequest{}
}

// WithPayload adds the payload to the get config bgp neigh all bad request response
func (o *GetConfigBgpNeighAllBadRequest) WithPayload(payload *models.Error) *GetConfigBgpNeighAllBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all bad request response
func (o *GetConfigBgpNeighAllBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpNeighAllUnauthorizedCode is the HTTP code returned for type GetConfigBgpNeighAllUnauthorized
const GetConfigBgpNeighAllUnauthorizedCode int = 401

/*
GetConfigBgpNeighAllUnauthorized Invalid authentication credentials

swagger:response getConfigBgpNeighAllUnauthorized
*/
type GetConfigBgpNeighAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllUnauthorized creates GetConfigBgpNeighAllUnauthorized with default headers values
func NewGetConfigBgpNeighAllUnauthorized() *GetConfigBgpNeighAllUnauthorized {

	return &GetConfigBgpNeighAllUnauthorized{}
}

// WithPayload adds the payload to the get config bgp neigh all unauthorized response
func (o *GetConfigBgpNeighAllUnauthorized) WithPayload(payload *models.Error) *GetConfigBgpNeighAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all unauthorized response
func (o *GetConfigBgpNeighAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpNeighAllForbiddenCode is the HTTP code returned for type GetConfigBgpNeighAllForbidden
const GetConfigBgpNeighAllForbiddenCode int = 403

/*
GetConfigBgpNeighAllForbidden Capacity insufficient

swagger:response getConfigBgpNeighAllForbidden
*/
type GetConfigBgpNeighAllForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllForbidden creates GetConfigBgpNeighAllForbidden with default headers values
func NewGetConfigBgpNeighAllForbidden() *GetConfigBgpNeighAllForbidden {

	return &GetConfigBgpNeighAllForbidden{}
}

// WithPayload adds the payload to the get config bgp neigh all forbidden response
func (o *GetConfigBgpNeighAllForbidden) WithPayload(payload *models.Error) *GetConfigBgpNeighAllForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all forbidden response
func (o *GetConfigBgpNeighAllForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpNeighAllNotFoundCode is the HTTP code returned for type GetConfigBgpNeighAllNotFound
const GetConfigBgpNeighAllNotFoundCode int = 404

/*
GetConfigBgpNeighAllNotFound Resource not found

swagger:response getConfigBgpNeighAllNotFound
*/
type GetConfigBgpNeighAllNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllNotFound creates GetConfigBgpNeighAllNotFound with default headers values
func NewGetConfigBgpNeighAllNotFound() *GetConfigBgpNeighAllNotFound {

	return &GetConfigBgpNeighAllNotFound{}
}

// WithPayload adds the payload to the get config bgp neigh all not found response
func (o *GetConfigBgpNeighAllNotFound) WithPayload(payload *models.Error) *GetConfigBgpNeighAllNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all not found response
func (o *GetConfigBgpNeighAllNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpNeighAllConflictCode is the HTTP code returned for type GetConfigBgpNeighAllConflict
const GetConfigBgpNeighAllConflictCode int = 409

/*
GetConfigBgpNeighAllConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response getConfigBgpNeighAllConflict
*/
type GetConfigBgpNeighAllConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllConflict creates GetConfigBgpNeighAllConflict with default headers values
func NewGetConfigBgpNeighAllConflict() *GetConfigBgpNeighAllConflict {

	return &GetConfigBgpNeighAllConflict{}
}

// WithPayload adds the payload to the get config bgp neigh all conflict response
func (o *GetConfigBgpNeighAllConflict) WithPayload(payload *models.Error) *GetConfigBgpNeighAllConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all conflict response
func (o *GetConfigBgpNeighAllConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpNeighAllInternalServerErrorCode is the HTTP code returned for type GetConfigBgpNeighAllInternalServerError
const GetConfigBgpNeighAllInternalServerErrorCode int = 500

/*
GetConfigBgpNeighAllInternalServerError Internal service error

swagger:response getConfigBgpNeighAllInternalServerError
*/
type GetConfigBgpNeighAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllInternalServerError creates GetConfigBgpNeighAllInternalServerError with default headers values
func NewGetConfigBgpNeighAllInternalServerError() *GetConfigBgpNeighAllInternalServerError {

	return &GetConfigBgpNeighAllInternalServerError{}
}

// WithPayload adds the payload to the get config bgp neigh all internal server error response
func (o *GetConfigBgpNeighAllInternalServerError) WithPayload(payload *models.Error) *GetConfigBgpNeighAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all internal server error response
func (o *GetConfigBgpNeighAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpNeighAllServiceUnavailableCode is the HTTP code returned for type GetConfigBgpNeighAllServiceUnavailable
const GetConfigBgpNeighAllServiceUnavailableCode int = 503

/*
GetConfigBgpNeighAllServiceUnavailable Maintanence mode

swagger:response getConfigBgpNeighAllServiceUnavailable
*/
type GetConfigBgpNeighAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpNeighAllServiceUnavailable creates GetConfigBgpNeighAllServiceUnavailable with default headers values
func NewGetConfigBgpNeighAllServiceUnavailable() *GetConfigBgpNeighAllServiceUnavailable {

	return &GetConfigBgpNeighAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config bgp neigh all service unavailable response
func (o *GetConfigBgpNeighAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigBgpNeighAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp neigh all service unavailable response
func (o *GetConfigBgpNeighAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpNeighAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}