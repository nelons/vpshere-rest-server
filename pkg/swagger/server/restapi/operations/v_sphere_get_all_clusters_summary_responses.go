// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nelons/vsphere-rest-server/pkg/swagger/server/models"
)

// VSphereGetAllClustersSummaryOKCode is the HTTP code returned for type VSphereGetAllClustersSummaryOK
const VSphereGetAllClustersSummaryOKCode int = 200

/*
VSphereGetAllClustersSummaryOK Successful request. Returns JSON containing a count and the collection of objects.

swagger:response vSphereGetAllClustersSummaryOK
*/
type VSphereGetAllClustersSummaryOK struct {

	/*
	  In: Body
	*/
	Payload *models.ObjectCollection `json:"body,omitempty"`
}

// NewVSphereGetAllClustersSummaryOK creates VSphereGetAllClustersSummaryOK with default headers values
func NewVSphereGetAllClustersSummaryOK() *VSphereGetAllClustersSummaryOK {

	return &VSphereGetAllClustersSummaryOK{}
}

// WithPayload adds the payload to the v sphere get all clusters summary o k response
func (o *VSphereGetAllClustersSummaryOK) WithPayload(payload *models.ObjectCollection) *VSphereGetAllClustersSummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get all clusters summary o k response
func (o *VSphereGetAllClustersSummaryOK) SetPayload(payload *models.ObjectCollection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetAllClustersSummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VSphereGetAllClustersSummaryBadRequestCode is the HTTP code returned for type VSphereGetAllClustersSummaryBadRequest
const VSphereGetAllClustersSummaryBadRequestCode int = 400

/*
VSphereGetAllClustersSummaryBadRequest A general failure occured, more details are contained within the message.

swagger:response vSphereGetAllClustersSummaryBadRequest
*/
type VSphereGetAllClustersSummaryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewVSphereGetAllClustersSummaryBadRequest creates VSphereGetAllClustersSummaryBadRequest with default headers values
func NewVSphereGetAllClustersSummaryBadRequest() *VSphereGetAllClustersSummaryBadRequest {

	return &VSphereGetAllClustersSummaryBadRequest{}
}

// WithPayload adds the payload to the v sphere get all clusters summary bad request response
func (o *VSphereGetAllClustersSummaryBadRequest) WithPayload(payload *models.BadRequestError) *VSphereGetAllClustersSummaryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere get all clusters summary bad request response
func (o *VSphereGetAllClustersSummaryBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereGetAllClustersSummaryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
