// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// NSXTResources Nsxt Upgrade Resources for an NSXT instance
//
// swagger:model NsxtResources
type NSXTResources struct {

	// Bundle id used to compute the upgradability
	BundleID string `json:"bundleId,omitempty"`

	// Domain id associated with the NSX-T cluster
	DomainID string `json:"domainId,omitempty"`

	// List of edge clusters that are candidates for upgrade
	NSXTEdgeClusters []*NSXTEdgeCluster `json:"nsxtEdgeClusters"`

	// List of host clusters that are candidates for upgrade
	NSXTHostClusters []*NSXTHostCluster `json:"nsxtHostClusters"`

	// Manager cluster that is a candidate for upgrade
	NSXTManagerCluster *NSXTManagerCluster `json:"nsxtManagerCluster,omitempty"`

	// Upgrade coordinator pertaining to the NSX-T instance
	NSXTUpgradeCoordinator *NSXTUpgradeCoordinator `json:"nsxtUpgradeCoordinator,omitempty"`
}

// Validate validates this Nsxt resources
func (m *NSXTResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNSXTEdgeClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNSXTHostClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNSXTManagerCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNSXTUpgradeCoordinator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTResources) validateNSXTEdgeClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.NSXTEdgeClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.NSXTEdgeClusters); i++ {
		if swag.IsZero(m.NSXTEdgeClusters[i]) { // not required
			continue
		}

		if m.NSXTEdgeClusters[i] != nil {
			if err := m.NSXTEdgeClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtEdgeClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtEdgeClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTResources) validateNSXTHostClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.NSXTHostClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.NSXTHostClusters); i++ {
		if swag.IsZero(m.NSXTHostClusters[i]) { // not required
			continue
		}

		if m.NSXTHostClusters[i] != nil {
			if err := m.NSXTHostClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtHostClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtHostClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTResources) validateNSXTManagerCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.NSXTManagerCluster) { // not required
		return nil
	}

	if m.NSXTManagerCluster != nil {
		if err := m.NSXTManagerCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtManagerCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtManagerCluster")
			}
			return err
		}
	}

	return nil
}

func (m *NSXTResources) validateNSXTUpgradeCoordinator(formats strfmt.Registry) error {
	if swag.IsZero(m.NSXTUpgradeCoordinator) { // not required
		return nil
	}

	if m.NSXTUpgradeCoordinator != nil {
		if err := m.NSXTUpgradeCoordinator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtUpgradeCoordinator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtUpgradeCoordinator")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Nsxt resources based on the context it is used
func (m *NSXTResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNSXTEdgeClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNSXTHostClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNSXTManagerCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNSXTUpgradeCoordinator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTResources) contextValidateNSXTEdgeClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NSXTEdgeClusters); i++ {

		if m.NSXTEdgeClusters[i] != nil {
			if err := m.NSXTEdgeClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtEdgeClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtEdgeClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTResources) contextValidateNSXTHostClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NSXTHostClusters); i++ {

		if m.NSXTHostClusters[i] != nil {
			if err := m.NSXTHostClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtHostClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtHostClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTResources) contextValidateNSXTManagerCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.NSXTManagerCluster != nil {
		if err := m.NSXTManagerCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtManagerCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtManagerCluster")
			}
			return err
		}
	}

	return nil
}

func (m *NSXTResources) contextValidateNSXTUpgradeCoordinator(ctx context.Context, formats strfmt.Registry) error {

	if m.NSXTUpgradeCoordinator != nil {
		if err := m.NSXTUpgradeCoordinator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtUpgradeCoordinator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtUpgradeCoordinator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NSXTResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTResources) UnmarshalBinary(b []byte) error {
	var res NSXTResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
