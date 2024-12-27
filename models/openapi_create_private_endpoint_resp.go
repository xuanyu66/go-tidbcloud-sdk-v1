// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiCreatePrivateEndpointResp CreatePrivateEndpointResp
//
// CreatePrivateEndpointResp is the response for creating a private endpoint for a private endpoint service.
//
// swagger:model openapiCreatePrivateEndpointResp
type OpenapiCreatePrivateEndpointResp struct {

	// private endpoint
	// Required: true
	PrivateEndpoint *OpenapiCreatePrivateEndpointRespPrivateEndpoint `json:"private_endpoint"`
}

// Validate validates this openapi create private endpoint resp
func (m *OpenapiCreatePrivateEndpointResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointResp) validatePrivateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("private_endpoint", "body", m.PrivateEndpoint); err != nil {
		return err
	}

	if m.PrivateEndpoint != nil {
		if err := m.PrivateEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("private_endpoint")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi create private endpoint resp based on the context it is used
func (m *OpenapiCreatePrivateEndpointResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointResp) contextValidatePrivateEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivateEndpoint != nil {

		if err := m.PrivateEndpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("private_endpoint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointResp) UnmarshalBinary(b []byte) error {
	var res OpenapiCreatePrivateEndpointResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiCreatePrivateEndpointRespPrivateEndpoint The newly endpoint created resource.
//
// swagger:model OpenapiCreatePrivateEndpointRespPrivateEndpoint
type OpenapiCreatePrivateEndpointRespPrivateEndpoint struct {

	// [Output Only] The cloud provider on which the private endpoint service is hosted.
	// - `"AWS"`: the Amazon Web Services cloud provider
	// - `"GCP"`: the Google Cloud cloud provider
	// Example: AWS
	// Enum: ["AWS","GCP"]
	CloudProvider string `json:"cloud_provider,omitempty"`

	// [Output Only] The ID of the cluster.
	// Example: 1
	ClusterID string `json:"cluster_id,omitempty"`

	// [Output Only] The name of the cluster.
	// Example: Cluster0
	ClusterName string `json:"cluster_name,omitempty"`

	// The format of the private endpoint name varies by cloud provider: `"vpce-xxxx"` for AWS and `"projects/xxx/regions/xxx/forwardingRules/xxx"` for Google Cloud.
	// Example: vpce-01234567890123456
	EndpointName string `json:"endpoint_name,omitempty"`

	// [Output Only] The ID of private endpoint. It is used when you [deleting the endpoint](#tag/Cluster/operation/DeletePrivateEndpoint).
	// Example: 1
	ID string `json:"id,omitempty"`

	// [Output Only] The detailed message when the `status` is "FAILED".
	// Example: The endpoint does not exist.
	Message string `json:"message,omitempty"`

	// [Output Only] The region where the private endpoint is hosted, such as Oregon in AWS.
	// Example: Oregon
	RegionName string `json:"region_name,omitempty"`

	// [Output Only] The service name of the private endpoint.
	// Example: com.amazonaws.vpce.us-west-2.vpce-svc-01234567890123456
	ServiceName string `json:"service_name,omitempty"`

	// [Output Only] The status of the private endpoint service.
	// Example: ACTIVE
	// Enum: ["CREATING","ACTIVE","DELETING"]
	ServiceStatus string `json:"service_status,omitempty"`

	// [Output Only] The status of the private endpoint.
	// Example: FAILED
	// Enum: ["PENDING","ACTIVE","DELETING","FAILED"]
	Status string `json:"status,omitempty"`
}

// Validate validates this openapi create private endpoint resp private endpoint
func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openapiCreatePrivateEndpointRespPrivateEndpointTypeCloudProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiCreatePrivateEndpointRespPrivateEndpointTypeCloudProviderPropEnum = append(openapiCreatePrivateEndpointRespPrivateEndpointTypeCloudProviderPropEnum, v)
	}
}

const (

	// OpenapiCreatePrivateEndpointRespPrivateEndpointCloudProviderAWS captures enum value "AWS"
	OpenapiCreatePrivateEndpointRespPrivateEndpointCloudProviderAWS string = "AWS"

	// OpenapiCreatePrivateEndpointRespPrivateEndpointCloudProviderGCP captures enum value "GCP"
	OpenapiCreatePrivateEndpointRespPrivateEndpointCloudProviderGCP string = "GCP"
)

// prop value enum
func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) validateCloudProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiCreatePrivateEndpointRespPrivateEndpointTypeCloudProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) validateCloudProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudProviderEnum("private_endpoint"+"."+"cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

var openapiCreatePrivateEndpointRespPrivateEndpointTypeServiceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiCreatePrivateEndpointRespPrivateEndpointTypeServiceStatusPropEnum = append(openapiCreatePrivateEndpointRespPrivateEndpointTypeServiceStatusPropEnum, v)
	}
}

const (

	// OpenapiCreatePrivateEndpointRespPrivateEndpointServiceStatusCREATING captures enum value "CREATING"
	OpenapiCreatePrivateEndpointRespPrivateEndpointServiceStatusCREATING string = "CREATING"

	// OpenapiCreatePrivateEndpointRespPrivateEndpointServiceStatusACTIVE captures enum value "ACTIVE"
	OpenapiCreatePrivateEndpointRespPrivateEndpointServiceStatusACTIVE string = "ACTIVE"

	// OpenapiCreatePrivateEndpointRespPrivateEndpointServiceStatusDELETING captures enum value "DELETING"
	OpenapiCreatePrivateEndpointRespPrivateEndpointServiceStatusDELETING string = "DELETING"
)

// prop value enum
func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) validateServiceStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiCreatePrivateEndpointRespPrivateEndpointTypeServiceStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) validateServiceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceStatusEnum("private_endpoint"+"."+"service_status", "body", m.ServiceStatus); err != nil {
		return err
	}

	return nil
}

var openapiCreatePrivateEndpointRespPrivateEndpointTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","ACTIVE","DELETING","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiCreatePrivateEndpointRespPrivateEndpointTypeStatusPropEnum = append(openapiCreatePrivateEndpointRespPrivateEndpointTypeStatusPropEnum, v)
	}
}

const (

	// OpenapiCreatePrivateEndpointRespPrivateEndpointStatusPENDING captures enum value "PENDING"
	OpenapiCreatePrivateEndpointRespPrivateEndpointStatusPENDING string = "PENDING"

	// OpenapiCreatePrivateEndpointRespPrivateEndpointStatusACTIVE captures enum value "ACTIVE"
	OpenapiCreatePrivateEndpointRespPrivateEndpointStatusACTIVE string = "ACTIVE"

	// OpenapiCreatePrivateEndpointRespPrivateEndpointStatusDELETING captures enum value "DELETING"
	OpenapiCreatePrivateEndpointRespPrivateEndpointStatusDELETING string = "DELETING"

	// OpenapiCreatePrivateEndpointRespPrivateEndpointStatusFAILED captures enum value "FAILED"
	OpenapiCreatePrivateEndpointRespPrivateEndpointStatusFAILED string = "FAILED"
)

// prop value enum
func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiCreatePrivateEndpointRespPrivateEndpointTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("private_endpoint"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi create private endpoint resp private endpoint based on context it is used
func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiCreatePrivateEndpointRespPrivateEndpoint) UnmarshalBinary(b []byte) error {
	var res OpenapiCreatePrivateEndpointRespPrivateEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
