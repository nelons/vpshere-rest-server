// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VSphereGetAllStoragePodsHandlerFunc turns a function with the right signature into a v sphere get all storage pods handler
type VSphereGetAllStoragePodsHandlerFunc func(VSphereGetAllStoragePodsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VSphereGetAllStoragePodsHandlerFunc) Handle(params VSphereGetAllStoragePodsParams) middleware.Responder {
	return fn(params)
}

// VSphereGetAllStoragePodsHandler interface for that can handle valid v sphere get all storage pods params
type VSphereGetAllStoragePodsHandler interface {
	Handle(VSphereGetAllStoragePodsParams) middleware.Responder
}

// NewVSphereGetAllStoragePods creates a new http.Handler for the v sphere get all storage pods operation
func NewVSphereGetAllStoragePods(ctx *middleware.Context, handler VSphereGetAllStoragePodsHandler) *VSphereGetAllStoragePods {
	return &VSphereGetAllStoragePods{Context: ctx, Handler: handler}
}

/*
	VSphereGetAllStoragePods swagger:route GET /vsphere/{vcenter}/storagepod vSphereGetAllStoragePods

Gets the hosts found at the endpoint.
*/
type VSphereGetAllStoragePods struct {
	Context *middleware.Context
	Handler VSphereGetAllStoragePodsHandler
}

func (o *VSphereGetAllStoragePods) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVSphereGetAllStoragePodsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
