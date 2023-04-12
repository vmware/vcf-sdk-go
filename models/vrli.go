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

// Vrli Spec contains parameters of a vRealize Log Insight instance
//
// swagger:model Vrli
type Vrli struct {

	// The ID of the vRealize Log Insight instance
	ID string `json:"id,omitempty"`

	// The Fully Qualified Domain Name of the vRealize Log Insight load balancer
	// Example: load-balancer.vrack.vsphere.local
	LoadBalancerFqdn string `json:"loadBalancerFqdn,omitempty"`

	// The IP address of the vRealize Log Insight load balancer
	// Example: 10.0.0.15
	LoadBalancerIPAddress string `json:"loadBalancerIpAddress,omitempty"`

	// The nodes of the vRealize Log Insight instance
	Nodes []*VrealizeProductNode `json:"nodes"`

	// The status of the vRealize Log Insight instance
	// Example: ACTIVE, ERROR
	Status string `json:"status,omitempty"`

	// The version of the vRealize Log Insight instance
	// Example: 4.8.0-13036238
	Version string `json:"version,omitempty"`
}

// Validate validates this vrli
func (m *Vrli) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vrli) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vrli based on the context it is used
func (m *Vrli) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vrli) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vrli) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vrli) UnmarshalBinary(b []byte) error {
	var res Vrli
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
