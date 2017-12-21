package models


func NewDeviceType(id int, name string) DeviceType {
	return DeviceType{id, name}
}

type DeviceType struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
