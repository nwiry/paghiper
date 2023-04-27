package pix

import "time"

// PixResponse struct represents the response received from the Pix API.
type PixResponse struct {
	// Result of the Pix transaction, either "reject" or "success".
	Result string `json:"result"`

	// Message regarding the Pix transaction, returned by the Pix API.
	ResponseMessage string `json:"response_message"`

	// Unique ID of the Pix transaction, generated by the Pix API.
	TransactionID string `json:"transaction_id"`

	// Date and time the Pix transaction was created.
	CreatedDate time.Time `json:"created_date"`

	// Value of the Pix transaction in cents.
	ValueCents int32 `json:"value_cents"`

	// Status of the Pix transaction. Possible values:
	// "pending" (waiting for payment), "canceled", "completed", "paid" (approved),
	// "processing" (under analysis), or "refunded".
	Status string `json:"status"`

	// ID of the order the Pix transaction is associated with.
	OrderID string `json:"order_id"`

	// Due date of the Pix transaction.
	DueDate time.Time `json:"due_date"`

	// Struct containing information about the Pix code associated with the transaction.
	PixCode struct {
		// Base64-encoded QR code image of the Pix transaction.
		QrcodeBase64 string `json:"qrcode_base64"`

		// URL for viewing the Pix transaction QR code image.
		QrcodeImageURL string `json:"qrcode_image_url"`

		// QRCode (Pix Copia e Cola)
		Emv string `json:"emv"`

		// URL for viewing the Pix transaction.
		PixURL string `json:"pix_url"`

		// Url Bacen Pix
		BacenURL string `json:"bacen_url"`
	} `json:"pix_code"`

	// HTTP status code returned by the Pix API.
	HTTPCode int `json:"http_code"`
}