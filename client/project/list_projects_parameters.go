// Code generated by go-swagger; DO NOT EDIT.

package project

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
	"github.com/go-openapi/swag"
)

// NewListProjectsParams creates a new ListProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectsParams() *ListProjectsParams {
	return &ListProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectsParamsWithTimeout creates a new ListProjectsParams object
// with the ability to set a timeout on a request.
func NewListProjectsParamsWithTimeout(timeout time.Duration) *ListProjectsParams {
	return &ListProjectsParams{
		timeout: timeout,
	}
}

// NewListProjectsParamsWithContext creates a new ListProjectsParams object
// with the ability to set a context for a request.
func NewListProjectsParamsWithContext(ctx context.Context) *ListProjectsParams {
	return &ListProjectsParams{
		Context: ctx,
	}
}

// NewListProjectsParamsWithHTTPClient creates a new ListProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectsParamsWithHTTPClient(client *http.Client) *ListProjectsParams {
	return &ListProjectsParams{
		HTTPClient: client,
	}
}

/*
ListProjectsParams contains all the parameters to send to the API endpoint

	for the list projects operation.

	Typically these are written to a http.Request.
*/
type ListProjectsParams struct {

	/* Page.

	   The number of pages.

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of a page.

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectsParams) WithDefaults() *ListProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListProjectsParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list projects params
func (o *ListProjectsParams) WithTimeout(timeout time.Duration) *ListProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list projects params
func (o *ListProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list projects params
func (o *ListProjectsParams) WithContext(ctx context.Context) *ListProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list projects params
func (o *ListProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list projects params
func (o *ListProjectsParams) WithHTTPClient(client *http.Client) *ListProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list projects params
func (o *ListProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the list projects params
func (o *ListProjectsParams) WithPage(page *int64) *ListProjectsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list projects params
func (o *ListProjectsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list projects params
func (o *ListProjectsParams) WithPageSize(pageSize *int64) *ListProjectsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list projects params
func (o *ListProjectsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}