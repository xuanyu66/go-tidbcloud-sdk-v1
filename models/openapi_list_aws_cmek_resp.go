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
	"github.com/go-openapi/validate"
)

// OpenapiListAwsCmekResp ListAwsCmekResp
//
// ListAwsCmekResp is the response for getting AWS Customer-Managed Encryption Keys for a project.
//
// swagger:model openapiListAwsCmekResp
type OpenapiListAwsCmekResp struct {

	// The specifications of the AWS CMEK.
	Items []*OpenapiListAwsCmekRespItemsItems0 `json:"items"`
}

// Validate validates this openapi list aws cmek resp
func (m *OpenapiListAwsCmekResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAwsCmekResp) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this openapi list aws cmek resp based on the context it is used
func (m *OpenapiListAwsCmekResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAwsCmekResp) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAwsCmekResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAwsCmekResp) UnmarshalBinary(b []byte) error {
	var res OpenapiListAwsCmekResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListAwsCmekRespItemsItems0 AwsCmekSpec
//
// AwsCmekSpec is the specification of the AWS CMEK.
//
// swagger:model OpenapiListAwsCmekRespItemsItems0
type OpenapiListAwsCmekRespItemsItems0 struct {

	// The KMS ARN of the AWS CMEK.
	// Example: arn:aws:kms:example
	// Required: true
	KmsArn *string `json:"kms_arn"`

	// The region name of the AWS CMEK. The region value should match the cloud provider's region code.
	//
	// You can get the complete list of available regions from the response of [List the cloud providers, regions and available specifications](#tag/Cluster/operation/ListProviderRegions).
	//
	// For the detailed information on each region, refer to the documentation of the [AWS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html) cloud provider.
	// Example: us-west-2
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this openapi list aws cmek resp items items0
func (m *OpenapiListAwsCmekRespItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKmsArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAwsCmekRespItemsItems0) validateKmsArn(formats strfmt.Registry) error {

	if err := validate.Required("kms_arn", "body", m.KmsArn); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiListAwsCmekRespItemsItems0) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi list aws cmek resp items items0 based on context it is used
func (m *OpenapiListAwsCmekRespItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAwsCmekRespItemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAwsCmekRespItemsItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListAwsCmekRespItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
