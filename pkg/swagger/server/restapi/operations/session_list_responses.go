// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SessionListOKCode is the HTTP code returned for type SessionListOK
const SessionListOKCode int = 200

/*
SessionListOK Lists the sessions in use on the server

swagger:response sessionListOK
*/
type SessionListOK struct {

	/*
	  In: Body
	*/
	Payload []*SessionListOKBodyItems0 `json:"body,omitempty"`
}

// NewSessionListOK creates SessionListOK with default headers values
func NewSessionListOK() *SessionListOK {

	return &SessionListOK{}
}

// WithPayload adds the payload to the session list o k response
func (o *SessionListOK) WithPayload(payload []*SessionListOKBodyItems0) *SessionListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the session list o k response
func (o *SessionListOK) SetPayload(payload []*SessionListOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SessionListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*SessionListOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
