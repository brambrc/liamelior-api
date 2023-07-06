package Cronjob

import (
	"fmt"
	"net/http"
)

func TriggerAPICronJob() {
	// Make a request to the API endpoint
	url := "/cron-job/cron-get-show-schedule"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	// Add any required headers or authentication tokens to the request
	// if needed

	// Send the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code and handle accordingly
	if resp.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status code:", resp.StatusCode)
		return
	}

	// Process the response body as needed
	// You can read and parse the response data here
}
