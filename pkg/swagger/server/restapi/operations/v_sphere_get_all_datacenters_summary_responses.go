// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nelons/vsphere-rest-server/pkg/swagger/server/models"
)

// VSphereGetAllDatacentersSummaryOKCode is the HTTP code returned for type VSphereGetAllDatacentersSummaryOK
const VSphereGetAllDatacentersSummaryOKCode int = 200

/*
VSphereGetAllDatacentersSummaryOK Successful request. Returns JSON containing a count and the collection of objects.

swagger:response vSphereGetAllDatacentersSummaryOK
*/
type VSphereGetAllDatacentersSummaryOK struct {

	/*
	  In: Body
	*/
	Payload *models.ObjectCollection `json:"body,omitempty"`
}

// NewVSphereGetAllDatacentersSummaryOK creates VSphereGetAllDatacentersSummaryOK with default headers values
func NewVSphereGetAllDatacentersSummaryOK() *VSphereGetAllDatacentersSummaryOK {

	return &VSphereGetAllDatacentersSummaryOK{}
}

// WithPayload adds the payload to the v sphere get all datacenters summary o k response
func (o *VSphereGetAllDatacentersSummaryOK) WithPayload(payload *models.ObjectCollection) *VSphereGetAllDatacentersSummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get all datacenters summary o k response
func (o *VSphereGetAllDatacentersSummaryOK) SetPayload(payload *models.ObjectCollection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetAllDatacentersSummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VSphereGetAllDatacentersSummaryBadRequestCode is the HTTP code returned for type VSphereGetAllDatacentersSummaryBadRequest
const VSphereGetAllDatacentersSummaryBadRequestCode int = 400

/*
VSphereGetAllDatacentersSummaryBadRequest A general failure occured, more details are contained within the message.

swagger:response vSphereGetAllDatacentersSummaryBadRequest
*/
type VSphereGetAllDatacentersSummaryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewVSphereGetAllDatacentersSummaryBadRequest creates VSphereGetAllDatacentersSummaryBadRequest with default headers values
func NewVSphereGetAllDatacentersSummaryBadRequest() *VSphereGetAllDatacentersSummaryBadRequest {

	return &VSphereGetAllDatacentersSummaryBadRequest{}
}

// WithPayload adds the payload to the v sphere get all datacenters summary bad request response
func (o *VSphereGetAllDatacentersSummaryBadRequest) WithPayload(payload *models.BadRequestError) *VSphereGetAllDatacentersSummaryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get all datacenters summary bad request response
func (o *VSphereGetAllDatacentersSummaryBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetAllDatacentersSummaryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
