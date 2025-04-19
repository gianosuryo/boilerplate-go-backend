package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendHTTPRequest(method string, url string, header map[string]string, formData map[string]interface{}) (map[string]interface{}, error) {
	// Create HTTP client
	client := &http.Client{}

	// Prepare request body
	formDataJson, err := json.Marshal(formData)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
		return nil, err
	}

	// Create POST request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(formDataJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Process response
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}

	return responseData, nil
}
