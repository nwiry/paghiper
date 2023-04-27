package paghiper

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/nwiry/paghiper/exception"
	"github.com/nwiry/paghiper/pix"
)

// BOLETO_ENDPOINT is the endpoint URL for boleto requests.
const BOLETO_ENDPOINT string = "https://api.paghiper.com/invoice/"

// PIX_ENDPOINT is the endpoint URL for pix requests.
const PIX_ENDPOINT string = "https://pix.paghiper.com/invoice/"

// CREATE_PREFIX is the prefix for create requests.
const CREATE_PREFIX string = "create/"

// NOTIFICATION_PREFIX is the prefix for notification requests.
const NOTIFICATION_PREFIX string = "notification/"

// paymentRequest is an interface for a payment request.
type paymentRequest interface {
	// ToJSON converts the payment request to JSON.
	ToJSON() ([]byte, error)
}

// PaghiperRequest represents a payment request to be sent to the PagHiper payment gateway.
// It contains all necessary information for the payment, such as the seller's API key,
// due date, discount amount, and order details.
type PaghiperRequest struct {
	// ApiKey is used to identify the seller.
	ApiKey string `json:"apiKey"`

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
func (pr *PaghiperRequest) ValidatePaymentRequest() error {
	// Check if ApiKey is set
	if strings.TrimSpace(pr.ApiKey) == "" {
		return errors.New("ApiKey is required")
	}

	// Check if DaysDueDate is set
	if pr.DaysDueDate <= 0 {
		return errors.New("DaysDueDate must be greater than zero")
	}

	// Check if NotificationURL is set
	if strings.TrimSpace(pr.NotificationURL) == "" {
		return errors.New("NotificationURL is required")
	}

	// Check if OrderID is set
	if strings.TrimSpace(pr.OrderID) == "" {
		return errors.New("OrderID is required")
	}

	return nil
}

// ToJSON converts the payment request to JSON.
func (r *PaghiperRequest) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

// NewPixRequest creates a new PIX payment request in Paghiper using the given `PaghiperRequest` and `PixRequest` objects.
// It returns a `PixResponse` object and a `PaghiperError` object in case of an error, or `nil` values for both in case of success.
func NewPixRequest(r *PaghiperRequest, p *pix.PixRequest) (*pix.PixResponse, *exception.PaghiperError, error) {
	// Validate the payment request and the PIX request using the validation functions defined in the respective structs.
	for _, f := range []func() error{
		r.ValidatePaymentRequest,
		p.ValidatePixRequest,
		p.ValidatePayer,
	} {
		if err := f(); err != nil {
			return nil, nil, err
		}
	}

	// Create a new map to hold the JSON data of the payment request and PIX request.
	m := make(map[string]interface{})

	// Convert the payment request and PIX request objects to JSON and add them to the map.
	var j []paymentRequest
	j = append(j, r, p)
	for _, j := range j {
		b, err := j.ToJSON()
		if err != nil {
			return nil, nil, err
		}
		if err = json.Unmarshal(b, &m); err != nil {
			return nil, nil, err
		}
	}

	// Make a request to create a new PIX payment using the data in the map.
	res, code, err := makeRequest(PIX_ENDPOINT, CREATE_PREFIX, m)
	if err != nil {
		return nil, nil, err
	}

	// Validate the response and return a PaghiperError object if the response is not valid.
	if err := validateRequestResponse(res, code); err != nil {
		return nil, &exception.PaghiperError{ResponseMessage: err.Error(), HTTPCode: code}, nil
	}

	// Create a new PixResponse object using the data in the response and return it.
	pixRes := &pix.PixResponse{}
	if err := setResponse(res, pixRes); err != nil {
		return nil, nil, err
	}
	return pixRes, nil, nil
}
