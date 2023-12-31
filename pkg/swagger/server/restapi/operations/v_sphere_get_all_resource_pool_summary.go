// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VSphereGetAllResourcePoolSummaryHandlerFunc turns a function with the right signature into a v sphere get all resource pool summary handler
type VSphereGetAllResourcePoolSummaryHandlerFunc func(VSphereGetAllResourcePoolSummaryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VSphereGetAllResourcePoolSummaryHandlerFunc) Handle(params VSphereGetAllResourcePoolSummaryParams) middleware.Responder {
	return fn(params)
}

// VSphereGetAllResourcePoolSummaryHandler interface for that can handle valid v sphere get all resource pool summary params
type VSphereGetAllResourcePoolSummaryHandler interface {
	Handle(VSphereGetAllResourcePoolSummaryParams) middleware.Responder
}

// NewVSphereGetAllResourcePoolSummary creates a new http.Handler for the v sphere get all resource pool summary operation
func NewVSphereGetAllResourcePoolSummary(ctx *middleware.Context, handler VSphereGetAllResourcePoolSummaryHandler) *VSphereGetAllResourcePoolSummary {
	return &VSphereGetAllResourcePoolSummary{Context: ctx, Handler: handler}
}

/*
	VSphereGetAllResourcePoolSummary swagger:route GET /vsphere/{vcenter}/resourcepool vSphereGetAllResourcePoolSummary

Gets the hosts found at the endpoint.
*/
type VSphereGetAllResourcePoolSummary struct {
	Context *middleware.Context
	Handler VSphereGetAllResourcePoolSummaryHandler
}

func (o *VSphereGetAllResourcePoolSummary) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVSphereGetAllResourcePoolSummaryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
