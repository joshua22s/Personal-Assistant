package database

import (
	"github.com/joshua22s/hub/models"
)

func NewRepository() Repository {
	conn := MySQLConnection{}
	conn.OpenConnection()
	account := NewAccountMySQLContext(conn.GetConnection())
	device := NewDeviceMySQLContext(conn.GetConnection())
	repo := Repository{&account, &device}
	return repo
}

type Repository struct {
	accountContext IAccountContext
	deviceContext IDeviceContext
}

func (this *Repository) Login(username string, password string) bool{
	return this.accountContext.login(username, password)
}

func (this *Repository) Register(username string, password string, long float64, lat float64) bool {
	return this.accountContext.register(username, password, long, lat)
}

func (this *Repository) CreateDevice(device models.Device) bool {
	return this.deviceContext.createDevice(device)
}

func (this *Repository) GetDevices() []models.Device {
	return this.deviceContext.getDevices()
}

func (this *Repository) GetDeviceTypes() []models.DeviceType {
	return this.deviceContext.getDeviceTypes()
}
 