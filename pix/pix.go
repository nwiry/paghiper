package pix

import (
	"encoding/json"
	"errors"

	"github.com/nwiry/paghiper/items"
	"github.com/nwiry/paghiper/payer"
)

// PixRequest represents a request for a Pix payment.
type PixRequest struct {
	// Payer contains information about the person making the payment.
	payer.Payer
	// Items contains a list of items included in the payment request.
	Items []items.Items `json:"items,omitempty"`
}

// validatePixRequest checks if a PixRequest contains at least one item.
// Returns an error if the items list is empty, otherwise returns nil.
func (r *PixRequest) ValidatePixRequest() error {
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
