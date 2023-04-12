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

// VsanRemoteDatastoreClusterSpec vSAN remote datastore configuration for the cluster
//
// swagger:model VsanRemoteDatastoreClusterSpec
type VsanRemoteDatastoreClusterSpec struct {

	// List of Remote vSAN datastore configuration for HCI Mesh compute client cluster
	VsanRemoteDatastoreSpec []*VsanRemoteDatastoreSpec `json:"vsanRemoteDatastoreSpec"`
}

// Validate validates this vsan remote datastore cluster spec
func (m *VsanRemoteDatastoreClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVsanRemoteDatastoreSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsanRemoteDatastoreClusterSpec) validateVsanRemoteDatastoreSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.VsanRemoteDatastoreSpec) { // not required
		return nil
	}

	for i := 0; i < len(m.VsanRemoteDatastoreSpec); i++ {
		if swag.IsZero(m.VsanRemoteDatastoreSpec[i]) { // not required
			continue
		}

		if m.VsanRemoteDatastoreSpec[i] != nil {
			if err := m.VsanRemoteDatastoreSpec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vsanRemoteDatastoreSpec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vsanRemoteDatastoreSpec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vsan remote datastore cluster spec based on the context it is used
func (m *VsanRemoteDatastoreClusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVsanRemoteDatastoreSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsanRemoteDatastoreClusterSpec) contextValidateVsanRemoteDatastoreSpec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VsanRemoteDatastoreSpec); i++ {

		if m.VsanRemoteDatastoreSpec[i] != nil {
			if err := m.VsanRemoteDatastoreSpec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vsanRemoteDatastoreSpec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vsanRemoteDatastoreSpec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VsanRemoteDatastoreClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsanRemoteDatastoreClusterSpec) UnmarshalBinary(b []byte) error {
	var res VsanRemoteDatastoreClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
