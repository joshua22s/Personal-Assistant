package main

type ILightingDevice interface {
	GetId() int
	GetName() string
	Setup()
}
