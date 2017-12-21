package models

func NewDevice(name string, deviceType DeviceType) Device {
	return Device{name, deviceType}
} 

type Device struct {
	Name string`json:"name"`
	Type DeviceType `json: "type"`
}