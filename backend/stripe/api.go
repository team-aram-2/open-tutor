package stripe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func stripeApiRequest[B any](path string, requestType string, requestBodyPayload any, requestHeaders map[string]string) (*B, error) {
	var jsonData []byte
	var err error
	if requestBodyPayload != nil {
		jsonData, err = json.Marshal(requestBodyPayload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling JSON for Stripe API request: %v", err)
		}
	} else {
		// Empty JSON object if body not supplied //
		jsonData = []byte("{}")
	}

	// Create the Zoom API request //
	req, err := http.NewRequest(requestType, "https://api.zoom.us/"+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add required headers //
	req.Header.Set("Authorization", "Bearer "+(*cachedAccessToken))
	req.Header.Set("Content-Type", "application/json")
	// Add custom headers if provided //
	for key, value := range requestHeaders {
		req.Header.Set(key, value)
	}

	// Send the request //
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// Decode response as JSON //
	var responseBodyPayload B
	if err := json.Unmarshal(responseBody, &responseBodyPayload); err != nil {
		errMsg := fmt.Errorf("error parsing Zoom API response to %s: %s", path, err)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	return &responseBodyPayload, nil
}
