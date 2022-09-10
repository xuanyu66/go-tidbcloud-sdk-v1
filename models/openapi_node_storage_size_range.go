// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiNodeStorageSizeRange openapi node storage size range
//
// swagger:model openapiNodeStorageSizeRange
type OpenapiNodeStorageSizeRange struct {

	// The maximum storage size for each node of the component in the cluster.
	Max int32 `json:"max,omitempty"`

	// The minimum storage size for each node of the component in the cluster.
	Min int32 `json:"min,omitempty"`
}

// Validate validates this openapi node storage size range
func (m *OpenapiNodeStorageSizeRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi node storage size range based on context it is used
func (m *OpenapiNodeStorageSizeRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiNodeStorageSizeRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiNodeStorageSizeRange) UnmarshalBinary(b []byte) error {
	var res OpenapiNodeStorageSizeRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}