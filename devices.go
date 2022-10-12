package gokirimwa

import (
	"fmt"
	"net/http"
	"time"
)

func (c Client) RegisterDevice(deviceID string) (resp RegisterDeviceResponse, err error) {
	var respBody RegisterDeviceResponse

	reqBody := RegisterDeviceRequest{DeviceID: deviceID}
	if err = c.do(http.MethodPost, DEVICE, nil, reqBody, &respBody); err != nil {
		return
	}

	return respBody, nil
}

func (c Client) DeleteDevice(deviceID string) (err error) {
	if err = c.do(http.MethodDelete, fmt.Sprintf("%s/%s", DEVICE, deviceID), nil, nil, nil); err != nil {
		return err
	}

	return nil
}

type (
	RegisterDeviceRequest struct {
		DeviceID string `json:"device_id"`
	}

	RegisterDeviceResponse struct {
		ID        string    `json:"id"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		Meta      struct {
			Location string `json:"location"`
		} `json:"meta"`
	}
)
