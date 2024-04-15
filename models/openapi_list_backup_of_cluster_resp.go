// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiListBackupOfClusterResp ListBackupOfClusterResp
//
// The response for listing backups of a cluster.
//
// swagger:model openapiListBackupOfClusterResp
type OpenapiListBackupOfClusterResp struct {

	// The items of all backups.
	Items []*OpenapiListBackupOfClusterRespItemsItems0 `json:"items"`

	// The total number of backups in the project.
	// Example: 10
	Total int64 `json:"total,omitempty"`
}

// Validate validates this openapi list backup of cluster resp
func (m *OpenapiListBackupOfClusterResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListBackupOfClusterResp) validateItems(formats strfmt.Registry) error {
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

// ContextValidate validate this openapi list backup of cluster resp based on the context it is used
func (m *OpenapiListBackupOfClusterResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListBackupOfClusterResp) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {

			if swag.IsZero(m.Items[i]) { // not required
				return nil
			}

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
func (m *OpenapiListBackupOfClusterResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListBackupOfClusterResp) UnmarshalBinary(b []byte) error {
	var res OpenapiListBackupOfClusterResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListBackupOfClusterRespItemsItems0 ListBackupItem
//
// The item of backup list.
//
// swagger:model OpenapiListBackupOfClusterRespItemsItems0
type OpenapiListBackupOfClusterRespItemsItems0 struct {

	// The creation time of the backup in UTC. The time format follows the [ISO8601](http://en.wikipedia.org/wiki/ISO_8601) standard, which is `YYYY-MM-DD` (year-month-day) + T +`HH:MM:SS` (hour-minutes-seconds) + Z. For example, `2020-01-01T00:00:00Z`.
	// Example: 2020-01-01T00:00:00Z
	// Format: date-time
	CreateTimestamp strfmt.DateTime `json:"create_timestamp,omitempty"`

	// The description of the backup. It is specified by the user when taking a manual type backup. It helps you add additional information to the backup.
	// Example: backup for cluster upgrade in 2022/06/07
	Description string `json:"description,omitempty"`

	// The ID of the backup. It is generated by TiDB Cloud.
	// Example: 1
	ID string `json:"id,omitempty"`

	// The name of the backup.
	// Example: backup-1
	Name string `json:"name,omitempty"`

	// The bytes of the backup.
	// Example: 102400
	Size string `json:"size,omitempty"`

	// The status of backup.
	// Example: SUCCESS
	// Enum: [PENDING RUNNING FAILED SUCCESS]
	Status string `json:"status,omitempty"`

	// The type of backup. TiDB Cloud only supports manual and auto backup. For more information, see [TiDB Cloud Documentation](https://docs.pingcap.com/tidbcloud/backup-and-restore#backup).
	// Example: MANUAL
	// Enum: [MANUAL AUTO]
	Type string `json:"type,omitempty"`
}

// Validate validates this openapi list backup of cluster resp items items0
func (m *OpenapiListBackupOfClusterRespItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListBackupOfClusterRespItemsItems0) validateCreateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("create_timestamp", "body", "date-time", m.CreateTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var openapiListBackupOfClusterRespItemsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","RUNNING","FAILED","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiListBackupOfClusterRespItemsItems0TypeStatusPropEnum = append(openapiListBackupOfClusterRespItemsItems0TypeStatusPropEnum, v)
	}
}

const (

	// OpenapiListBackupOfClusterRespItemsItems0StatusPENDING captures enum value "PENDING"
	OpenapiListBackupOfClusterRespItemsItems0StatusPENDING string = "PENDING"

	// OpenapiListBackupOfClusterRespItemsItems0StatusRUNNING captures enum value "RUNNING"
	OpenapiListBackupOfClusterRespItemsItems0StatusRUNNING string = "RUNNING"

	// OpenapiListBackupOfClusterRespItemsItems0StatusFAILED captures enum value "FAILED"
	OpenapiListBackupOfClusterRespItemsItems0StatusFAILED string = "FAILED"

	// OpenapiListBackupOfClusterRespItemsItems0StatusSUCCESS captures enum value "SUCCESS"
	OpenapiListBackupOfClusterRespItemsItems0StatusSUCCESS string = "SUCCESS"
)

// prop value enum
func (m *OpenapiListBackupOfClusterRespItemsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiListBackupOfClusterRespItemsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiListBackupOfClusterRespItemsItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var openapiListBackupOfClusterRespItemsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiListBackupOfClusterRespItemsItems0TypeTypePropEnum = append(openapiListBackupOfClusterRespItemsItems0TypeTypePropEnum, v)
	}
}

const (

	// OpenapiListBackupOfClusterRespItemsItems0TypeMANUAL captures enum value "MANUAL"
	OpenapiListBackupOfClusterRespItemsItems0TypeMANUAL string = "MANUAL"

	// OpenapiListBackupOfClusterRespItemsItems0TypeAUTO captures enum value "AUTO"
	OpenapiListBackupOfClusterRespItemsItems0TypeAUTO string = "AUTO"
)

// prop value enum
func (m *OpenapiListBackupOfClusterRespItemsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiListBackupOfClusterRespItemsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiListBackupOfClusterRespItemsItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi list backup of cluster resp items items0 based on context it is used
func (m *OpenapiListBackupOfClusterRespItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListBackupOfClusterRespItemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListBackupOfClusterRespItemsItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListBackupOfClusterRespItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
