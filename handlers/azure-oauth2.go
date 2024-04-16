package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type TokenResponseSuccess struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GenericHTTPRequest(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	// Create a new request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func RequestAccessTokenInfo(tenantId string) (*TokenResponseSuccess, error) {
	var (
		client_id     = os.Getenv("AZ_CLIENT_ID")
		client_secret = os.Getenv("AZ_CLIENT_SECRET")
		scope         = os.Getenv("AZ_OAUTH_SCOPE")
	)

	if tenantId == "" {
		return nil, fmt.Errorf("tenantId is empty")
	}

	authUrl := "https://login.microsoftonline.com/" + tenantId + "/oauth2/v2.0/token"

	bodyData := url.Values{}
	bodyData.Set("client_id", client_id)
	bodyData.Set("client_secret", client_secret)
	bodyData.Set("grant_type", "client_credentials")
	bodyData.Set("scope", scope)
	body := bodyData.Encode()
	bodyBuffer := bytes.NewBuffer([]byte(body))

	// Create request
	request, err := http.NewRequest(http.MethodPost, authUrl, bodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(url.QueryEscape(client_id), url.QueryEscape(client_secret))

	// post request
	client := &http.Client{}
	httpResponse, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error on request: %v", err)
	}
	defer httpResponse.Body.Close()

	responseBuffer, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading http response body: %v", err)
	}

	if httpResponse.StatusCode != http.StatusOK {
		// Pretty format error output
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, responseBuffer, "", "    ")
		return nil, fmt.Errorf("response returned with status-code %d and body:\n%s", httpResponse.StatusCode, prettyJSON.String())
	}

	var response TokenResponseSuccess
	err = json.Unmarshal(responseBuffer, &response)
	if err != nil {
		return nil, fmt.Errorf("error on unmarshal: %v", err)
	}

	return &response, nil
}

func RequestAccessToken(tenantId string) (string, error) {

	response, err := RequestAccessTokenInfo(tenantId)
	if err != nil {
		return "", err
	}

	if response.AccessToken == "" {
		return "", fmt.Errorf("received access-token is empty")
	}

	return response.AccessToken, nil
}
