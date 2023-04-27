package paghiper

import (
	"errors"
	"strings"

	"github.com/nwiry/paghiper/payer"
)

// BOLETO_ENDPOINT is the endpoint URL for boleto requests.
const BOLETO_ENDPOINT string = "https://api.paghiper.com/invoice/"

type PaghiperRequest struct {
	// ApiKey is used to identify the seller.
	ApiKey string `json:"apiKey"`

	// Token is used to identify the seller's transaction.
	Token string `json:"token,omitempty"`
}

// PaghiperPaymentRequest represents a payment request to be sent to the PagHiper payment gateway.
// It contains all necessary information for the payment, such as the seller's API key,
// due date, discount amount, and order details.
type PaghiperPaymentRequest struct {
	PaghiperRequest

	// DaysDueDate sets the number of calendar days until expiration.
	DaysDueDate int32 `json:"days_due_date"`

	// DiscountCents sets the total purchase discount amount in cents.
	DiscountCents int32 `json:"discount_cents,omitempty"`

	// FixedDescription defines whether the pre-configured phrase on the PagHiper panel will be displayed.
	FixedDescription bool `json:"fixed_description,omitempty"`

	// NotificationURL defines the address of the page where PagHiper
	// will send the POST with transaction information.
	NotificationURL string `json:"notification_url"`

	// NumberNtfiscal will display the invoice number in the transaction request.
	NumberNtfiscal int32 `json:"number_ntfiscal,omitempty"`

	// OrderID defines a code to reference the payment.
	OrderID string `json:"order_id"`

	// SellerDescription defines a text that will vary according to each specific transaction,
	// and may include information that refers to the request/service purchased by the customer.
	SellerDescription string `json:"seller_description,omitempty"`

	// ShippingPriceCents sets the total shipping amount in cents.
	ShippingPriceCents int32 `json:"shipping_price_cents,omitempty"`

	// ShippingMethods defines the order delivery method.
	ShippingMethods string `json:"shipping_methods,omitempty"`
}

// ValidatePaymentRequest checks if all required fields in the payment request are filled.
// Returns an error if any required field is missing.
func (r *PaghiperPaymentRequest) ValidatePaymentRequest() error {
	// Check if ApiKey is set
	if strings.TrimSpace(r.ApiKey) == "" {
		return errors.New("ApiKey is required")
	}

	// Check if DaysDueDate is set
	if r.DaysDueDate <= 0 {
		return errors.New("DaysDueDate must be greater than zero")
	}

	// Check if NotificationURL is set
	if strings.TrimSpace(r.NotificationURL) == "" {
		return errors.New("NotificationURL is required")
	}

	// Check if OrderID is set
	if strings.TrimSpace(r.OrderID) == "" {
		return errors.New("OrderID is required")
	}

	return nil
}

// NewPayer Ensures the import of the brdoc package by creating a new instance of payer
func NewPayer() *payer.Payer {
	p := &payer.Payer{}
	p.ValidatePayer()
	return p
}
