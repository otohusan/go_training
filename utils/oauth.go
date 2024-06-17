package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GoogleUserInfo holds the information about the user from Google
type GoogleUserInfo struct {
	ID    string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// FetchGoogleUserInfo retrieves user information from Google using the access token
func FetchGoogleUserInfo(accessToken string) (*GoogleUserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: %s", resp.Status)
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
