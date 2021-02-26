package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AuthToken struct defines 42 api oAuth Token response
type AuthToken struct {
	AccessToken string    `json:"access_Token"`
	TokenType   string    `json:"Token_type"`
	ExpiresIn   int       `json:"expires_in"`
	LastUpdate  time.Time `json:"last_update"`
}

// API 42 interface
type API interface {
	Auth() error
	GetUser(login string) (User, error)
	Token() AuthToken
}

// APIClient implements 42 API interface
type APIClient struct {
	Url    string
	Uid    string
	Secret string
	Token  AuthToken
}

// Auth method implements 42 API oAuth authentication
func (s *APIClient) Auth() error {
	endpoint := BuildAuthURL(s.Url, s.Uid, s.Secret)
	response, err := http.Post(endpoint, "application/x-www-form-Urlencoded", nil)
	if err != nil {
		return err
	}
	body, err := ReadHTTPResponse(response)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &s.Token)
	s.Token.LastUpdate = time.Now()
	return nil
}

// GetUser fetch user data form 42 api based on login
func (s *APIClient) GetUser(login string) (User, error) {
	endpoint := fmt.Sprintf("%s/v2/users/%s", s.Url, login)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.Token.AccessToken))
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	if resp.StatusCode != 200 {
		return User{}, fmt.Errorf(resp.Status)
	}
	body, err := ReadHTTPResponse(resp)
	if err != nil {
		return User{}, err
	}
	var user User
	json.Unmarshal(body, &user)
	return user, nil
}
