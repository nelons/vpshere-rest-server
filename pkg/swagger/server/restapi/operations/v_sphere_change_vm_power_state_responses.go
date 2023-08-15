// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nelons/vsphere-rest-server/pkg/swagger/server/models"
)

// VSphereChangeVMPowerStateOKCode is the HTTP code returned for type VSphereChangeVMPowerStateOK
const VSphereChangeVMPowerStateOKCode int = 200

/*
VSphereChangeVMPowerStateOK The power state change task has been initiated.

swagger:response vSphereChangeVmPowerStateOK
*/
type VSphereChangeVMPowerStateOK struct {
}

// NewVSphereChangeVMPowerStateOK creates VSphereChangeVMPowerStateOK with default headers values
func NewVSphereChangeVMPowerStateOK() *VSphereChangeVMPowerStateOK {

	return &VSphereChangeVMPowerStateOK{}
}

// WriteResponse to the client
func (o *VSphereChangeVMPowerStateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// VSphereChangeVMPowerStateBadRequestCode is the HTTP code returned for type VSphereChangeVMPowerStateBadRequest
const VSphereChangeVMPowerStateBadRequestCode int = 400

/*
VSphereChangeVMPowerStateBadRequest A general failure occured, more details are contained within the message.

swagger:response vSphereChangeVmPowerStateBadRequest
*/
type VSphereChangeVMPowerStateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewVSphereChangeVMPowerStateBadRequest creates VSphereChangeVMPowerStateBadRequest with default headers values
func NewVSphereChangeVMPowerStateBadRequest() *VSphereChangeVMPowerStateBadRequest {

	return &VSphereChangeVMPowerStateBadRequest{}
}

// WithPayload adds the payload to the v sphere change Vm power state bad request response
func (o *VSphereChangeVMPowerStateBadRequest) WithPayload(payload *models.BadRequestError) *VSphereChangeVMPowerStateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v sphere change Vm power state bad request response
func (o *VSphereChangeVMPowerStateBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VSphereChangeVMPowerStateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}