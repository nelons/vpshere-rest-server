// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nelons/vsphere-rest-server/pkg/swagger/server/models"
)

// VSphereGetAllResourcePoolSummaryOKCode is the HTTP code returned for type VSphereGetAllResourcePoolSummaryOK
const VSphereGetAllResourcePoolSummaryOKCode int = 200

/*
VSphereGetAllResourcePoolSummaryOK Successful request. Returns JSON containing a count and the collection of objects.

swagger:response vSphereGetAllResourcePoolSummaryOK
*/
type VSphereGetAllResourcePoolSummaryOK struct {

	/*
	  In: Body
	*/
	Payload *models.ObjectCollection `json:"body,omitempty"`
}

// NewVSphereGetAllResourcePoolSummaryOK creates VSphereGetAllResourcePoolSummaryOK with default headers values
func NewVSphereGetAllResourcePoolSummaryOK() *VSphereGetAllResourcePoolSummaryOK {

	return &VSphereGetAllResourcePoolSummaryOK{}
}

// WithPayload adds the payload to the v sphere get all resource pool summary o k response
func (o *VSphereGetAllResourcePoolSummaryOK) WithPayload(payload *models.ObjectCollection) *VSphereGetAllResourcePoolSummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get all resource pool summary o k response
func (o *VSphereGetAllResourcePoolSummaryOK) SetPayload(payload *models.ObjectCollection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetAllResourcePoolSummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VSphereGetAllResourcePoolSummaryBadRequestCode is the HTTP code returned for type VSphereGetAllResourcePoolSummaryBadRequest
const VSphereGetAllResourcePoolSummaryBadRequestCode int = 400

/*
VSphereGetAllResourcePoolSummaryBadRequest A general failure occured, more details are contained within the message.

swagger:response vSphereGetAllResourcePoolSummaryBadRequest
*/
type VSphereGetAllResourcePoolSummaryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewVSphereGetAllResourcePoolSummaryBadRequest creates VSphereGetAllResourcePoolSummaryBadRequest with default headers values
func NewVSphereGetAllResourcePoolSummaryBadRequest() *VSphereGetAllResourcePoolSummaryBadRequest {

	return &VSphereGetAllResourcePoolSummaryBadRequest{}
}

// WithPayload adds the payload to the v sphere get all resource pool summary bad request response
func (o *VSphereGetAllResourcePoolSummaryBadRequest) WithPayload(payload *models.BadRequestError) *VSphereGetAllResourcePoolSummaryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get all resource pool summary bad request response
func (o *VSphereGetAllResourcePoolSummaryBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetAllResourcePoolSummaryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
