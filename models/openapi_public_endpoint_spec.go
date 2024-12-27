// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiPublicEndpointSpec openapi public endpoint spec
//
// swagger:model openapiPublicEndpointSpec
type OpenapiPublicEndpointSpec struct {

	// The IP allowlist for the public endpoint.
	IPAllowlist []*OpenapiPublicEndpointSpecIPAllowlistItems0 `json:"ip_allowlist"`
}

// Validate validates this openapi public endpoint spec
func (m *OpenapiPublicEndpointSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAllowlist(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiPublicEndpointSpec) validateIPAllowlist(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAllowlist) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAllowlist); i++ {
		if swag.IsZero(m.IPAllowlist[i]) { // not required
			continue
		}

		if m.IPAllowlist[i] != nil {
			if err := m.IPAllowlist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this openapi public endpoint spec based on the context it is used
func (m *OpenapiPublicEndpointSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAllowlist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiPublicEndpointSpec) contextValidateIPAllowlist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAllowlist); i++ {

		if m.IPAllowlist[i] != nil {

			if swag.IsZero(m.IPAllowlist[i]) { // not required
				return nil
			}

			if err := m.IPAllowlist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiPublicEndpointSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiPublicEndpointSpec) UnmarshalBinary(b []byte) error {
	var res OpenapiPublicEndpointSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiPublicEndpointSpecIPAllowlistItems0 openapi public endpoint spec IP allowlist items0
//
// swagger:model OpenapiPublicEndpointSpecIPAllowlistItems0
type OpenapiPublicEndpointSpecIPAllowlistItems0 struct {

	// CIDR.
	// Example: 0.0.0.0/0
	Cidr string `json:"cidr,omitempty"`
}

// Validate validates this openapi public endpoint spec IP allowlist items0
func (m *OpenapiPublicEndpointSpecIPAllowlistItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi public endpoint spec IP allowlist items0 based on context it is used
func (m *OpenapiPublicEndpointSpecIPAllowlistItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiPublicEndpointSpecIPAllowlistItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiPublicEndpointSpecIPAllowlistItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiPublicEndpointSpecIPAllowlistItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}