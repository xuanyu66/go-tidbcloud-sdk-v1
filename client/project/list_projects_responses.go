// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListProjectsReader is a Reader for the ListProjects structure.
type ListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListProjectsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectsOK creates a ListProjectsOK with default headers values
func NewListProjectsOK() *ListProjectsOK {
	return &ListProjectsOK{}
}

/*
ListProjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProjectsOK struct {
	Payload *ListProjectsOKBody
}

// IsSuccess returns true when this list projects o k response has a 2xx status code
func (o *ListProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list projects o k response has a 3xx status code
func (o *ListProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects o k response has a 4xx status code
func (o *ListProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list projects o k response has a 5xx status code
func (o *ListProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects o k response a status code equal to that given
func (o *ListProjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list projects o k response
func (o *ListProjectsOK) Code() int {
	return 200
}

func (o *ListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) GetPayload() *ListProjectsOKBody {
	return o.Payload
}

func (o *ListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListProjectsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsBadRequest creates a ListProjectsBadRequest with default headers values
func NewListProjectsBadRequest() *ListProjectsBadRequest {
	return &ListProjectsBadRequest{}
}

/*
ListProjectsBadRequest describes a response with status code 400, with default header values.

A request field is invalid.
*/
type ListProjectsBadRequest struct {
	Payload *ListProjectsBadRequestBody
}

