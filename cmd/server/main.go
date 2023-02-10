package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)

func main() {
	var (
		username = "YOUR_ACCESS_KEY"
		password = "YOUR_SECRET_KEY"
	)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "URL OF YOUR GTFS-REALTIME SOURCE GOES HERE", nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range feed.Entity {
		tripUpdate := entity.GetTripUpdate()
		trip := tripUpdate.GetTrip()
		tripId := trip.GetTripId()
		fmt.Printf("Trip ID: %s\n", tripId)
	}
}
