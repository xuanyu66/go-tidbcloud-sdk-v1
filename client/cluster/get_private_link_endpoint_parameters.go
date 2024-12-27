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

// NewGetPrivateLinkEndpointParams creates a new GetPrivateLinkEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPrivateLinkEndpointParams() *GetPrivateLinkEndpointParams {
	return &GetPrivateLinkEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateLinkEndpointParamsWithTimeout creates a new GetPrivateLinkEndpointParams object
// with the ability to set a timeout on a request.
func NewGetPrivateLinkEndpointParamsWithTimeout(timeout time.Duration) *GetPrivateLinkEndpointParams {
	return &GetPrivateLinkEndpointParams{
		timeout: timeout,
	}
}

// NewGetPrivateLinkEndpointParamsWithContext creates a new GetPrivateLinkEndpointParams object
// with the ability to set a context for a request.
func NewGetPrivateLinkEndpointParamsWithContext(ctx context.Context) *GetPrivateLinkEndpointParams {
	return &GetPrivateLinkEndpointParams{
		Context: ctx,
	}
}

// NewGetPrivateLinkEndpointParamsWithHTTPClient creates a new GetPrivateLinkEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPrivateLinkEndpointParamsWithHTTPClient(client *http.Client) *GetPrivateLinkEndpointParams {
	return &GetPrivateLinkEndpointParams{
		HTTPClient: client,
	}
}

/*
GetPrivateLinkEndpointParams contains all the parameters to send to the API endpoint

	for the get private link endpoint operation.

	Typically these are written to a http.Request.
*/
type GetPrivateLinkEndpointParams struct {

	/* AccesspointID.

	   The ID of the accesspoint.

	   Format: uint64
	*/
	AccesspointID string

	/* ClusterID.

	   The ID of the cluster.

	   Format: uint64
	*/
	ClusterID string

	/* PrivateLinkEndpointID.

	   The ID of the private link endpoint.

	   Format: uint64
	*/
	PrivateLinkEndpointID string

	/* ProjectID.

	   The ID of the project.

	   Format: uint64
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get private link endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrivateLinkEndpointParams) WithDefaults() *GetPrivateLinkEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get private link endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrivateLinkEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) WithTimeout(timeout time.Duration) *GetPrivateLinkEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) WithContext(ctx context.Context) *GetPrivateLinkEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) WithHTTPClient(client *http.Client) *GetPrivateLinkEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccesspointID adds the accesspointID to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) WithAccesspointID(accesspointID string) *GetPrivateLinkEndpointParams {
	o.SetAccesspointID(accesspointID)
	return o
}

// SetAccesspointID adds the accesspointId to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) SetAccesspointID(accesspointID string) {
	o.AccesspointID = accesspointID
}

// WithClusterID adds the clusterID to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) WithClusterID(clusterID string) *GetPrivateLinkEndpointParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithPrivateLinkEndpointID adds the privateLinkEndpointID to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) WithPrivateLinkEndpointID(privateLinkEndpointID string) *GetPrivateLinkEndpointParams {
	o.SetPrivateLinkEndpointID(privateLinkEndpointID)
	return o
}

// SetPrivateLinkEndpointID adds the privateLinkEndpointId to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) SetPrivateLinkEndpointID(privateLinkEndpointID string) {
	o.PrivateLinkEndpointID = privateLinkEndpointID
}

// WithProjectID adds the projectID to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) WithProjectID(projectID string) *GetPrivateLinkEndpointParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get private link endpoint params
func (o *GetPrivateLinkEndpointParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateLinkEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accesspoint_id
	if err := r.SetPathParam("accesspoint_id", o.AccesspointID); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param private_link_endpoint_id
	if err := r.SetPathParam("private_link_endpoint_id", o.PrivateLinkEndpointID); err != nil {
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