// IsSuccess returns true when this list projects bad request response has a 2xx status code
func (o *ListProjectsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects bad request response has a 3xx status code
func (o *ListProjectsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects bad request response has a 4xx status code
func (o *ListProjectsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects bad request response has a 5xx status code
func (o *ListProjectsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects bad request response a status code equal to that given
func (o *ListProjectsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list projects bad request response
func (o *ListProjectsBadRequest) Code() int {
	return 400
}

func (o *ListProjectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectsBadRequest) GetPayload() *ListProjectsBadRequestBody {
	return o.Payload
}

func (o *ListProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListProjectsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsUnauthorized creates a ListProjectsUnauthorized with default headers values
func NewListProjectsUnauthorized() *ListProjectsUnauthorized {
	return &ListProjectsUnauthorized{}
}

/*
ListProjectsUnauthorized describes a response with status code 401, with default header values.

The API key cannot be authenticated.
*/
type ListProjectsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this list projects unauthorized response has a 2xx status code
func (o *ListProjectsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects unauthorized response has a 3xx status code
func (o *ListProjectsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects unauthorized response has a 4xx status code
func (o *ListProjectsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects unauthorized response has a 5xx status code
func (o *ListProjectsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects unauthorized response a status code equal to that given
func (o *ListProjectsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list projects unauthorized response
func (o *ListProjectsUnauthorized) Code() int {
	return 401
}

func (o *ListProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProjectsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProjectsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsForbidden creates a ListProjectsForbidden with default headers values
func NewListProjectsForbidden() *ListProjectsForbidden {
	return &ListProjectsForbidden{}
}

/*
ListProjectsForbidden describes a response with status code 403, with default header values.

The API key does not have permission to access the resource.
*/
type ListProjectsForbidden struct {
	Payload *ListProjectsForbiddenBody
}

// IsSuccess returns true when this list projects forbidden response has a 2xx status code
func (o *ListProjectsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects forbidden response has a 3xx status code
func (o *ListProjectsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects forbidden response has a 4xx status code
func (o *ListProjectsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects forbidden response has a 5xx status code
func (o *ListProjectsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects forbidden response a status code equal to that given
func (o *ListProjectsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list projects forbidden response
func (o *ListProjectsForbidden) Code() int {
	return 403
}

func (o *ListProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectsForbidden) GetPayload() *ListProjectsForbiddenBody {
	return o.Payload
}

func (o *ListProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListProjectsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsNotFound creates a ListProjectsNotFound with default headers values
func NewListProjectsNotFound() *ListProjectsNotFound {
	return &ListProjectsNotFound{}
}

/*
ListProjectsNotFound describes a response with status code 404, with default header values.

The requested resource does not exist.
*/
type ListProjectsNotFound struct {
	Payload *ListProjectsNotFoundBody
}

// IsSuccess returns true when this list projects not found response has a 2xx status code
func (o *ListProjectsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects not found response has a 3xx status code
func (o *ListProjectsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects not found response has a 4xx status code
func (o *ListProjectsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects not found response has a 5xx status code
func (o *ListProjectsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects not found response a status code equal to that given
func (o *ListProjectsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list projects not found response
func (o *ListProjectsNotFound) Code() int {
	return 404
}

func (o *ListProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectsNotFound) GetPayload() *ListProjectsNotFoundBody {
	return o.Payload
}

func (o *ListProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListProjectsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsTooManyRequests creates a ListProjectsTooManyRequests with default headers values
func NewListProjectsTooManyRequests() *ListProjectsTooManyRequests {
	return &ListProjectsTooManyRequests{}
}

/*
ListProjectsTooManyRequests describes a response with status code 429, with default header values.

You have exceed the rate limit.
*/
type ListProjectsTooManyRequests struct {
	Payload *ListProjectsTooManyRequestsBody
}

// IsSuccess returns true when this list projects too many requests response has a 2xx status code
func (o *ListProjectsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects too many requests response has a 3xx status code
func (o *ListProjectsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects too many requests response has a 4xx status code
func (o *ListProjectsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects too many requests response has a 5xx status code
func (o *ListProjectsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects too many requests response a status code equal to that given
func (o *ListProjectsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the list projects too many requests response
func (o *ListProjectsTooManyRequests) Code() int {
	return 429
}

func (o *ListProjectsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListProjectsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListProjectsTooManyRequests) GetPayload() *ListProjectsTooManyRequestsBody {
	return o.Payload
}

func (o *ListProjectsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListProjectsTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsInternalServerError creates a ListProjectsInternalServerError with default headers values
func NewListProjectsInternalServerError() *ListProjectsInternalServerError {
	return &ListProjectsInternalServerError{}
}

/*
ListProjectsInternalServerError describes a response with status code 500, with default header values.

Server error.
*/
type ListProjectsInternalServerError struct {
	Payload *ListProjectsInternalServerErrorBody
}

// IsSuccess returns true when this list projects internal server error response has a 2xx status code
func (o *ListProjectsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects internal server error response has a 3xx status code
func (o *ListProjectsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects internal server error response has a 4xx status code
func (o *ListProjectsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list projects internal server error response has a 5xx status code
func (o *ListProjectsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list projects internal server error response a status code equal to that given
func (o *ListProjectsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list projects internal server error response
func (o *ListProjectsInternalServerError) Code() int {
	return 500
}

func (o *ListProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] listProjectsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectsInternalServerError) GetPayload() *ListProjectsInternalServerErrorBody {
	return o.Payload
}

func (o *ListProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListProjectsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsDefault creates a ListProjectsDefault with default headers values
func NewListProjectsDefault(code int) *ListProjectsDefault {
	return &ListProjectsDefault{
		_statusCode: code,
	}
}

/*
ListProjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProjectsDefault struct {
	_statusCode int

	Payload *ListProjectsDefaultBody
}

// IsSuccess returns true when this list projects default response has a 2xx status code
func (o *ListProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list projects default response has a 3xx status code
func (o *ListProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list projects default response has a 4xx status code
func (o *ListProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list projects default response has a 5xx status code
func (o *ListProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list projects default response a status code equal to that given
func (o *ListProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list projects default response
func (o *ListProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] ListProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1beta/projects][%d] ListProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectsDefault) GetPayload() *ListProjectsDefaultBody {
	return o.Payload
}

func (o *ListProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListProjectsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListProjectsBadRequestBody list projects bad request body
swagger:model ListProjectsBadRequestBody
*/
type ListProjectsBadRequestBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this list projects bad request body
func (o *ListProjectsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list projects bad request body based on context it is used
func (o *ListProjectsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ListProjectsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsDefaultBody list projects default body
swagger:model ListProjectsDefaultBody
*/
type ListProjectsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ListProjectsDefaultBodyDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this list projects default body
func (o *ListProjectsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListProjectsDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListProjects default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListProjects default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list projects default body based on the context it is used
func (o *ListProjectsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListProjectsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListProjects default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListProjects default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListProjectsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsDefaultBodyDetailsItems0 list projects default body details items0
swagger:model ListProjectsDefaultBodyDetailsItems0
*/
type ListProjectsDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list projects default body details items0
func (o *ListProjectsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list projects default body details items0 based on context it is used
func (o *ListProjectsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListProjectsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsForbiddenBody list projects forbidden body
swagger:model ListProjectsForbiddenBody
*/
type ListProjectsForbiddenBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this list projects forbidden body
func (o *ListProjectsForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list projects forbidden body based on context it is used
func (o *ListProjectsForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ListProjectsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsInternalServerErrorBody list projects internal server error body
swagger:model ListProjectsInternalServerErrorBody
*/
type ListProjectsInternalServerErrorBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this list projects internal server error body
func (o *ListProjectsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list projects internal server error body based on context it is used
func (o *ListProjectsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ListProjectsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsNotFoundBody list projects not found body
swagger:model ListProjectsNotFoundBody
*/
type ListProjectsNotFoundBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this list projects not found body
func (o *ListProjectsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list projects not found body based on context it is used
func (o *ListProjectsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListProjectsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsOKBody GetProjectsResp
//
// GetProjectsResp is the response for getting accessible projects.
swagger:model ListProjectsOKBody
*/
type ListProjectsOKBody struct {

	// The items of accessible projects.
	// Required: true
	Items []*ListProjectsOKBodyItemsItems0 `json:"items"`

	// The total number of accessible projects.
	// Example: 1
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this list projects o k body
func (o *ListProjectsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListProjectsOKBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("listProjectsOK"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listProjectsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listProjectsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListProjectsOKBody) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("listProjectsOK"+"."+"total", "body", o.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list projects o k body based on the context it is used
func (o *ListProjectsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListProjectsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listProjectsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listProjectsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsOKBody) UnmarshalBinary(b []byte) error {
	var res ListProjectsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsOKBodyItemsItems0 ListProjectItem
//
// ListProjectItem is the item of projects.
swagger:model ListProjectsOKBodyItemsItems0
*/
type ListProjectsOKBodyItemsItems0 struct {

	// Flag that indicates whether to enable AWS Customer-Managed Encryption Keys (CMEK). For more information, see [Encryption at Rest using CMEK](https://docs.pingcap.com/tidbcloud/tidb-cloud-encrypt-cmek).
	//
	// **Note:** Currently, this feature is only available upon request. If you need to try out this feature, contact [support](https://docs.pingcap.com/tidbcloud/tidb-cloud-support).
	// Example: false
	AwsCmekEnabled *bool `json:"aws_cmek_enabled,omitempty"`

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

// Validate validates this list projects o k body items items0
func (o *ListProjectsOKBodyItemsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list projects o k body items items0 based on context it is used
func (o *ListProjectsOKBodyItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsOKBodyItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsOKBodyItemsItems0) UnmarshalBinary(b []byte) error {
	var res ListProjectsOKBodyItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListProjectsTooManyRequestsBody list projects too many requests body
swagger:model ListProjectsTooManyRequestsBody
*/
type ListProjectsTooManyRequestsBody struct {

	// code
	//
	// Error code returned with this error.
	Code int64 `json:"code,omitempty"`

	// details
	//
	// Error details returned with this error.
	Details []string `json:"details"`

	// message
	//
	// Error message returned with this error.
	Message string `json:"message,omitempty"`
}

// Validate validates this list projects too many requests body
func (o *ListProjectsTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list projects too many requests body based on context it is used
func (o *ListProjectsTooManyRequestsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListProjectsTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListProjectsTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res ListProjectsTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
