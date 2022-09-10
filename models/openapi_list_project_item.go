// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiListProjectItem ListProjectItem
//
// ListProjectItem is the item of projects.
//
// swagger:model openapiListProjectItem
type OpenapiListProjectItem struct {

	// The number of TiDB Cloud clusters deployed in the project.
	// Example: 4
	ClusterCount int64 `json:"cluster_count,omitempty"`

	// The creation time of the cluster in Unix timestamp seconds (epoch time).
	// Example: 1656991448
	CreateTimestamp string `json:"create_timestamp,omitempty"`

	// The ID of the project.
	// Example: 1
	ID string `json:"id,omitempty"`

	// The name of the project.
	// Example: default_project
	Name string `json:"name,omitempty"`

	// The ID of the TiDB Cloud organization to which the project belongs.
	// Example: 1
	OrgID string `json:"org_id,omitempty"`

	// The number of users in the project.
	// Example: 10
	UserCount int64 `json:"user_count,omitempty"`
}

// Validate validates this openapi list project item
func (m *OpenapiListProjectItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list project item based on context it is used
func (m *OpenapiListProjectItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProjectItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProjectItem) UnmarshalBinary(b []byte) error {
	var res OpenapiListProjectItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}