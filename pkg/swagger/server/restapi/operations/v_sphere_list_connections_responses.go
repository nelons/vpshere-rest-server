// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// VSphereListConnectionsOKCode is the HTTP code returned for type VSphereListConnectionsOK
const VSphereListConnectionsOKCode int = 200

/*
VSphereListConnectionsOK Success.

swagger:response vSphereListConnectionsOK
*/
type VSphereListConnectionsOK struct {

	/*
	  In: Body
	*/
	Payload []*VSphereListConnectionsOKBodyItems0 `json:"body,omitempty"`
}

// NewVSphereListConnectionsOK creates VSphereListConnectionsOK with default headers values
func NewVSphereListConnectionsOK() *VSphereListConnectionsOK {

	return &VSphereListConnectionsOK{}
}

// WithPayload adds the payload to the v sphere list connections o k response
func (o *VSphereListConnectionsOK) WithPayload(payload []*VSphereListConnectionsOKBodyItems0) *VSphereListConnectionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere list connections o k response
func (o *VSphereListConnectionsOK) SetPayload(payload []*VSphereListConnectionsOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereListConnectionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*VSphereListConnectionsOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// VSphereListConnectionsUnauthorizedCode is the HTTP code returned for type VSphereListConnectionsUnauthorized
const VSphereListConnectionsUnauthorizedCode int = 401

/*
VSphereListConnectionsUnauthorized Invalid Token was passed, this is an unauthenicated request.

swagger:response vSphereListConnectionsUnauthorized
*/
type VSphereListConnectionsUnauthorized struct {
}

// NewVSphereListConnectionsUnauthorized creates VSphereListConnectionsUnauthorized with default headers values
func NewVSphereListConnectionsUnauthorized() *VSphereListConnectionsUnauthorized {

	return &VSphereListConnectionsUnauthorized{}
}

// WriteResponse to the client
func (o *VSphereListConnectionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
