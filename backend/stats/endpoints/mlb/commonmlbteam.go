package mlb

import (
	"encoding/json"
	"fmt"
	"sports_api/globals/mlb"
)

type MLBTeam struct {
	Abbreviation  string `json:"abbreviation"`
	Active        bool   `json:"active"`
	AllStarStatus string `json:"allStarStatus"`
	ClubName      string `json:"clubName"`
	Division      struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"division"`
	FileCode        string `json:"fileCode"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	FranchiseName   string `json:"franchiseName"`
	Id              int    `json:"id"`
	Link            string `json:"link"`
	LocationName    string `json:"locationName"`
	Name            string `json:"name"`
	Season          int    `json:"season"`
	ShortName       string `json:"shortName"`
	TeamCode        string `json:"teamCode"`
	TeamName        string `json:"teamName"`
	Roster          []MLBPlayer
}

func GetAllMLBTeams() (*mlb.MLBResponse, error) {
	params := map[string]string{
		"sportId": "1",
	}

	endpoint := "/api/v1/teams"

	return mlb.MLBSession.MLBGetRequest(endpoint, params, "", nil)
}

func GetAndParseMLBTeams() []*MLBTeam {
	resp, err := GetAllMLBTeams()
	if err != nil {
		return nil
	}
	i := resp.Data.(map[string]interface{})["teams"]
	var teams []*MLBTeam
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(marshal, &teams)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, v := range teams {
		err = v.GetRoster()
	}
	return teams
}
