// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListApplicationDefinitionsParams creates a new ListApplicationDefinitionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListApplicationDefinitionsParams() *ListApplicationDefinitionsParams {
	return &ListApplicationDefinitionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListApplicationDefinitionsParamsWithTimeout creates a new ListApplicationDefinitionsParams object
// with the ability to set a timeout on a request.
func NewListApplicationDefinitionsParamsWithTimeout(timeout time.Duration) *ListApplicationDefinitionsParams {
	return &ListApplicationDefinitionsParams{
		timeout: timeout,
	}
}

// NewListApplicationDefinitionsParamsWithContext creates a new ListApplicationDefinitionsParams object
// with the ability to set a context for a request.
func NewListApplicationDefinitionsParamsWithContext(ctx context.Context) *ListApplicationDefinitionsParams {
	return &ListApplicationDefinitionsParams{
		Context: ctx,
	}
}

// NewListApplicationDefinitionsParamsWithHTTPClient creates a new ListApplicationDefinitionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListApplicationDefinitionsParamsWithHTTPClient(client *http.Client) *ListApplicationDefinitionsParams {
	return &ListApplicationDefinitionsParams{
		HTTPClient: client,
	}
}

/*
ListApplicationDefinitionsParams contains all the parameters to send to the API endpoint

	for the list application definitions operation.

	Typically these are written to a http.Request.
*/
type ListApplicationDefinitionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list application definitions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListApplicationDefinitionsParams) WithDefaults() *ListApplicationDefinitionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list application definitions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListApplicationDefinitionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list application definitions params
func (o *ListApplicationDefinitionsParams) WithTimeout(timeout time.Duration) *ListApplicationDefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list application definitions params
func (o *ListApplicationDefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list application definitions params
func (o *ListApplicationDefinitionsParams) WithContext(ctx context.Context) *ListApplicationDefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list application definitions params
func (o *ListApplicationDefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list application definitions params
func (o *ListApplicationDefinitionsParams) WithHTTPClient(client *http.Client) *ListApplicationDefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list application definitions params
func (o *ListApplicationDefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListApplicationDefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}