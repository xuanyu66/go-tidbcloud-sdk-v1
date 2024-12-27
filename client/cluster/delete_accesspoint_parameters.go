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

// NewDeleteAccesspointParams creates a new DeleteAccesspointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAccesspointParams() *DeleteAccesspointParams {
	return &DeleteAccesspointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccesspointParamsWithTimeout creates a new DeleteAccesspointParams object
// with the ability to set a timeout on a request.
func NewDeleteAccesspointParamsWithTimeout(timeout time.Duration) *DeleteAccesspointParams {
	return &DeleteAccesspointParams{
		timeout: timeout,
	}
}

// NewDeleteAccesspointParamsWithContext creates a new DeleteAccesspointParams object
// with the ability to set a context for a request.
func NewDeleteAccesspointParamsWithContext(ctx context.Context) *DeleteAccesspointParams {
	return &DeleteAccesspointParams{
		Context: ctx,
	}
}

// NewDeleteAccesspointParamsWithHTTPClient creates a new DeleteAccesspointParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAccesspointParamsWithHTTPClient(client *http.Client) *DeleteAccesspointParams {
	return &DeleteAccesspointParams{
		HTTPClient: client,
	}
}

/*
DeleteAccesspointParams contains all the parameters to send to the API endpoint

	for the delete accesspoint operation.

	Typically these are written to a http.Request.
*/
type DeleteAccesspointParams struct {

	/* AccesspointID.

	   The ID of the accesspoint to be deleted.

	   Format: uint64
	*/
	AccesspointID string

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

// WithDefaults hydrates default values in the delete accesspoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccesspointParams) WithDefaults() *DeleteAccesspointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete accesspoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccesspointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete accesspoint params
func (o *DeleteAccesspointParams) WithTimeout(timeout time.Duration) *DeleteAccesspointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete accesspoint params
func (o *DeleteAccesspointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete accesspoint params
func (o *DeleteAccesspointParams) WithContext(ctx context.Context) *DeleteAccesspointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete accesspoint params
func (o *DeleteAccesspointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete accesspoint params
func (o *DeleteAccesspointParams) WithHTTPClient(client *http.Client) *DeleteAccesspointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete accesspoint params
func (o *DeleteAccesspointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccesspointID adds the accesspointID to the delete accesspoint params
func (o *DeleteAccesspointParams) WithAccesspointID(accesspointID string) *DeleteAccesspointParams {
	o.SetAccesspointID(accesspointID)
	return o
}

// SetAccesspointID adds the accesspointId to the delete accesspoint params
func (o *DeleteAccesspointParams) SetAccesspointID(accesspointID string) {
	o.AccesspointID = accesspointID
}

// WithClusterID adds the clusterID to the delete accesspoint params
func (o *DeleteAccesspointParams) WithClusterID(clusterID string) *DeleteAccesspointParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete accesspoint params
func (o *DeleteAccesspointParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the delete accesspoint params
func (o *DeleteAccesspointParams) WithProjectID(projectID string) *DeleteAccesspointParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete accesspoint params
func (o *DeleteAccesspointParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccesspointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}