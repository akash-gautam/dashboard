// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EKSClusterRole EKSClusterRole represents a EKS Cluster Service Role.
//
// swagger:model EKSClusterRole
type EKSClusterRole struct {

	// The Amazon Resource Name (ARN) specifying the role. For more information
	// about ARNs and how to use them in policies, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide guide.
	Arn string `json:"arn,omitempty"`

	// RoleName  represents the friendly name that identifies the role.
	RoleName string `json:"roleName,omitempty"`
}

// Validate validates this e k s cluster role
func (m *EKSClusterRole) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this e k s cluster role based on context it is used
func (m *EKSClusterRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EKSClusterRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EKSClusterRole) UnmarshalBinary(b []byte) error {
	var res EKSClusterRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}