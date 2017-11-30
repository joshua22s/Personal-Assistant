package main

import (
	hub "github.com/joshua22s/Personal-Assistant/Hub"
	web "github.com/joshua22s/Personal-Assistant/Webserver"
)

func main() {
	hub.SetupWakeUpTimeCalculator()
	hub.InitializeDevices()
	web.StartWebServer()
}
