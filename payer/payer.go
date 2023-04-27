package payer

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/paemuri/brdoc"
)

// Payer represents a payer object with the necessary information for payment.
type Payer struct {
	// CpfCnpj represents the payer's individual taxpayer identification number (CPF) or
	// legal entity identifier (CNPJ).
	CpfCnpj string `json:"payer_cpf_cnpj"`

	// Email is the email address of the payer.
	Email string `json:"payer_email"`

	// Name is the name of the payer.
	Name string `json:"payer_name"`

	// Phone is the phone number of the payer.
	Phone int32 `json:"payer_phone,omitempty"`

	// City is the city of the payer.
	City string `json:"payer_city,omitempty"`

	// Complement is the complement of the payer's address.
	Complement string `json:"payer_complement,omitempty"`

	// District is the district of the payer's address.
	District string `json:"payer_district,omitempty"`

	// Number is the number of the payer's address.
	Number int32 `json:"payer_number,omitempty"`

	// Street is the street of the payer's address.
	Street string `json:"payer_street,omitempty"`

	// State is the state of the payer's address.
	State string `json:"payer_state,omitempty"`

	// ZipCode is the zip code of the payer's address.
	ZipCode int32 `json:"payer_zip_code,omitempty"`
}

// ValidatePayer validates a Payer struct, returning an error if any required fields are missing.
func (p *Payer) ValidatePayer() error {
	// Check if Email is set
	if strings.TrimSpace(p.Email) == "" {
		return errors.New("Email is required")
	}

	if _, err := mail.ParseAddress(p.Email); err != nil {
		return errors.New("Email is invalid")
	}

	// Check if Name is set
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("Name is required")
	}

	// Check if CpfCnpj is set and is valid
	if strings.TrimSpace(p.CpfCnpj) == "" {
		return errors.New("CpfCnpj is required")
	}

	if !brdoc.IsCPF(p.CpfCnpj) && !brdoc.IsCNPJ(p.CpfCnpj) {
		return errors.New("CpfCnpj is invalid")
	}

	return nil
}
