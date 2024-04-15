// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewListPrivateEndpointsParams creates a new ListPrivateEndpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPrivateEndpointsParams() *ListPrivateEndpointsParams {
	return &ListPrivateEndpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPrivateEndpointsParamsWithTimeout creates a new ListPrivateEndpointsParams object
// with the ability to set a timeout on a request.
func NewListPrivateEndpointsParamsWithTimeout(timeout time.Duration) *ListPrivateEndpointsParams {
	return &ListPrivateEndpointsParams{
		timeout: timeout,
	}
}

// NewListPrivateEndpointsParamsWithContext creates a new ListPrivateEndpointsParams object
// with the ability to set a context for a request.
func NewListPrivateEndpointsParamsWithContext(ctx context.Context) *ListPrivateEndpointsParams {
	return &ListPrivateEndpointsParams{
		Context: ctx,
	}
}

// NewListPrivateEndpointsParamsWithHTTPClient creates a new ListPrivateEndpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPrivateEndpointsParamsWithHTTPClient(client *http.Client) *ListPrivateEndpointsParams {
	return &ListPrivateEndpointsParams{
		HTTPClient: client,
	}
}

/*
ListPrivateEndpointsParams contains all the parameters to send to the API endpoint

	for the list private endpoints operation.

	Typically these are written to a http.Request.
*/
type ListPrivateEndpointsParams struct {

	/* ClusterID.

	   The ID of the cluster.

	   Format: uint64
	*/
	ClusterID string

	/* ProjectID.

	   The ID of the project. You can get the project ID from the response of [List all accessible projects](#tag/Project/operation/ListProjects).

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list private endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPrivateEndpointsParams) WithDefaults() *ListPrivateEndpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list private endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPrivateEndpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list private endpoints params
func (o *ListPrivateEndpointsParams) WithTimeout(timeout time.Duration) *ListPrivateEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list private endpoints params
func (o *ListPrivateEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list private endpoints params
func (o *ListPrivateEndpointsParams) WithContext(ctx context.Context) *ListPrivateEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list private endpoints params
func (o *ListPrivateEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list private endpoints params
func (o *ListPrivateEndpointsParams) WithHTTPClient(client *http.Client) *ListPrivateEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list private endpoints params
func (o *ListPrivateEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list private endpoints params
func (o *ListPrivateEndpointsParams) WithClusterID(clusterID string) *ListPrivateEndpointsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list private endpoints params
func (o *ListPrivateEndpointsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list private endpoints params
func (o *ListPrivateEndpointsParams) WithProjectID(projectID string) *ListPrivateEndpointsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list private endpoints params
func (o *ListPrivateEndpointsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPrivateEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
