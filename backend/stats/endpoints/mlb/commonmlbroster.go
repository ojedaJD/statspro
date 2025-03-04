package mlb

import (
	"encoding/json"
	"sports_api/globals/mlb"
)

type MLBPlayer struct {
	JerseyNumber string `json:"jerseyNumber"`
	ParentTeamId int    `json:"parentTeamId"`
	Person       struct {
		FullName string `json:"fullName"`
		Id       int    `json:"id"`
		Link     string `json:"link"`
	} `json:"person"`
	Position struct {
		Abbreviation string `json:"abbreviation"`
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"position"`
	Status struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Hitting  []HittingStats
	Pitching []PitchingStats
}

func (team *MLBTeam) GetRoster() error {
	params := map[string]string{}

	resp, err := mlb.MLBSession.MLBGetRequest(team.Link+"/roster", params, "", nil)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(resp.Data.(map[string]interface{})["roster"])
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, &team.Roster)
	if err != nil {
		return err
	}

	return nil
}
