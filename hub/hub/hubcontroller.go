package hub

import (
	"fmt"
	"github.com/joshua22s/hub/database"
	models "github.com/joshua22s/hub/models"
)

func NewHubController() HubController {
	controller := HubController{database.NewRepository()}
	return controller
}

type HubController struct {
	repo database.Repository
}

func (this *HubController) Login(username string, password string) {
	if (this.repo.Login(username, password)) {
		fmt.Println("Login succesfull")
	} else {
		fmt.Println("Login unsuccesfull")
	}
}

func (this *HubController) Register(username string, password string, long float64, lat float64) {
	if (this.repo.Register(username, password, long, lat)) {
		fmt.Println("Register succesfull")
	} else {
		fmt.Println("Register unsuccesfull")
	}
}

func (this *HubController) CreateDevice(device models.Device) bool {
	return this.repo.CreateDevice(device)
}

func (this *HubController) GetDeviceTypes() []models.DeviceType {
	return this.repo.GetDeviceTypes()
}

func (this *HubController) GetDevices() []models.Device {
	return this.repo.GetDevices()
}