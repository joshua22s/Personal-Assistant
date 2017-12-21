package main

import (
	api "github.com/joshua22s/hub/api"
)

func main() {
	apicontroller := api.NewApiController()
	apicontroller.StartApi(8000)
}
