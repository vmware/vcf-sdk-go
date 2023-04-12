// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NsxTClusterReference Represents an NSX-T Cluster reference
//
// swagger:model NsxTClusterReference
type NsxTClusterReference struct {

	// ID of the NSX-T cluster
	ID string `json:"id,omitempty"`

	// VIP (Virtual IP Address) of the NSX-T cluster
	Vip string `json:"vip,omitempty"`

	// FQDN for VIP of the NSX-T cluster
	VipFqdn string `json:"vipFqdn,omitempty"`
}

// Validate validates this nsx t cluster reference
func (m *NsxTClusterReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nsx t cluster reference based on context it is used
func (m *NsxTClusterReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NsxTClusterReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxTClusterReference) UnmarshalBinary(b []byte) error {
	var res NsxTClusterReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}