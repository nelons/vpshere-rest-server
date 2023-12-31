// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SessionListHandlerFunc turns a function with the right signature into a session list handler
type SessionListHandlerFunc func(SessionListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SessionListHandlerFunc) Handle(params SessionListParams) middleware.Responder {
	return fn(params)
}

// SessionListHandler interface for that can handle valid session list params
type SessionListHandler interface {
	Handle(SessionListParams) middleware.Responder
}

// NewSessionList creates a new http.Handler for the session list operation
func NewSessionList(ctx *middleware.Context, handler SessionListHandler) *SessionList {
	return &SessionList{Context: ctx, Handler: handler}
}

/*
	SessionList swagger:route GET /session/list sessionList

SessionList session list API
*/
type SessionList struct {
	Context *middleware.Context
	Handler SessionListHandler
}

func (o *SessionList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSessionListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SessionListOKBodyItems0 session list o k body items0
//
// swagger:model SessionListOKBodyItems0
type SessionListOKBodyItems0 struct {

	// connections
	Connections []*SessionListOKBodyItems0ConnectionsItems0 `json:"Connections"`

	// host
	Host string `json:"Host,omitempty"`

	// last access
	LastAccess string `json:"LastAccess,omitempty"`

	// secret
	Secret string `json:"Secret,omitempty"`
}

// Validate validates this session list o k body items0
func (o *SessionListOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SessionListOKBodyItems0) validateConnections(formats strfmt.Registry) error {
	if swag.IsZero(o.Connections) { // not required
		return nil
	}

	for i := 0; i < len(o.Connections); i++ {
		if swag.IsZero(o.Connections[i]) { // not required
			continue
		}

		if o.Connections[i] != nil {
			if err := o.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this session list o k body items0 based on the context it is used
func (o *SessionListOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SessionListOKBodyItems0) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Connections); i++ {

		if o.Connections[i] != nil {

			if swag.IsZero(o.Connections[i]) { // not required
				return nil
			}

			if err := o.Connections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SessionListOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SessionListOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res SessionListOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// SessionListOKBodyItems0ConnectionsItems0 session list o k body items0 connections items0
//
// swagger:model SessionListOKBodyItems0ConnectionsItems0
type SessionListOKBodyItems0ConnectionsItems0 struct {

	// url
	URL string `json:"url,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this session list o k body items0 connections items0
func (o *SessionListOKBodyItems0ConnectionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this session list o k body items0 connections items0 based on context it is used
func (o *SessionListOKBodyItems0ConnectionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SessionListOKBodyItems0ConnectionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SessionListOKBodyItems0ConnectionsItems0) UnmarshalBinary(b []byte) error {
	var res SessionListOKBodyItems0ConnectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
