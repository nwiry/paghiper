package payer

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/paemuri/brdoc"
)

// Payer represents a customer that will be paying for a service or product.
type Payer struct {
	// CpfCnpj represents the payer's individual taxpayer identification number (CPF) or
	// legal entity identifier (CNPJ).
	CpfCnpj string `json:"payer_cpf_cnpj"`

	// Email represents the payer's email address.
	Email string `json:"payer_email"`

	// Name represents the payer's full name.
	Name string `json:"payer_name"`

	// Phone represents the payer's phone number.
	Phone int32 `json:"payer_phone,omitempty"`
}

// ValidatePayer checks if all required fields in the payer are filled and if CPF/CNPJ is valid.
// Returns an error if any required field is missing or if CPF/CNPJ is invalid.
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
