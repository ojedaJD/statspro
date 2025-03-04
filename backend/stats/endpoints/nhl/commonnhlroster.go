package nhl

import (
	"encoding/json"
	"fmt"
	"sports_api/globals/nhl"
)

type NHLRoster struct {
	Defensemen []struct {
		FirstName struct {
			Default string `json:"default"`
		} `json:"firstName"`
		Id       int `json:"id"`
		LastName struct {
			Default string `json:"default"`
		} `json:"lastName"`
		PositionCode  string `json:"positionCode"`
		ShootsCatches string `json:"shootsCatches"`
		SweaterNumber int    `json:"sweaterNumber"`
	} `json:"defensemen"`
	Forwards []struct {
		FirstName struct {
			Default string `json:"default"`
		} `json:"firstName"`
		Id       int `json:"id"`
		LastName struct {
			Default string `json:"default"`
		} `json:"lastName"`
		PositionCode  string `json:"positionCode"`
		ShootsCatches string `json:"shootsCatches"`
		SweaterNumber int    `json:"sweaterNumber"`
	} `json:"forwards"`
	Goalies []struct {
		FirstName struct {
			Default string `json:"default"`
		} `json:"firstName"`
		Id       int `json:"id"`
		LastName struct {
			Default string `json:"default"`
		} `json:"lastName"`
		PositionCode  string `json:"positionCode"`
		ShootsCatches string `json:"shootsCatches"`
		SweaterNumber int    `json:"sweaterNumber"`
	} `json:"goalies"`
}

// GetNHLRoster SeasonID in form 20242025
func GetNHLRoster(teamAbbreviation, seasonID string) (*nhl.NHLResponse, error) {
	nhl.NHLSession.SetBaseUrl("https://api-web.nhle.com/v1/")

	return nhl.NHLSession.NHLGetRequest(fmt.Sprintf("roster/%s/%s", teamAbbreviation, seasonID), nil, "", nil)
}

func GetAndParseNHLRoster(teamAbbreviation, seasonID string) NHLRoster {
	teams, err := GetNHLRoster(teamAbbreviation, seasonID)
	if err != nil {
		fmt.Println(err)
	}
	marshal, err := json.Marshal(teams.Data)
	if err != nil {
		return NHLRoster{}
	}
	var roster NHLRoster
	err = json.Unmarshal(marshal, &roster)
	if err != nil {
		return NHLRoster{}
	}

	return roster

}

//r.json()["data"]
