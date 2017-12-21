package models

type ILightingDevice interface {
	GetId() int
	GetName() string
	Setup()
	TurnOn()
}
