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

// EdgeClusterCreationSpec This specification contains the parameters required to add a NSX-T edge cluster spanning multiple VI clusters.
//
// swagger:model EdgeClusterCreationSpec
type EdgeClusterCreationSpec struct {

	// ASN to be used for the edge cluster
	Asn int64 `json:"asn,omitempty"`

	// Edge Password for admin user
	// Required: true
	EdgeAdminPassword *string `json:"edgeAdminPassword"`

	// Edge Password for audit
	// Required: true
	EdgeAuditPassword *string `json:"edgeAuditPassword"`

	// Name for the edge cluster.
	// Required: true
	EdgeClusterName *string `json:"edgeClusterName"`

	// Specifications for Edge Node
	// Required: true
	EdgeClusterProfileSpec *NsxTEdgeClusterProfileSpec `json:"edgeClusterProfileSpec"`

	// Type of edge cluster profile
	// Example: One among: DEFAULT, CUSTOM
	// Required: true
	EdgeClusterProfileType *string `json:"edgeClusterProfileType"`

	// Type of edge cluster
	// Example: One among: NSX-T
	// Required: true
	EdgeClusterType *string `json:"edgeClusterType"`

	// Edge Form Factor
	// Example: One among: XLARGE, LARGE, MEDIUM, SMALL
	// Required: true
	EdgeFormFactor *string `json:"edgeFormFactor"`

	// Specifications for Edge Node. Number of Edge Nodes cannot exceed 8 if HA mode is ACTIVE-ACTIVE and can not exceed 2 if HA mode is ACTIVE-STANDBY
	// Required: true
	EdgeNodeSpecs []*NsxTEdgeNodeSpec `json:"edgeNodeSpecs"`

	// Edge Password for root user.
	// Required: true
	EdgeRootPassword *string `json:"edgeRootPassword"`

	// Subnet addresses in CIDR notation that are used to assign addresses to logical links connecting service routers and distributed routers
	InternalTransitSubnets []string `json:"internalTransitSubnets"`

	// Maximum transmission unit
	// Required: true
	Mtu *int32 `json:"mtu"`

	// Name for the Tier-0
	// Required: true
	Tier0Name *string `json:"tier0Name"`

	// Tier 0 Routing type -eg eBGP, Static
	// Example: One among: EBGP, STATIC
	// Required: true
	Tier0RoutingType *string `json:"tier0RoutingType"`

	// High-availability Mode for Tier-0
	// Example: One among: ACTIVE_ACTIVE, ACTIVE_STANDBY
	// Required: true
	Tier0ServicesHighAvailability *string `json:"tier0ServicesHighAvailability"`

	// Name for the Tier-1
	// Required: true
	Tier1Name *string `json:"tier1Name"`

	// Select whether Tier-1 being created per this spec is hosted on the new Edge cluster or not (default value is false, meaning hosted)
	Tier1Unhosted bool `json:"tier1Unhosted,omitempty"`

	// Transit subnet addresses in CIDR notation that are used to assign addresses to logical links connecting Tier-0 and Tier-1s
	TransitSubnets []string `json:"transitSubnets"`
}

// Validate validates this edge cluster creation spec
func (m *EdgeClusterCreationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeAdminPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeAuditPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeClusterProfileSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeClusterProfileType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeClusterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeFormFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeNodeSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeRootPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier0Name(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier0RoutingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier0ServicesHighAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier1Name(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeAdminPassword(formats strfmt.Registry) error {

	if err := validate.Required("edgeAdminPassword", "body", m.EdgeAdminPassword); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeAuditPassword(formats strfmt.Registry) error {

	if err := validate.Required("edgeAuditPassword", "body", m.EdgeAuditPassword); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeClusterName(formats strfmt.Registry) error {

	if err := validate.Required("edgeClusterName", "body", m.EdgeClusterName); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeClusterProfileSpec(formats strfmt.Registry) error {

	if err := validate.Required("edgeClusterProfileSpec", "body", m.EdgeClusterProfileSpec); err != nil {
		return err
	}

	if m.EdgeClusterProfileSpec != nil {
		if err := m.EdgeClusterProfileSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeClusterProfileSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeClusterProfileSpec")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeClusterProfileType(formats strfmt.Registry) error {

	if err := validate.Required("edgeClusterProfileType", "body", m.EdgeClusterProfileType); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeClusterType(formats strfmt.Registry) error {

	if err := validate.Required("edgeClusterType", "body", m.EdgeClusterType); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeFormFactor(formats strfmt.Registry) error {

	if err := validate.Required("edgeFormFactor", "body", m.EdgeFormFactor); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeNodeSpecs(formats strfmt.Registry) error {

	if err := validate.Required("edgeNodeSpecs", "body", m.EdgeNodeSpecs); err != nil {
		return err
	}

	for i := 0; i < len(m.EdgeNodeSpecs); i++ {
		if swag.IsZero(m.EdgeNodeSpecs[i]) { // not required
			continue
		}

		if m.EdgeNodeSpecs[i] != nil {
			if err := m.EdgeNodeSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateEdgeRootPassword(formats strfmt.Registry) error {

	if err := validate.Required("edgeRootPassword", "body", m.EdgeRootPassword); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateMtu(formats strfmt.Registry) error {

	if err := validate.Required("mtu", "body", m.Mtu); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateTier0Name(formats strfmt.Registry) error {

	if err := validate.Required("tier0Name", "body", m.Tier0Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateTier0RoutingType(formats strfmt.Registry) error {

	if err := validate.Required("tier0RoutingType", "body", m.Tier0RoutingType); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateTier0ServicesHighAvailability(formats strfmt.Registry) error {

	if err := validate.Required("tier0ServicesHighAvailability", "body", m.Tier0ServicesHighAvailability); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterCreationSpec) validateTier1Name(formats strfmt.Registry) error {

	if err := validate.Required("tier1Name", "body", m.Tier1Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this edge cluster creation spec based on the context it is used
func (m *EdgeClusterCreationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeClusterProfileSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeNodeSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterCreationSpec) contextValidateEdgeClusterProfileSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeClusterProfileSpec != nil {
		if err := m.EdgeClusterProfileSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeClusterProfileSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeClusterProfileSpec")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeClusterCreationSpec) contextValidateEdgeNodeSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EdgeNodeSpecs); i++ {

		if m.EdgeNodeSpecs[i] != nil {
			if err := m.EdgeNodeSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterCreationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterCreationSpec) UnmarshalBinary(b []byte) error {
	var res EdgeClusterCreationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
