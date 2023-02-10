package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func WebArchiveDownloadHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//ctx := request.Context()
		specUrl := "http://motionbuscard.org.cy/opendata/downloadfile?file=GTFS%5C6_google_transit.zip&rel=True"

		resp, err := http.Get(specUrl)
		if err != nil {
			fmt.Printf("err: %s", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			http.Error(writer, "response status code is not 200", http.StatusInternalServerError)
		}
		log.Println("status", resp.Status)

		// Create the file
		out, err := os.Create("gtfs_limassol.zip")
		if err != nil {
			fmt.Printf("err: %s", err)
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		log.Println("Archive saved")
	}
}
