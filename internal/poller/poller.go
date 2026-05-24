package poller

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func start() {

	for {
		flightAPI, err := http.Get("https://example.com")

		if err != nil {
			log.Fatal(err)

			resp, err := http.Get("https://example.com")
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close() // Always close the body to prevent leaks

			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}
	}

}
