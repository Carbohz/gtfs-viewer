package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func WebArchiveDownloadHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//ctx := request.Context()
		url := "http://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png"

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(len(body))
		//download the file in browser
	}
}