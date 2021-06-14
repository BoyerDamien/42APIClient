package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Campus struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	TimeZone string `json:"time_zone"`
	Language struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Identifier string `json:"identifier"`
	} `json:"language"`
	UsersCount  int         `json:"users_count"`
	VogsphereID int         `json:"vogsphere_id"`
	Endpoint    interface{} `json:"endpoint"`
}

// GetCampus fetch campus data from 42 api from the campus id
func (s *APIClient) GetCampus(id string) (Campus, error) {
	endpoint := fmt.Sprintf("%s/v2/campus/%s", s.Url, id)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Campus{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token.AccessToken))
	resp, err := client.Do(req)
	if err != nil {
		return Campus{}, err
	}
	if resp.StatusCode != 200 {
		return Campus{}, fmt.Errorf(resp.Status)
	}
	body, err := ReadHTTPResponse(resp)
	if err != nil {
		return Campus{}, err
	}
	var campus Campus
	_ = json.Unmarshal(body, &campus)
	return campus, nil
}
