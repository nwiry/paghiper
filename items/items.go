package items

// Items represents an item included in a payment request.
type Items struct {
	// Description contains a description of the item being sold.
	Description string `json:"description"`
	// Quantity contains the number of units of the item being sold.
	Quantity int32 `json:"quantity"`
	// ItemID contains an identifier for the item being sold.
	ItemID string `json:"item_id"`
	// PriceCents contains the price of the item in cents.
	PriceCents string `json:"price_cents"`
}