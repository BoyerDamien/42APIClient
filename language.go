package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Language struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

// GetLanguage fetch language data from 42 api from the language id
func (s *APIClient) GetLanguage(id string) (Language, error) {
	endpoint := fmt.Sprintf("%s/v2/languages/%s", s.Url, id)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Language{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token.AccessToken))
	resp, err := client.Do(req)
	if err != nil {
		return Language{}, err
	}
	if resp.StatusCode != 200 {
		return Language{}, fmt.Errorf(resp.Status)
	}
	body, err := ReadHTTPResponse(resp)
	if err != nil {
		return Language{}, err
	}
	var language Language
	_ = json.Unmarshal(body, &language)
	return language, nil
}
