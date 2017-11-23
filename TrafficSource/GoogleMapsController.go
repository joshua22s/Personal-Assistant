package trafficsource

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var (
	client *maps.Client
)

func Start(mapskey string) {
	client = getClient(mapskey)
}

func getClient(mapskey string) *maps.Client {
	clnt, err := maps.NewClient(maps.WithAPIKey(mapskey))
	if err != nil {
		log.Fatalf("Unable to create maps client: %v", err)
	}
	return clnt
}

func GetTravelTime(origin string, destination string, mode string, arrivalTime time.Time) time.Time {
	r := &maps.DistanceMatrixRequest{}
	r.Origins = append(r.Origins, origin)
	r.Destinations = append(r.Destinations, destination)
	lookupMode(mode, r)
	resp, err := client.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("Unable te receive request from maps API: %v", err)
	}
	var departureTime time.Time
	for _, row := range resp.Rows {
		for _, element := range row.Elements {
			fmt.Println(element.Duration)
			fmt.Println(element.Distance)
			departureTime = arrivalTime.Add(-element.Duration)
		}
	}
	return departureTime
}

func lookupMode(mode string, r *maps.DistanceMatrixRequest) {

	switch mode {
	case "driving":
		r.Mode = maps.TravelModeDriving
	case "walking":
		r.Mode = maps.TravelModeWalking
	case "bicycling":
		r.Mode = maps.TravelModeBicycling
	case "transit":
		r.Mode = maps.TravelModeTransit
	case "":
		// ignore
	}
}
