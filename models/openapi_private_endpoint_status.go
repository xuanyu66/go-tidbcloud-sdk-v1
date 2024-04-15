// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OpenapiPrivateEndpointStatus openapi private endpoint status
//
// swagger:model openapiPrivateEndpointStatus
type OpenapiPrivateEndpointStatus string

func NewOpenapiPrivateEndpointStatus(value OpenapiPrivateEndpointStatus) *OpenapiPrivateEndpointStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OpenapiPrivateEndpointStatus.
func (m OpenapiPrivateEndpointStatus) Pointer() *OpenapiPrivateEndpointStatus {
	return &m
}

const (

	// OpenapiPrivateEndpointStatusPENDING captures enum value "PENDING"
	OpenapiPrivateEndpointStatusPENDING OpenapiPrivateEndpointStatus = "PENDING"

	// OpenapiPrivateEndpointStatusACTIVE captures enum value "ACTIVE"
	OpenapiPrivateEndpointStatusACTIVE OpenapiPrivateEndpointStatus = "ACTIVE"

	// OpenapiPrivateEndpointStatusDELETING captures enum value "DELETING"
	OpenapiPrivateEndpointStatusDELETING OpenapiPrivateEndpointStatus = "DELETING"

	// OpenapiPrivateEndpointStatusFAILED captures enum value "FAILED"
	OpenapiPrivateEndpointStatusFAILED OpenapiPrivateEndpointStatus = "FAILED"
)

// for schema
var openapiPrivateEndpointStatusEnum []interface{}

func init() {
	var res []OpenapiPrivateEndpointStatus
	if err := json.Unmarshal([]byte(`["PENDING","ACTIVE","DELETING","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiPrivateEndpointStatusEnum = append(openapiPrivateEndpointStatusEnum, v)
	}
}

func (m OpenapiPrivateEndpointStatus) validateOpenapiPrivateEndpointStatusEnum(path, location string, value OpenapiPrivateEndpointStatus) error {
	if err := validate.EnumCase(path, location, value, openapiPrivateEndpointStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openapi private endpoint status
func (m OpenapiPrivateEndpointStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenapiPrivateEndpointStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openapi private endpoint status based on context it is used
func (m OpenapiPrivateEndpointStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}