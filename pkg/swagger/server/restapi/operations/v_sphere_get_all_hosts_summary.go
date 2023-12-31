// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VSphereGetAllHostsSummaryHandlerFunc turns a function with the right signature into a v sphere get all hosts summary handler
type VSphereGetAllHostsSummaryHandlerFunc func(VSphereGetAllHostsSummaryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VSphereGetAllHostsSummaryHandlerFunc) Handle(params VSphereGetAllHostsSummaryParams) middleware.Responder {
	return fn(params)
}

// VSphereGetAllHostsSummaryHandler interface for that can handle valid v sphere get all hosts summary params
type VSphereGetAllHostsSummaryHandler interface {
	Handle(VSphereGetAllHostsSummaryParams) middleware.Responder
}

// NewVSphereGetAllHostsSummary creates a new http.Handler for the v sphere get all hosts summary operation
func NewVSphereGetAllHostsSummary(ctx *middleware.Context, handler VSphereGetAllHostsSummaryHandler) *VSphereGetAllHostsSummary {
	return &VSphereGetAllHostsSummary{Context: ctx, Handler: handler}
}

/*
	VSphereGetAllHostsSummary swagger:route GET /vsphere/{vcenter}/host vSphereGetAllHostsSummary

Gets the hosts found at the endpoint.
*/
type VSphereGetAllHostsSummary struct {
	Context *middleware.Context
	Handler VSphereGetAllHostsSummaryHandler
}

func (o *VSphereGetAllHostsSummary) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVSphereGetAllHostsSummaryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
