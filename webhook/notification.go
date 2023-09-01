package webhook

import (
	"github.com/nwiry/paghiper"
	"github.com/nwiry/paghiper/exception"
	"github.com/nwiry/paghiper/request"
)

type PaghiperNotification struct {
	paghiper.PaghiperRequest
	TransactionID  string `json:"transaction_id"`
	NotificationID string `json:"notification_id"`
}

func (n *PaghiperNotification) GetStatus(endpoint string, payment interface{}) (map[string]interface{}, *exception.PaghiperError, error) {

	res, code, err := request.Make(endpoint, "notification/", n)
	if err != nil {
		return nil, nil, err
	}

	// Validate the response and return a PaghiperError object if the response is not valid.
	if err := request.ValidateRequestResponse(res, code); err != nil {
		return nil, &exception.PaghiperError{ResponseMessage: err.Error(), HTTPCode: code}, nil
	}

	return res, nil, nil
}
