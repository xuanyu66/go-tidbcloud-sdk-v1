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

// OpenapiListProviderRegionsResp GetProviderRegionsResp
//
// GetProviderRegionsResp is the response for getting provider regions.
//
// swagger:model openapiListProviderRegionsResp
type OpenapiListProviderRegionsResp struct {

	// Items of provider regions.
	// Example: [{"cloud_provider":"AWS","cluster_type":"DEDICATED","region":"us-west-2","tidb":[{"node_quantity_range":{"min":1,"step":1},"node_size":"8C16G"}],"tiflash":[{"node_quantity_range":{"min":0,"step":1},"node_size":"8C64G","storage_size_gib_range":{"max":2048,"min":500}}],"tikv":[{"node_quantity_range":{"min":3,"step":3},"node_size":"8C32G","storage_size_gib_range":{"max":4096,"min":500}}]},{"cloud_provider":"AWS","cluster_type":"DEVELOPER","region":"us-west-2","tidb":[{"node_quantity_range":{"min":1,"step":1},"node_size":"Shared0"}],"tiflash":[{"node_quantity_range":{"min":1,"step":1},"node_size":"Shared0","storage_size_gib_range":{"max":1,"min":1}}],"tikv":[{"node_quantity_range":{"min":1,"step":1},"node_size":"Shared0","storage_size_gib_range":{"max":1,"min":1}}]}]
	Items []*OpenapiListProviderRegionsRespItemsItems0 `json:"items"`
}

// Validate validates this openapi list provider regions resp
func (m *OpenapiListProviderRegionsResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsResp) validateItems(formats strfmt.Registry) error {
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

// ContextValidate validate this openapi list provider regions resp based on the context it is used
func (m *OpenapiListProviderRegionsResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsResp) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

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
func (m *OpenapiListProviderRegionsResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsResp) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0 ListProviderRegionsItem
//
// ListProviderRegionsItem is the item of provider regions.
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0
type OpenapiListProviderRegionsRespItemsItems0 struct {

	// The cloud provider on which your TiDB cluster is hosted.
	// - `"AWS"`: the Amazon Web Services cloud provider
	// - `"GCP"`: the Google Cloud Platform cloud provider
	// Example: AWS
	// Enum: [AWS GCP]
	CloudProvider string `json:"cloud_provider,omitempty"`

	// The cluster type.
	// - `"DEVELOPER"`: a [Serverless Tier](https://docs.pingcap.com/tidbcloud/select-cluster-tier#serverless-tier) cluster
	// - `"DEDICATED"`: a [Dedicated Tier](https://docs.pingcap.com/tidbcloud/select-cluster-tier#dedicated-tier) cluster
	//
	// **Warning:** `"DEVELOPER"` will soon be changed to `"SERVERLESS"` to represent Serverless Tier clusters.
	// Example: DEDICATED
	// Enum: [DEDICATED DEVELOPER]
	ClusterType string `json:"cluster_type,omitempty"`

	// The region in which your TiDB cluster is hosted.
	//
	// For the detailed information on each region, refer to the documentation of the corresponding cloud provider ([AWS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html) | [GCP](https://cloud.google.com/about/locations#americas)).
	//
	// For example, `"us-west-2"` refers to Oregon for AWS.
	// Example: us-west-2
	Region string `json:"region,omitempty"`

	// The list of TiDB specifications in the region.
	Tidb []*OpenapiListProviderRegionsRespItemsItems0TidbItems0 `json:"tidb"`

	// The list of TiFlash specifications in the region.
	Tiflash []*OpenapiListProviderRegionsRespItemsItems0TiflashItems0 `json:"tiflash"`

	// The list of TiKV specifications in the region.
	Tikv []*OpenapiListProviderRegionsRespItemsItems0TikvItems0 `json:"tikv"`
}

// Validate validates this openapi list provider regions resp items items0
func (m *OpenapiListProviderRegionsRespItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTidb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTiflash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTikv(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openapiListProviderRegionsRespItemsItems0TypeCloudProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiListProviderRegionsRespItemsItems0TypeCloudProviderPropEnum = append(openapiListProviderRegionsRespItemsItems0TypeCloudProviderPropEnum, v)
	}
}

const (

	// OpenapiListProviderRegionsRespItemsItems0CloudProviderAWS captures enum value "AWS"
	OpenapiListProviderRegionsRespItemsItems0CloudProviderAWS string = "AWS"

	// OpenapiListProviderRegionsRespItemsItems0CloudProviderGCP captures enum value "GCP"
	OpenapiListProviderRegionsRespItemsItems0CloudProviderGCP string = "GCP"
)

// prop value enum
func (m *OpenapiListProviderRegionsRespItemsItems0) validateCloudProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiListProviderRegionsRespItemsItems0TypeCloudProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) validateCloudProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudProviderEnum("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

var openapiListProviderRegionsRespItemsItems0TypeClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEDICATED","DEVELOPER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiListProviderRegionsRespItemsItems0TypeClusterTypePropEnum = append(openapiListProviderRegionsRespItemsItems0TypeClusterTypePropEnum, v)
	}
}

