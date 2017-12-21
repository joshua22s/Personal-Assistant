package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	models "github.com/joshua22s/hub/models"
)

func NewDeviceMySQLContext(db *sql.DB) DeviceMySQLContext {
	context := DeviceMySQLContext{db}
	return context
}

type DeviceMySQLContext struct {
	db *sql.DB
}

func (this *DeviceMySQLContext) createDevice(device models.Device) bool {
	stm, err := this.db.Prepare("INSERT INTO Device(name, type) VALUES(?,?)")
	defer stm.Close()
	if (err != nil) {
		panic(err.Error())
	}
	res, err := stm.Exec(device.Name, device.Type.Id)
	if (err!= nil || res == nil) {
		return false
	}
	return true
}

func (this *DeviceMySQLContext) getDevices() []models.Device {
	var (
		devices []models.Device
		deviceName string
		typeId int
		typeName string		
	)
	rows, err := this.db.Query("SELECT d.name, dt.id, dt.name FROM Device d JOIN Devicetype dt ON d.type = dt.id")
	defer rows.Close()
	if (err != nil) {
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&deviceName, &typeId, &typeName)
		if (err != nil) {
			return nil
		}
		devices = append(devices, models.Device{Name: deviceName, Type: models.DeviceType{Id: typeId, Name: typeName}})
	}
	return devices
}

func (this *DeviceMySQLContext) getDeviceTypes() []models.DeviceType {
	var (
		deviceTypes []models.DeviceType
		typeId int
		typeName string
	)
	rows, err := this.db.Query("SELECT id, name FROM Devicetype")
	defer rows.Close()
	if (err != nil) {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&typeId, &typeName)
		if (err != nil) {
			fmt.Println(err)
			return nil
		}
		deviceTypes = append(deviceTypes, models.DeviceType{Id: typeId, Name: typeName})
	}
	return deviceTypes
}