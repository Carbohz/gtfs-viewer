package handlers

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/Carbohz/gtfs-viewer/api/rest/models"
	"github.com/gocarina/gocsv"
	"io"
	"io/ioutil"
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

		body, err := ioutil.ReadAll(resp.Body)
		//log.Println(string(body))
		zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
		if err != nil {
			log.Fatal(err)
		}

		// Read all the files from zip archive
		for _, zipFile := range zipReader.File {
			if zipFile.Name == "routes.txt" {
				log.Println("Reading file:", zipFile.Name)
				unzippedFileBytes, err := readZipFile(zipFile)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println(string(unzippedFileBytes))

				var routes []*models.Route
				if err := gocsv.UnmarshalBytes(unzippedFileBytes, &routes); err != nil {
					log.Printf("Failed to unmarshal routes.txt: %s", string(unzippedFileBytes))
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}
				log.Println(routes[0])
			}
		}

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

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
