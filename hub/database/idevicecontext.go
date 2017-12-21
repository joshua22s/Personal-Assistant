package database

import (
	models "github.com/joshua22s/hub/models"
)

type IDeviceContext interface {
	createDevice(device models.Device) bool
	getDevices() []models.Device
	getDeviceTypes() []models.DeviceType
}