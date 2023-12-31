// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VSphereListConnectionsHandlerFunc turns a function with the right signature into a v sphere list connections handler
type VSphereListConnectionsHandlerFunc func(VSphereListConnectionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VSphereListConnectionsHandlerFunc) Handle(params VSphereListConnectionsParams) middleware.Responder {
	return fn(params)
}

// VSphereListConnectionsHandler interface for that can handle valid v sphere list connections params
type VSphereListConnectionsHandler interface {
	Handle(VSphereListConnectionsParams) middleware.Responder
}

// NewVSphereListConnections creates a new http.Handler for the v sphere list connections operation
func NewVSphereListConnections(ctx *middleware.Context, handler VSphereListConnectionsHandler) *VSphereListConnections {
	return &VSphereListConnections{Context: ctx, Handler: handler}
}

/*
	VSphereListConnections swagger:route GET /vsphere/list vSphereListConnections

Provides a list of all vSphere Connections
*/
type VSphereListConnections struct {
	Context *middleware.Context
	Handler VSphereListConnectionsHandler
}

func (o *VSphereListConnections) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVSphereListConnectionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// VSphereListConnectionsOKBodyItems0 v sphere list connections o k body items0
//
// swagger:model VSphereListConnectionsOKBodyItems0
type VSphereListConnectionsOKBodyItems0 struct {

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this v sphere list connections o k body items0
func (o *VSphereListConnectionsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v sphere list connections o k body items0 based on context it is used
func (o *VSphereListConnectionsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VSphereListConnectionsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VSphereListConnectionsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res VSphereListConnectionsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
