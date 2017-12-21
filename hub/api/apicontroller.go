package api

import (
	"strconv"
	"encoding/json"
	"net/http"
	"log"
	hub "github.com/joshua22s/hub/hub"
	models "github.com/joshua22s/hub/models"
	"github.com/gorilla/mux"
)

func NewApiController() ApiController {
	return ApiController{hub.NewHubController()}
}

type ApiController struct {
	hub hub.HubController
}

func (this *ApiController) StartApi(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/devicetypes", this.getDeviceTypes).Methods("GET")
	router.HandleFunc("/devices", this.getDevices).Methods("GET")
	router.HandleFunc("/devices", this.createDevice).Methods("POST")
	log.Println("Started API on " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), router))
}

func (this *ApiController) getDeviceTypes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(this.hub.GetDeviceTypes())
}

func (this *ApiController) getDevices(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(this.hub.GetDevices())
}

func (this *ApiController) createDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	if (this.hub.CreateDevice(device)) {
		json.NewEncoder(w).Encode(this.hub.GetDevices())
	}
}
 