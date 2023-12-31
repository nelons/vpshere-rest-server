// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nelons/vsphere-rest-server/pkg/swagger/server/models"
)

// VSphereGetHostByMoRefOKCode is the HTTP code returned for type VSphereGetHostByMoRefOK
const VSphereGetHostByMoRefOKCode int = 200

/*
VSphereGetHostByMoRefOK Successful request. Returns JSON containing a count and the collection of objects.

swagger:response vSphereGetHostByMoRefOK
*/
type VSphereGetHostByMoRefOK struct {

	/*
	  In: Body
	*/
	Payload *models.ObjectCollection `json:"body,omitempty"`
}

// NewVSphereGetHostByMoRefOK creates VSphereGetHostByMoRefOK with default headers values
func NewVSphereGetHostByMoRefOK() *VSphereGetHostByMoRefOK {

	return &VSphereGetHostByMoRefOK{}
}

// WithPayload adds the payload to the v sphere get host by mo ref o k response
func (o *VSphereGetHostByMoRefOK) WithPayload(payload *models.ObjectCollection) *VSphereGetHostByMoRefOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get host by mo ref o k response
func (o *VSphereGetHostByMoRefOK) SetPayload(payload *models.ObjectCollection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetHostByMoRefOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VSphereGetHostByMoRefBadRequestCode is the HTTP code returned for type VSphereGetHostByMoRefBadRequest
const VSphereGetHostByMoRefBadRequestCode int = 400

/*
VSphereGetHostByMoRefBadRequest A general failure occured, more details are contained within the message.

swagger:response vSphereGetHostByMoRefBadRequest
*/
type VSphereGetHostByMoRefBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewVSphereGetHostByMoRefBadRequest creates VSphereGetHostByMoRefBadRequest with default headers values
func NewVSphereGetHostByMoRefBadRequest() *VSphereGetHostByMoRefBadRequest {

	return &VSphereGetHostByMoRefBadRequest{}
}

// WithPayload adds the payload to the v sphere get host by mo ref bad request response
func (o *VSphereGetHostByMoRefBadRequest) WithPayload(payload *models.BadRequestError) *VSphereGetHostByMoRefBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get host by mo ref bad request response
func (o *VSphereGetHostByMoRefBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetHostByMoRefBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
