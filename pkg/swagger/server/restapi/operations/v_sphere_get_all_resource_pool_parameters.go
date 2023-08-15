// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewVSphereGetAllResourcePoolParams creates a new VSphereGetAllResourcePoolParams object
//
// There are no default values defined in the spec.
func NewVSphereGetAllResourcePoolParams() VSphereGetAllResourcePoolParams {

	return VSphereGetAllResourcePoolParams{}
}

// VSphereGetAllResourcePoolParams contains all the bound params for the v sphere get all resource pool operation
// typically these are obtained from a http.Request
//
// swagger:parameters vSphereGetAllResourcePool
type VSphereGetAllResourcePoolParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	VRSToken string
	/*The nicename for the connection.
	  Required: true
	  In: path
	*/
	Vcenter string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewVSphereGetAllResourcePoolParams() beforehand.
func (o *VSphereGetAllResourcePoolParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindVRSToken(r.Header[http.CanonicalHeaderKey("VRS-Token")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rVcenter, rhkVcenter, _ := route.Params.GetOK("vcenter")
	if err := o.bindVcenter(rVcenter, rhkVcenter, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVRSToken binds and validates parameter VRSToken from header.
func (o *VSphereGetAllResourcePoolParams) bindVRSToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("VRS-Token", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("VRS-Token", "header", raw); err != nil {
		return err
	}
	o.VRSToken = raw

	return nil
}

// bindVcenter binds and validates parameter Vcenter from path.
func (o *VSphereGetAllResourcePoolParams) bindVcenter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Vcenter = raw

	return nil
}
