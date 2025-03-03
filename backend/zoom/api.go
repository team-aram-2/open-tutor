package zoom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var ACCESS_TOKEN_REFRESH_INTERVAL time.Duration = 10 * time.Minute

var cachedAccessToken *string

func refreshAccessToken() error {
	fmt.Printf("Refreshing Zoom access token...\n")
	accountId := os.Getenv("ZOOM_ACCOUNT_ID")
	clientId := os.Getenv("ZOOM_CLIENT_ID")
	clientSecret := os.Getenv("ZOOM_CLIENT_SECRET")

	url := fmt.Sprintf("https://zoom.us/oauth/token?grant_type=account_credentials&account_id=%s", accountId)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("error creating token request: %v", err)
	}

	req.SetBasicAuth(clientId, clientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error getting token: %v", err)
	}
	defer resp.Body.Close()

	// Decode response as JSON //
	respBodyStream, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read access token response: %s\n", err)
	}
	var responseBody struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(respBodyStream, &responseBody); err != nil {
		errMsg := fmt.Errorf("error parsing Zoom access token repsonse: %s", err)
		fmt.Println(errMsg)
		return errMsg
	}

	cachedAccessToken = &responseBody.AccessToken
	fmt.Printf("Zoom access token refreshed\n")
	return nil
}

func zoomApiRequest[B any](path string, requestType string, requestBodyPayload any, requestHeaders map[string]string) (*B, error) {
	var jsonData []byte
	var err error
	if requestBodyPayload != nil {
		jsonData, err = json.Marshal(requestBodyPayload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling JSON for Zoom API request: %v", err)
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

func LoadAccessToken() {
	var err error

	// Refresh token on interval //
	go func() {
		for {
			err = refreshAccessToken()
			if err != nil {
				fmt.Printf("error fetching Zoom access token: %s\n", err)
			}

			time.Sleep(ACCESS_TOKEN_REFRESH_INTERVAL)
		}
	}()
}
