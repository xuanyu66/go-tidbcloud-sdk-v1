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

// OpenapiListAccesspointsResp ListAccesspointsResp
//
// ListAccesspointsResp is the response for listing accesspoints for a cluster.
//
// swagger:model openapiListAccesspointsResp
type OpenapiListAccesspointsResp struct {

	// The accesspoints for the cluster.
	// Required: true
	Items []*OpenapiListAccesspointsRespItemsItems0 `json:"items"`

	// The total number of accesspoints.
	// Example: 1
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this openapi list accesspoints resp
func (m *OpenapiListAccesspointsResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsResp) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
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

func (m *OpenapiListAccesspointsResp) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openapi list accesspoints resp based on the context it is used
func (m *OpenapiListAccesspointsResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsResp) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

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
func (m *OpenapiListAccesspointsResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAccesspointsResp) UnmarshalBinary(b []byte) error {
	var res OpenapiListAccesspointsResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListAccesspointsRespItemsItems0 Accesspoint
//
// # Accesspoint
//
// swagger:model OpenapiListAccesspointsRespItemsItems0
type OpenapiListAccesspointsRespItemsItems0 struct {

	// The ID of the cluster.
	// Example: 1
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// The display name of the accesspoint. The name must be 4-64 characters that can only include numbers, letters, and hyphens, and the first and last character must be a letter or number.
	// Example: Accesspoint
	// Required: true
	// Pattern: ^[A-Za-z0-9][-A-Za-z0-9]{2,62}[A-Za-z0-9]$
	DisplayName *string `json:"display_name"`

	// The ID of the accesspoint.
	// Example: 1
	// Required: true
	ID *string `json:"id"`

	// private endpoint
	PrivateEndpoint *OpenapiListAccesspointsRespItemsItems0PrivateEndpoint `json:"private_endpoint,omitempty"`

	// The ID of the project. You can get the project ID from the response of [List all accessible projects](#tag/Project/operation/ListProjects).
	// Example: 1
	// Required: true
	ProjectID *string `json:"project_id"`

	// public endpoint
	PublicEndpoint *OpenapiListAccesspointsRespItemsItems0PublicEndpoint `json:"public_endpoint,omitempty"`

	// The number of tidb nodes in the accesspoint.
	// Example: 2
	TidbNodeQuantity int32 `json:"tidb_node_quantity,omitempty"`

	// vpc peering endpoint
	VpcPeeringEndpoint *OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint `json:"vpc_peering_endpoint,omitempty"`
}

// Validate validates this openapi list accesspoints resp items items0
func (m *OpenapiListAccesspointsRespItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcPeeringEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	if err := validate.Pattern("display_name", "body", *m.DisplayName, `^[A-Za-z0-9][-A-Za-z0-9]{2,62}[A-Za-z0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) validatePrivateEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateEndpoint) { // not required
		return nil
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

func (m *OpenapiListAccesspointsRespItemsItems0) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("project_id", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) validatePublicEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicEndpoint) { // not required
		return nil
	}

	if m.PublicEndpoint != nil {
		if err := m.PublicEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) validateVpcPeeringEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcPeeringEndpoint) { // not required
		return nil
	}

	if m.VpcPeeringEndpoint != nil {
		if err := m.VpcPeeringEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_peering_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_peering_endpoint")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi list accesspoints resp items items0 based on the context it is used
func (m *OpenapiListAccesspointsRespItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpcPeeringEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) contextValidatePrivateEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivateEndpoint != nil {

		if swag.IsZero(m.PrivateEndpoint) { // not required
			return nil
		}

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

func (m *OpenapiListAccesspointsRespItemsItems0) contextValidatePublicEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicEndpoint != nil {

		if swag.IsZero(m.PublicEndpoint) { // not required
			return nil
		}

		if err := m.PublicEndpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0) contextValidateVpcPeeringEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcPeeringEndpoint != nil {

		if swag.IsZero(m.VpcPeeringEndpoint) { // not required
			return nil
		}

		if err := m.VpcPeeringEndpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_peering_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_peering_endpoint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListAccesspointsRespItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListAccesspointsRespItemsItems0PrivateEndpoint The private endpoint.
//
// swagger:model OpenapiListAccesspointsRespItemsItems0PrivateEndpoint
type OpenapiListAccesspointsRespItemsItems0PrivateEndpoint struct {

	// Availability zones for the private endpoint service. This field is only applicable when the `cloud_provider` is `"AWS"`.
	// Example: ["usw2-az2","usw2-az1"]
	AzIds []string `json:"az_ids"`

	// The DNS name of the private endpoint service.
	// Example: privatelink-tidb.01234567890.clusters.tidb-cloud.com
	Host string `json:"host,omitempty"`

	// The port of the private endpoint service.
	// Example: 4000
	Port int32 `json:"port,omitempty"`

	// private link service
	PrivateLinkService *OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService `json:"private_link_service,omitempty"`
}

// Validate validates this openapi list accesspoints resp items items0 private endpoint
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateLinkService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpoint) validatePrivateLinkService(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateLinkService) { // not required
		return nil
	}

	if m.PrivateLinkService != nil {
		if err := m.PrivateLinkService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_endpoint" + "." + "private_link_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("private_endpoint" + "." + "private_link_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi list accesspoints resp items items0 private endpoint based on the context it is used
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateLinkService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpoint) contextValidatePrivateLinkService(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivateLinkService != nil {

		if swag.IsZero(m.PrivateLinkService) { // not required
			return nil
		}

		if err := m.PrivateLinkService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_endpoint" + "." + "private_link_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("private_endpoint" + "." + "private_link_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpoint) UnmarshalBinary(b []byte) error {
	var res OpenapiListAccesspointsRespItemsItems0PrivateEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService The private endpoint service.
//
// swagger:model OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService
type OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService struct {

	// The name of the private endpoint service, which is used for connection.
	// Example: com.amazonaws.vpce.us-west-2.vpce-svc-01234567890123456
	ServiceName string `json:"service_name,omitempty"`

	// The status of the private endpoint service.
	// Example: ACTIVE
	// Enum: ["CREATING","ACTIVE","DELETING"]
	Status string `json:"status,omitempty"`
}

// Validate validates this openapi list accesspoints resp items items0 private endpoint private link service
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceTypeStatusPropEnum = append(openapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceTypeStatusPropEnum, v)
	}
}

const (

	// OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceStatusCREATING captures enum value "CREATING"
	OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceStatusCREATING string = "CREATING"

	// OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceStatusACTIVE captures enum value "ACTIVE"
	OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceStatusACTIVE string = "ACTIVE"

	// OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceStatusDELETING captures enum value "DELETING"
	OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceStatusDELETING string = "DELETING"
)

// prop value enum
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkServiceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("private_endpoint"+"."+"private_link_service"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi list accesspoints resp items items0 private endpoint private link service based on context it is used
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService) UnmarshalBinary(b []byte) error {
	var res OpenapiListAccesspointsRespItemsItems0PrivateEndpointPrivateLinkService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListAccesspointsRespItemsItems0PublicEndpoint The public endpoint.
//
// swagger:model OpenapiListAccesspointsRespItemsItems0PublicEndpoint
type OpenapiListAccesspointsRespItemsItems0PublicEndpoint struct {

	// The host of standard connection.
	// Example: tidb.f69f3808.acea1f2a.us-east-1.shared.aws.tidbcloud.com
	Host string `json:"host,omitempty"`

	// The IP allowlist for the public endpoint.
	IPAllowlist []*OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0 `json:"ip_allowlist"`

	// The TiDB port for connection. The port must be in the range of 1024-65535 except 10080.
	//
	// Example: 4000
	// Maximum: 65535
	// Minimum: 1024
	Port int32 `json:"port,omitempty"`
}

// Validate validates this openapi list accesspoints resp items items0 public endpoint
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAllowlist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpoint) validateIPAllowlist(formats strfmt.Registry) error {
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
					return ve.ValidateName("public_endpoint" + "." + "ip_allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("public_endpoint" + "." + "ip_allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpoint) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("public_endpoint"+"."+"port", "body", int64(m.Port), 1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("public_endpoint"+"."+"port", "body", int64(m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openapi list accesspoints resp items items0 public endpoint based on the context it is used
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAllowlist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpoint) contextValidateIPAllowlist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAllowlist); i++ {

		if m.IPAllowlist[i] != nil {

			if swag.IsZero(m.IPAllowlist[i]) { // not required
				return nil
			}

			if err := m.IPAllowlist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("public_endpoint" + "." + "ip_allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("public_endpoint" + "." + "ip_allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpoint) UnmarshalBinary(b []byte) error {
	var res OpenapiListAccesspointsRespItemsItems0PublicEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0 openapi list accesspoints resp items items0 public endpoint IP allowlist items0
//
// swagger:model OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0
type OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0 struct {

	// CIDR.
	// Example: 0.0.0.0/0
	Cidr string `json:"cidr,omitempty"`
}

// Validate validates this openapi list accesspoints resp items items0 public endpoint IP allowlist items0
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list accesspoints resp items items0 public endpoint IP allowlist items0 based on context it is used
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListAccesspointsRespItemsItems0PublicEndpointIPAllowlistItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint The VPC peering endpoint.
//
// swagger:model OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint
type OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint struct {

	// The host of VPC peering connection.
	// Example: private-tidb.f69f3808.acea1f2a.us-east-1.shared.aws.tidbcloud.com
	Host string `json:"host,omitempty"`

	// The TiDB port for connection. The port must be in the range of 1024-65535 except 10080.
	//
	// Example: 4000
	// Maximum: 65535
	// Minimum: 1024
	Port int32 `json:"port,omitempty"`
}

// Validate validates this openapi list accesspoints resp items items0 vpc peering endpoint
func (m *OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("vpc_peering_endpoint"+"."+"port", "body", int64(m.Port), 1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vpc_peering_endpoint"+"."+"port", "body", int64(m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi list accesspoints resp items items0 vpc peering endpoint based on context it is used
func (m *OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint) UnmarshalBinary(b []byte) error {
	var res OpenapiListAccesspointsRespItemsItems0VpcPeeringEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
