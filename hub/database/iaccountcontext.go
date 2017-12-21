package database

type IAccountContext interface {
	register(username string, password string, long float64, lat float64) bool
	login(username string, password string) bool
}