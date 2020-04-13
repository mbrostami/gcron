package logs_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_config "github.com/mbrostami/gcron/internal/config/mocks"
	"github.com/mbrostami/gcron/internal/logs"
)

func TestInitialize(t *testing.T) {
	ctrl := gomock.NewController(t)
	path := "../../tests/log"
	mockGeneralConfig := mock_config.NewMockGeneralConfig(ctrl)
	mockGeneralConfig.EXPECT().GetKey("log.enable").Return(true)
	mockGeneralConfig.EXPECT().GetKey("log.path").Return(path)
	mockGeneralConfig.EXPECT().GetLogLevel()
	file := logs.Initialize(mockGeneralConfig)
	if path != file.Name() {
		t.Errorf("Expected %s, got %s", path, file.Name())
	}
	err := file.Close()
	if err != nil {
		t.Errorf("Can not close file %s", err)
	}
}

func TestInitializeInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	path := "../../tests-invalid-dir/log"
	mockGeneralConfig := mock_config.NewMockGeneralConfig(ctrl)
	mockGeneralConfig.EXPECT().GetKey("log.enable").Return(true)
	mockGeneralConfig.EXPECT().GetKey("log.path").Return(path)
	mockGeneralConfig.EXPECT().GetLogLevel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	logs.Initialize(mockGeneralConfig)
}

func TestInitializeDisabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGeneralConfig := mock_config.NewMockGeneralConfig(ctrl)
	mockGeneralConfig.EXPECT().GetKey("log.enable").Return(false)
	mockGeneralConfig.EXPECT().GetLogLevel()
	file := logs.Initialize(mockGeneralConfig)
	err := file.Close()
	if err == nil {
		t.Error("file close should return error")
	}
}
