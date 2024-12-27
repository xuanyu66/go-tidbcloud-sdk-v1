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

// NewDeletePrivateLinkEndpointParams creates a new DeletePrivateLinkEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePrivateLinkEndpointParams() *DeletePrivateLinkEndpointParams {
	return &DeletePrivateLinkEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrivateLinkEndpointParamsWithTimeout creates a new DeletePrivateLinkEndpointParams object
// with the ability to set a timeout on a request.
func NewDeletePrivateLinkEndpointParamsWithTimeout(timeout time.Duration) *DeletePrivateLinkEndpointParams {
	return &DeletePrivateLinkEndpointParams{
		timeout: timeout,
	}
}

// NewDeletePrivateLinkEndpointParamsWithContext creates a new DeletePrivateLinkEndpointParams object
// with the ability to set a context for a request.
func NewDeletePrivateLinkEndpointParamsWithContext(ctx context.Context) *DeletePrivateLinkEndpointParams {
	return &DeletePrivateLinkEndpointParams{
		Context: ctx,
	}
}

// NewDeletePrivateLinkEndpointParamsWithHTTPClient creates a new DeletePrivateLinkEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePrivateLinkEndpointParamsWithHTTPClient(client *http.Client) *DeletePrivateLinkEndpointParams {
	return &DeletePrivateLinkEndpointParams{
		HTTPClient: client,
	}
}

/*
DeletePrivateLinkEndpointParams contains all the parameters to send to the API endpoint

	for the delete private link endpoint operation.

	Typically these are written to a http.Request.
*/
type DeletePrivateLinkEndpointParams struct {

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

	   The ID of the private link endpoint to be deleted.

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

// WithDefaults hydrates default values in the delete private link endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateLinkEndpointParams) WithDefaults() *DeletePrivateLinkEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete private link endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateLinkEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) WithTimeout(timeout time.Duration) *DeletePrivateLinkEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) WithContext(ctx context.Context) *DeletePrivateLinkEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) WithHTTPClient(client *http.Client) *DeletePrivateLinkEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccesspointID adds the accesspointID to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) WithAccesspointID(accesspointID string) *DeletePrivateLinkEndpointParams {
	o.SetAccesspointID(accesspointID)
	return o
}

// SetAccesspointID adds the accesspointId to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) SetAccesspointID(accesspointID string) {
	o.AccesspointID = accesspointID
}

// WithClusterID adds the clusterID to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) WithClusterID(clusterID string) *DeletePrivateLinkEndpointParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithPrivateLinkEndpointID adds the privateLinkEndpointID to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) WithPrivateLinkEndpointID(privateLinkEndpointID string) *DeletePrivateLinkEndpointParams {
	o.SetPrivateLinkEndpointID(privateLinkEndpointID)
	return o
}

// SetPrivateLinkEndpointID adds the privateLinkEndpointId to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) SetPrivateLinkEndpointID(privateLinkEndpointID string) {
	o.PrivateLinkEndpointID = privateLinkEndpointID
}

// WithProjectID adds the projectID to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) WithProjectID(projectID string) *DeletePrivateLinkEndpointParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete private link endpoint params
func (o *DeletePrivateLinkEndpointParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateLinkEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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