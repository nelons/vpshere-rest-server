// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewVSphereGetHostByNameParams creates a new VSphereGetHostByNameParams object
//
// There are no default values defined in the spec.
func NewVSphereGetHostByNameParams() VSphereGetHostByNameParams {

	return VSphereGetHostByNameParams{}
}

// VSphereGetHostByNameParams contains all the bound params for the v sphere get host by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters vSphereGetHostByName
type VSphereGetHostByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	VRSToken string
	/*The name of the ESXi host
	  Required: true
	  In: path
	*/
	Hostname string
	/*A comma-delimited list of the properties to retrieve from the VM.
	  In: query
	*/
	Props *string
	/*The nicename for the connection.
	  Required: true
	  In: path
	*/
	Vcenter string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewVSphereGetHostByNameParams() beforehand.
func (o *VSphereGetHostByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindVRSToken(r.Header[http.CanonicalHeaderKey("VRS-Token")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rHostname, rhkHostname, _ := route.Params.GetOK("hostname")
	if err := o.bindHostname(rHostname, rhkHostname, route.Formats); err != nil {
		res = append(res, err)
	}

	qProps, qhkProps, _ := qs.GetOK("props")
	if err := o.bindProps(qProps, qhkProps, route.Formats); err != nil {
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
func (o *VSphereGetHostByNameParams) bindVRSToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindHostname binds and validates parameter Hostname from path.
func (o *VSphereGetHostByNameParams) bindHostname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Hostname = raw

	return nil
}

// bindProps binds and validates parameter Props from query.
func (o *VSphereGetHostByNameParams) bindProps(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Props = &raw

	return nil
}

// bindVcenter binds and validates parameter Vcenter from path.
func (o *VSphereGetHostByNameParams) bindVcenter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Vcenter = raw

	return nil
}
