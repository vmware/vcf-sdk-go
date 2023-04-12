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

// Certificate Represents certificate and its attributes
//
// swagger:model Certificate
type Certificate struct {

	// Certificate chain ordered from intermediate to root certificates
	// Required: true
	CaChain []*Certificate `json:"caChain"`

	// Domain of the resource certificate
	// Required: true
	Domain *string `json:"domain"`

	// Certificate expiry status
	// Example: One among: ACTIVE, ABOUT_TO_EXPIRE, EXPIRED
	// Required: true
	ExpirationStatus *string `json:"expirationStatus"`

	// Error if certificate cannot be fetched
	// Required: true
	GetCertificateError *string `json:"getCertificateError"`

	// Whether the certificate is installed or not
	// Required: true
	IsInstalled *bool `json:"isInstalled"`

	// The certificate authority that issued the certificate
	// Required: true
	IssuedBy *string `json:"issuedBy"`

	// To whom the certificate is issued
	// Required: true
	IssuedTo *string `json:"issuedTo"`

	// The keysize of the certificate
	// Example: One among: 2048, 3072, 4096
	// Required: true
	KeySize *string `json:"keySize"`

	// The timestamp after which certificate is not valid
	// Required: true
	NotAfter *string `json:"notAfter"`

	// The timestamp before which certificate is not valid
	// Required: true
	NotBefore *string `json:"notBefore"`

	// Number of days left for the certificate to expire
	// Required: true
	NumberOfDaysToExpire *int32 `json:"numberOfDaysToExpire"`

	// The PEM encoded certificate content
	// Required: true
	PemEncoded *string `json:"pemEncoded"`

	// The public key of the certificate
	// Required: true
	PublicKey *string `json:"publicKey"`

	// The public key algorithm of the certificate
	// Example: One among: RSA
	// Required: true
	PublicKeyAlgorithm *string `json:"publicKeyAlgorithm"`

	// The serial number of the certificate
	// Required: true
	SerialNumber *string `json:"serialNumber"`

	// Algorithm used to sign the certificate
	// Required: true
	SignatureAlgorithm *string `json:"signatureAlgorithm"`

	// Complete distinguished name to which the certificate is issued
	// Required: true
	Subject *string `json:"subject"`

	// The alternative names to which the certificate is issued
	// Required: true
	SubjectAlternativeName []string `json:"subjectAlternativeName"`

	// Thumbprint generated using certificate content
	// Required: true
	Thumbprint *string `json:"thumbprint"`

	// Algorithm used to generate thumbprint
	// Required: true
	ThumbprintAlgorithm *string `json:"thumbprintAlgorithm"`

	// The X.509 version of the certificate
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this certificate
func (m *Certificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGetCertificateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsInstalled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeySize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfDaysToExpire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePemEncoded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKeyAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatureAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectAlternativeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbprintAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Certificate) validateCaChain(formats strfmt.Registry) error {

	if err := validate.Required("caChain", "body", m.CaChain); err != nil {
		return err
	}

	for i := 0; i < len(m.CaChain); i++ {
		if swag.IsZero(m.CaChain[i]) { // not required
			continue
		}

		if m.CaChain[i] != nil {
			if err := m.CaChain[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("caChain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("caChain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateExpirationStatus(formats strfmt.Registry) error {

	if err := validate.Required("expirationStatus", "body", m.ExpirationStatus); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateGetCertificateError(formats strfmt.Registry) error {

	if err := validate.Required("getCertificateError", "body", m.GetCertificateError); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateIsInstalled(formats strfmt.Registry) error {

	if err := validate.Required("isInstalled", "body", m.IsInstalled); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateIssuedBy(formats strfmt.Registry) error {

	if err := validate.Required("issuedBy", "body", m.IssuedBy); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateIssuedTo(formats strfmt.Registry) error {

	if err := validate.Required("issuedTo", "body", m.IssuedTo); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateKeySize(formats strfmt.Registry) error {

	if err := validate.Required("keySize", "body", m.KeySize); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateNotAfter(formats strfmt.Registry) error {

	if err := validate.Required("notAfter", "body", m.NotAfter); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateNotBefore(formats strfmt.Registry) error {

	if err := validate.Required("notBefore", "body", m.NotBefore); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateNumberOfDaysToExpire(formats strfmt.Registry) error {

	if err := validate.Required("numberOfDaysToExpire", "body", m.NumberOfDaysToExpire); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validatePemEncoded(formats strfmt.Registry) error {

	if err := validate.Required("pemEncoded", "body", m.PemEncoded); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validatePublicKeyAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("publicKeyAlgorithm", "body", m.PublicKeyAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.Required("serialNumber", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateSignatureAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("signatureAlgorithm", "body", m.SignatureAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateSubjectAlternativeName(formats strfmt.Registry) error {

	if err := validate.Required("subjectAlternativeName", "body", m.SubjectAlternativeName); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateThumbprint(formats strfmt.Registry) error {

	if err := validate.Required("thumbprint", "body", m.Thumbprint); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateThumbprintAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("thumbprintAlgorithm", "body", m.ThumbprintAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this certificate based on the context it is used
func (m *Certificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCaChain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Certificate) contextValidateCaChain(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CaChain); i++ {

		if m.CaChain[i] != nil {
			if err := m.CaChain[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("caChain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("caChain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Certificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Certificate) UnmarshalBinary(b []byte) error {
	var res Certificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
