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

// ImportStatusImportTaskPhase import status import task phase
//
// swagger:model ImportStatusImportTaskPhase
type ImportStatusImportTaskPhase string

func NewImportStatusImportTaskPhase(value ImportStatusImportTaskPhase) *ImportStatusImportTaskPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ImportStatusImportTaskPhase.
func (m ImportStatusImportTaskPhase) Pointer() *ImportStatusImportTaskPhase {
	return &m
}

const (

	// ImportStatusImportTaskPhasePREPARING captures enum value "PREPARING"
	ImportStatusImportTaskPhasePREPARING ImportStatusImportTaskPhase = "PREPARING"

	// ImportStatusImportTaskPhaseIMPORTING captures enum value "IMPORTING"
	ImportStatusImportTaskPhaseIMPORTING ImportStatusImportTaskPhase = "IMPORTING"

	// ImportStatusImportTaskPhaseCOMPLETED captures enum value "COMPLETED"
	ImportStatusImportTaskPhaseCOMPLETED ImportStatusImportTaskPhase = "COMPLETED"

	// ImportStatusImportTaskPhaseFAILED captures enum value "FAILED"
	ImportStatusImportTaskPhaseFAILED ImportStatusImportTaskPhase = "FAILED"

	// ImportStatusImportTaskPhaseCANCELING captures enum value "CANCELING"
	ImportStatusImportTaskPhaseCANCELING ImportStatusImportTaskPhase = "CANCELING"

	// ImportStatusImportTaskPhaseCANCELED captures enum value "CANCELED"
	ImportStatusImportTaskPhaseCANCELED ImportStatusImportTaskPhase = "CANCELED"
)

// for schema
var importStatusImportTaskPhaseEnum []interface{}

func init() {
	var res []ImportStatusImportTaskPhase
	if err := json.Unmarshal([]byte(`["PREPARING","IMPORTING","COMPLETED","FAILED","CANCELING","CANCELED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		importStatusImportTaskPhaseEnum = append(importStatusImportTaskPhaseEnum, v)
	}
}

func (m ImportStatusImportTaskPhase) validateImportStatusImportTaskPhaseEnum(path, location string, value ImportStatusImportTaskPhase) error {
	if err := validate.EnumCase(path, location, value, importStatusImportTaskPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this import status import task phase
func (m ImportStatusImportTaskPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateImportStatusImportTaskPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this import status import task phase based on context it is used
func (m ImportStatusImportTaskPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
