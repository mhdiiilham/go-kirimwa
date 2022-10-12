package gokirimwa_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	gokirimwa "github.com/mhdiiilham/go-kirimwa"
	"github.com/mhdiiilham/go-kirimwa/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeviceTestSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	mockHTTP *mock.MockHTTPClient
	apiKey   string
	client   *gokirimwa.Client
}

func TestDevices(t *testing.T) {
	suite.Run(t, new(DeviceTestSuite))
}

func (suite *DeviceTestSuite) TestRegisterNewDevice() {
	testCases := []struct {
		name         string
		deviceID     string
		doMocks      func()
		expectedResp gokirimwa.RegisterDeviceResponse
		expectedErr  error
	}{
		{
			name:     "success",
			deviceID: "fake_device_id",
			doMocks: func() {
				var bResp bytes.Buffer
				json.NewEncoder(&bResp).Encode(gokirimwa.RegisterDeviceResponse{
					ID:        "fake_device_id",
					Status:    "disconnect",
					CreatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
				})

				suite.
					mockHTTP.
					EXPECT().
					Do(gomock.Any()).
					Return(&http.Response{
						StatusCode: http.StatusCreated,
						Body:       ioutil.NopCloser(&bResp),
					}, nil).
					Times(1)
			},
			expectedResp: gokirimwa.RegisterDeviceResponse{
				ID:        "fake_device_id",
				Status:    "disconnect",
				CreatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expectedErr: nil,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assertion := assert.New(t)
			tt.doMocks()

			actual, actualErr := suite.client.RegisterDevice(tt.deviceID)
			assertion.Equal(tt.expectedResp, actual)
			assertion.Equal(tt.expectedErr, actualErr)
		})
	}
}

func (suite *DeviceTestSuite) TestDeleteDevice() {
	testCases := []struct {
		name        string
		deviceID    string
		doMocks     func()
		expectedErr error
	}{
		{
			name:     "return nil",
			deviceID: "fake_device_id",
			doMocks: func() {
				req, _ := http.NewRequest(http.MethodDelete, "https://api.kirimwa.id/v1/devices/fake_device_id", nil)
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", "Bearer "+suite.apiKey)
				suite.
					mockHTTP.
					EXPECT().
					Do(req).
					Return(&http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(nil)),
					}, nil).
					Times(1)
			},
			expectedErr: nil,
		},
		{
			name:     "device not found",
			deviceID: "fake_device_id",
			doMocks: func() {
				req, _ := http.NewRequest(http.MethodDelete, "https://api.kirimwa.id/v1/devices/fake_device_id", nil)
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", "Bearer "+suite.apiKey)
				suite.
					mockHTTP.
					EXPECT().
					Do(req).
					Return(&http.Response{
						StatusCode: http.StatusNotFound,
						Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"message": "Error: Device not found."}`))),
					}, nil).
					Times(1)
			},
			expectedErr: gokirimwa.ErrorResponse{Message: "Error: Device not found."},
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.doMocks()
			actual := suite.client.DeleteDevice(tt.deviceID)
			assert.Equal(t, tt.expectedErr, actual)
		})
	}
}

func (suite *DeviceTestSuite) SetupTest() {
	t := suite.T()
	suite.ctrl = gomock.NewController(t)
	suite.mockHTTP = mock.NewMockHTTPClient(suite.ctrl)
	suite.apiKey = "fake_api_key"
	suite.client = gokirimwa.NewKirimWAMockWithClient(suite.apiKey, suite.mockHTTP)
}

func (suite *DeviceTestSuite) TearDownTest() {
	defer suite.ctrl.Finish()
}
