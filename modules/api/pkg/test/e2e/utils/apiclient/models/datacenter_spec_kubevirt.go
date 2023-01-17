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

// DatacenterSpecKubevirt DatacenterSpecKubevirt describes a kubevirt datacenter.
//
// swagger:model DatacenterSpecKubevirt
type DatacenterSpecKubevirt struct {

	// CustomNetworkPolicies (optional) allows to add some extra custom NetworkPolicies, that are deployed
	// in the dedicated infra KubeVirt cluster. They are added to the defaults.
	CustomNetworkPolicies []*CustomNetworkPolicy `json:"customNetworkPolicies"`

	// DNSPolicy represents the dns policy for the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst',
	// 'Default' or 'None'. Defaults to "ClusterFirst". DNS parameters given in DNSConfig will be merged with the
	// policy selected with DNSPolicy.
	DNSPolicy string `json:"dnsPolicy,omitempty"`

	// dns config
	DNSConfig *PodDNSConfig `json:"dnsConfig,omitempty"`

	// images
	Images *ImageSources `json:"images,omitempty"`
}

// Validate validates this datacenter spec kubevirt
func (m *DatacenterSpecKubevirt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomNetworkPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecKubevirt) validateCustomNetworkPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomNetworkPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomNetworkPolicies); i++ {
		if swag.IsZero(m.CustomNetworkPolicies[i]) { // not required
			continue
		}

		if m.CustomNetworkPolicies[i] != nil {
			if err := m.CustomNetworkPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customNetworkPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customNetworkPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterSpecKubevirt) validateDNSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSConfig) { // not required
		return nil
	}

	if m.DNSConfig != nil {
		if err := m.DNSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpecKubevirt) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if m.Images != nil {
		if err := m.Images.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec kubevirt based on the context it is used
func (m *DatacenterSpecKubevirt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomNetworkPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecKubevirt) contextValidateCustomNetworkPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomNetworkPolicies); i++ {

		if m.CustomNetworkPolicies[i] != nil {
			if err := m.CustomNetworkPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customNetworkPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customNetworkPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterSpecKubevirt) contextValidateDNSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DNSConfig != nil {
		if err := m.DNSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpecKubevirt) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	if m.Images != nil {
		if err := m.Images.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecKubevirt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecKubevirt) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecKubevirt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}