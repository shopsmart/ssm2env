package testutils

import (
	"testing"

	"github.com/golang/mock/gomock"
)

//go:generate mockgen -source=../../pkg/service/service.go -destination=gen_service_mocks.go -package testutils SSMClient,Service

// MockController is a gomock controller
var MockController *gomock.Controller

// Setup initializes the mock controller
func Setup(t *testing.T) {
	MockController = gomock.NewController(t)
}

// Teardown cleans up after gomock
func Teardown() {
	MockController.Finish()
}
