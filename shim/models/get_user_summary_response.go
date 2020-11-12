// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetUserSummaryResponse get user summary response
//
// swagger:model getUserSummaryResponse
type GetUserSummaryResponse struct {

	// The audited Organizations
	AuditedOrganizations []GenericObject `json:"audited_organizations"`

	// The audited Spaces
	AuditedSpaces []GenericObject `json:"audited_spaces"`

	// The billing Managed Organizations
	BillingManagedOrganizations []GenericObject `json:"billing_managed_organizations"`

	// The managed Organizations
	ManagedOrganizations []GenericObject `json:"managed_organizations"`

	// The managed Spaces
	ManagedSpaces []GenericObject `json:"managed_spaces"`

	// The organizations
	Organizations []GenericObject `json:"organizations"`

	// The spaces
	Spaces []GenericObject `json:"spaces"`
}

// Validate validates this get user summary response
func (m *GetUserSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditedOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditedSpaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingManagedOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedSpaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUserSummaryResponse) validateAuditedOrganizations(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditedOrganizations) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditedOrganizations); i++ {

		if err := m.AuditedOrganizations[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audited_organizations" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GetUserSummaryResponse) validateAuditedSpaces(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditedSpaces) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditedSpaces); i++ {

		if err := m.AuditedSpaces[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audited_spaces" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GetUserSummaryResponse) validateBillingManagedOrganizations(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingManagedOrganizations) { // not required
		return nil
	}

	for i := 0; i < len(m.BillingManagedOrganizations); i++ {

		if err := m.BillingManagedOrganizations[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_managed_organizations" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GetUserSummaryResponse) validateManagedOrganizations(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedOrganizations) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagedOrganizations); i++ {

		if err := m.ManagedOrganizations[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_organizations" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GetUserSummaryResponse) validateManagedSpaces(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedSpaces) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagedSpaces); i++ {

		if err := m.ManagedSpaces[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_spaces" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GetUserSummaryResponse) validateOrganizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organizations); i++ {

		if err := m.Organizations[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GetUserSummaryResponse) validateSpaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Spaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Spaces); i++ {

		if err := m.Spaces[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spaces" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUserSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserSummaryResponse) UnmarshalBinary(b []byte) error {
	var res GetUserSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}