// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DependencyCredentials dependency credentials
//
// swagger:model DependencyCredentials
type DependencyCredentials struct {

	// helm credentials
	HelmCredentials *HelmCredentials `json:"helmCredentials,omitempty"`
}

// Validate validates this dependency credentials
func (m *DependencyCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHelmCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DependencyCredentials) validateHelmCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.HelmCredentials) { // not required
		return nil
	}

	if m.HelmCredentials != nil {
		if err := m.HelmCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helmCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helmCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dependency credentials based on the context it is used
func (m *DependencyCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHelmCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DependencyCredentials) contextValidateHelmCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.HelmCredentials != nil {
		if err := m.HelmCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helmCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helmCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DependencyCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DependencyCredentials) UnmarshalBinary(b []byte) error {
	var res DependencyCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}