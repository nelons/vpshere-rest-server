// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// VSphereGetHostByMoRefURL generates an URL for the v sphere get host by mo ref operation
type VSphereGetHostByMoRefURL struct {
	Moref   string
	Vcenter string

	Props *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *VSphereGetHostByMoRefURL) WithBasePath(bp string) *VSphereGetHostByMoRefURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *VSphereGetHostByMoRefURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *VSphereGetHostByMoRefURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/vsphere/{vcenter}/host/ref/{moref}"

	moref := o.Moref
	if moref != "" {
		_path = strings.Replace(_path, "{moref}", moref, -1)
	} else {
		return nil, errors.New("moref is required on VSphereGetHostByMoRefURL")
	}

	vcenter := o.Vcenter
	if vcenter != "" {
		_path = strings.Replace(_path, "{vcenter}", vcenter, -1)
	} else {
		return nil, errors.New("vcenter is required on VSphereGetHostByMoRefURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var propsQ string
	if o.Props != nil {
		propsQ = *o.Props
	}
	if propsQ != "" {
		qs.Set("props", propsQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *VSphereGetHostByMoRefURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *VSphereGetHostByMoRefURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *VSphereGetHostByMoRefURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on VSphereGetHostByMoRefURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on VSphereGetHostByMoRefURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *VSphereGetHostByMoRefURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}