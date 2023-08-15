// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VSphereGetAllDatacentersHandlerFunc turns a function with the right signature into a v sphere get all datacenters handler
type VSphereGetAllDatacentersHandlerFunc func(VSphereGetAllDatacentersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VSphereGetAllDatacentersHandlerFunc) Handle(params VSphereGetAllDatacentersParams) middleware.Responder {
	return fn(params)
}

// VSphereGetAllDatacentersHandler interface for that can handle valid v sphere get all datacenters params
type VSphereGetAllDatacentersHandler interface {
	Handle(VSphereGetAllDatacentersParams) middleware.Responder
}

// NewVSphereGetAllDatacenters creates a new http.Handler for the v sphere get all datacenters operation
func NewVSphereGetAllDatacenters(ctx *middleware.Context, handler VSphereGetAllDatacentersHandler) *VSphereGetAllDatacenters {
	return &VSphereGetAllDatacenters{Context: ctx, Handler: handler}
}

/*
	VSphereGetAllDatacenters swagger:route GET /vsphere/{vcenter}/datacenter vSphereGetAllDatacenters

Gets the datacenters.
*/
type VSphereGetAllDatacenters struct {
	Context *middleware.Context
	Handler VSphereGetAllDatacentersHandler
}

func (o *VSphereGetAllDatacenters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVSphereGetAllDatacentersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
