// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VSphereTag VSphereTag represents the tags that are attached or created on the cluster level, that are then propagated down to the
// MachineDeployments. In order to attach tags on MachineDeployment, users must create the tag on a cluster level first
// then attach that tag on the MachineDeployment.
//
// swagger:model VSphereTag
type VSphereTag struct {

	// CategoryID is the id of the vsphere category that the tag belongs to. If the category id is left empty, the default
	// category id for the cluster will be used.
	CategoryID string `json:"categoryID,omitempty"`

	// Tags represents the name of the created tags.
	Tags []string `json:"tags"`
}

// Validate validates this v sphere tag
func (m *VSphereTag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v sphere tag based on context it is used
func (m *VSphereTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VSphereTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VSphereTag) UnmarshalBinary(b []byte) error {
	var res VSphereTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}