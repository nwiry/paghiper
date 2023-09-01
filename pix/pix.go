package pix

import (
	"encoding/json"
	"errors"

	"github.com/nwiry/paghiper"
	"github.com/nwiry/paghiper/exception"
	"github.com/nwiry/paghiper/items"
	"github.com/nwiry/paghiper/payer"
	"github.com/nwiry/paghiper/request"
)

// PIX_ENDPOINT is the endpoint URL for pix requests.
const PIX_ENDPOINT string = "https://pix.paghiper.com/invoice/"

// PixRequest represents a request for a Pix payment.
type PixRequest struct {
	// PaghiperPaymentRequest contains information about seller and transaction
	paghiper.PaghiperPaymentRequest
	// Payer contains information about the person making the payment.
	payer.Payer
	// Items contains a list of items included in the payment request.
	Items []items.Items `json:"items,omitempty"`
}

// validatePixRequest checks if a PixRequest contains at least one item.
// Returns an error if the items list is empty, otherwise returns nil.
func (r *PixRequest) ValidateCreatePixRequest() error {
	// Check if the length of the items list is zero
	if len(r.Items) == 0 {
		// If it is, return an error with a message
		return errors.New("at least one item is required")
	}
	// If the items list is not empty, return nil to indicate success
	return nil
}

func (p *PixRequest) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// Create creates a new PIX payment request in Paghiper using the given `PaghiperRequest` and `PixRequest` objects.
// It returns a `PixResponse` object and a `PaghiperError` object in case of an error, or `nil` values for both in case of success.
func (p *PixRequest) Create() (*PixResponse, *exception.PaghiperError, error) {
	// Validate the payment request and the PIX request using the validation functions defined in the respective structs.
	for _, f := range []func() error{
		p.ValidatePaymentRequest,
		p.ValidatePayer,
		p.ValidateCreatePixRequest,
	} {
		if err := f(); err != nil {
			return nil, nil, err
		}
	}

	// Create a new map to hold the JSON data of the payment request and PIX request.
	m := make(map[string]interface{})

	// Convert the payment request and PIX request objects to JSON and add them to the map.
	b, err := p.ToJSON()
	if err != nil {
		return nil, nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, nil, err
	}

	// Make a request to create a new PIX payment using the data in the map.
	res, code, err := request.Make(PIX_ENDPOINT, "create/", m)
	if err != nil {
		return nil, nil, err
	}

	// Validate the response and return a PaghiperError object if the response is not valid.
	if err := request.ValidateRequestResponse(res, code); err != nil {
		return nil, &exception.PaghiperError{ResponseMessage: err.Error(), HTTPCode: code}, nil
	}

	// Create a new PixResponse object using the data in the response and return it.
	pixRes := &PixResponse{}
	if err := request.SetResponse(res, pixRes); err != nil {
		return nil, nil, err
	}

	return pixRes, nil, nil
}