const (

	// OpenapiListProviderRegionsRespItemsItems0ClusterTypeDEDICATED captures enum value "DEDICATED"
	OpenapiListProviderRegionsRespItemsItems0ClusterTypeDEDICATED string = "DEDICATED"

	// OpenapiListProviderRegionsRespItemsItems0ClusterTypeDEVELOPER captures enum value "DEVELOPER"
	OpenapiListProviderRegionsRespItemsItems0ClusterTypeDEVELOPER string = "DEVELOPER"
)

// prop value enum
func (m *OpenapiListProviderRegionsRespItemsItems0) validateClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiListProviderRegionsRespItemsItems0TypeClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) validateClusterType(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterTypeEnum("cluster_type", "body", m.ClusterType); err != nil {
		return err
	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) validateTidb(formats strfmt.Registry) error {
	if swag.IsZero(m.Tidb) { // not required
		return nil
	}

	for i := 0; i < len(m.Tidb); i++ {
		if swag.IsZero(m.Tidb[i]) { // not required
			continue
		}

		if m.Tidb[i] != nil {
			if err := m.Tidb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tidb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tidb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) validateTiflash(formats strfmt.Registry) error {
	if swag.IsZero(m.Tiflash) { // not required
		return nil
	}

	for i := 0; i < len(m.Tiflash); i++ {
		if swag.IsZero(m.Tiflash[i]) { // not required
			continue
		}

		if m.Tiflash[i] != nil {
			if err := m.Tiflash[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiflash" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiflash" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) validateTikv(formats strfmt.Registry) error {
	if swag.IsZero(m.Tikv) { // not required
		return nil
	}

	for i := 0; i < len(m.Tikv); i++ {
		if swag.IsZero(m.Tikv[i]) { // not required
			continue
		}

		if m.Tikv[i] != nil {
			if err := m.Tikv[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tikv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tikv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this openapi list provider regions resp items items0 based on the context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTidb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTiflash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTikv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) contextValidateTidb(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tidb); i++ {

		if m.Tidb[i] != nil {
			if err := m.Tidb[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tidb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tidb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) contextValidateTiflash(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tiflash); i++ {

		if m.Tiflash[i] != nil {
			if err := m.Tiflash[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiflash" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiflash" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0) contextValidateTikv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tikv); i++ {

		if m.Tikv[i] != nil {
			if err := m.Tikv[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tikv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tikv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TidbItems0 openapi list provider regions resp items items0 tidb items0
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TidbItems0
type OpenapiListProviderRegionsRespItemsItems0TidbItems0 struct {

	// node quantity range
	NodeQuantityRange *OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange `json:"node_quantity_range,omitempty"`

	// The size of the TiDB component in the cluster.
	// Example: 8C16G
	NodeSize string `json:"node_size,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tidb items0
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeQuantityRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0) validateNodeQuantityRange(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeQuantityRange) { // not required
		return nil
	}

	if m.NodeQuantityRange != nil {
		if err := m.NodeQuantityRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_quantity_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_quantity_range")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi list provider regions resp items items0 tidb items0 based on the context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeQuantityRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0) contextValidateNodeQuantityRange(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeQuantityRange != nil {
		if err := m.NodeQuantityRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_quantity_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_quantity_range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TidbItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange The range and step of node quantity of the TiDB component in the cluster.
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange
type OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange struct {

	// The minimum node quantity of the component in the cluster.
	Min int32 `json:"min,omitempty"`

	// The step of node quantity of the component in the cluster.
	Step int32 `json:"step,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tidb items0 node quantity range
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list provider regions resp items items0 tidb items0 node quantity range based on context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TidbItems0NodeQuantityRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TiflashItems0 openapi list provider regions resp items items0 tiflash items0
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TiflashItems0
type OpenapiListProviderRegionsRespItemsItems0TiflashItems0 struct {

	// node quantity range
	NodeQuantityRange *OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange `json:"node_quantity_range,omitempty"`

	// The size of the TiFlash component in the cluster.
	// Example: 8C64G
	NodeSize string `json:"node_size,omitempty"`

	// storage size gib range
	StorageSizeGibRange *OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange `json:"storage_size_gib_range,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tiflash items0
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeQuantityRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageSizeGibRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) validateNodeQuantityRange(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeQuantityRange) { // not required
		return nil
	}

	if m.NodeQuantityRange != nil {
		if err := m.NodeQuantityRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_quantity_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_quantity_range")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) validateStorageSizeGibRange(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageSizeGibRange) { // not required
		return nil
	}

	if m.StorageSizeGibRange != nil {
		if err := m.StorageSizeGibRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_size_gib_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_size_gib_range")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi list provider regions resp items items0 tiflash items0 based on the context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeQuantityRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageSizeGibRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) contextValidateNodeQuantityRange(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeQuantityRange != nil {
		if err := m.NodeQuantityRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_quantity_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_quantity_range")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) contextValidateStorageSizeGibRange(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageSizeGibRange != nil {
		if err := m.StorageSizeGibRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_size_gib_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_size_gib_range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TiflashItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange The range and step of node quantity of the TiFlash component in the cluster.
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange
type OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange struct {

	// The minimum node quantity of the component in the cluster.
	Min int32 `json:"min,omitempty"`

	// The step of node quantity of the component in the cluster.
	Step int32 `json:"step,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tiflash items0 node quantity range
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list provider regions resp items items0 tiflash items0 node quantity range based on context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TiflashItems0NodeQuantityRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange The storage size range for each node of the TiFlash component in the cluster.
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange
type OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange struct {

	// The maximum storage size for each node of the component in the cluster.
	Max int32 `json:"max,omitempty"`

	// The minimum storage size for each node of the component in the cluster.
	Min int32 `json:"min,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tiflash items0 storage size gib range
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list provider regions resp items items0 tiflash items0 storage size gib range based on context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TiflashItems0StorageSizeGibRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TikvItems0 openapi list provider regions resp items items0 tikv items0
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TikvItems0
type OpenapiListProviderRegionsRespItemsItems0TikvItems0 struct {

	// node quantity range
	NodeQuantityRange *OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange `json:"node_quantity_range,omitempty"`

	// The size of the TiKV component in the cluster.
	// Example: 8C32G
	NodeSize string `json:"node_size,omitempty"`

	// storage size gib range
	StorageSizeGibRange *OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange `json:"storage_size_gib_range,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tikv items0
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeQuantityRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageSizeGibRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) validateNodeQuantityRange(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeQuantityRange) { // not required
		return nil
	}

	if m.NodeQuantityRange != nil {
		if err := m.NodeQuantityRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_quantity_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_quantity_range")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) validateStorageSizeGibRange(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageSizeGibRange) { // not required
		return nil
	}

	if m.StorageSizeGibRange != nil {
		if err := m.StorageSizeGibRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_size_gib_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_size_gib_range")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openapi list provider regions resp items items0 tikv items0 based on the context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeQuantityRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageSizeGibRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) contextValidateNodeQuantityRange(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeQuantityRange != nil {
		if err := m.NodeQuantityRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_quantity_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_quantity_range")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) contextValidateStorageSizeGibRange(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageSizeGibRange != nil {
		if err := m.StorageSizeGibRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_size_gib_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_size_gib_range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TikvItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange The range and step of node quantity of the TiKV component in the cluster.
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange
type OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange struct {

	// The minimum node quantity of the component in the cluster.
	Min int32 `json:"min,omitempty"`

	// The step of node quantity of the component in the cluster.
	Step int32 `json:"step,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tikv items0 node quantity range
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list provider regions resp items items0 tikv items0 node quantity range based on context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TikvItems0NodeQuantityRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange The storage size range for each node of the TiKV component in the cluster.
//
// swagger:model OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange
type OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange struct {

	// The maximum storage size for each node of the component in the cluster.
	Max int32 `json:"max,omitempty"`

	// The minimum storage size for each node of the component in the cluster.
	Min int32 `json:"min,omitempty"`
}

// Validate validates this openapi list provider regions resp items items0 tikv items0 storage size gib range
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list provider regions resp items items0 tikv items0 storage size gib range based on context it is used
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange) UnmarshalBinary(b []byte) error {
	var res OpenapiListProviderRegionsRespItemsItems0TikvItems0StorageSizeGibRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
