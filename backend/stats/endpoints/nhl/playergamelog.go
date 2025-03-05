package nhl

import (
	"encoding/json"
	"fmt"
	"sports_api/globals/nhl"
)

type GameLog struct {
	Assists           int    `json:"assists"`
	GameDate          string `json:"gameDate"`
	GameId            int    `json:"gameId"`
	GameWinningGoals  int    `json:"gameWinningGoals"`
	Goals             int    `json:"goals"`
	HomeRoadFlag      string `json:"homeRoadFlag"`
	OpponentAbbrev    string `json:"opponentAbbrev"`
	OtGoals           int    `json:"otGoals"`
	Pim               int    `json:"pim"`
	PlusMinus         int    `json:"plusMinus"`
	Points            int    `json:"points"`
	PowerPlayGoals    int    `json:"powerPlayGoals"`
	PowerPlayPoints   int    `json:"powerPlayPoints"`
	Shifts            int    `json:"shifts"`
	ShorthandedGoals  int    `json:"shorthandedGoals"`
	ShorthandedPoints int    `json:"shorthandedPoints"`
	Shots             int    `json:"shots"`
	TeamAbbrev        string `json:"teamAbbrev"`
	Toi               string `json:"toi"`
}

type GoalieGameLog struct {
	Assists        int     `json:"assists"`
	Decision       string  `json:"decision,omitempty"`
	GameDate       string  `json:"gameDate"`
	GameId         int     `json:"gameId"`
	GamesStarted   int     `json:"gamesStarted"`
	Goals          int     `json:"goals"`
	GoalsAgainst   int     `json:"goalsAgainst"`
	HomeRoadFlag   string  `json:"homeRoadFlag"`
	OpponentAbbrev string  `json:"opponentAbbrev"`
	Pim            int     `json:"pim"`
	SavePctg       float64 `json:"savePctg"`
	ShotsAgainst   int     `json:"shotsAgainst"`
	Shutouts       int     `json:"shutouts"`
	TeamAbbrev     string  `json:"teamAbbrev"`
	Toi            string  `json:"toi"`
}

func (p *Player) GetGameLog(seasonYear string, seasonType int) error {

	nhl.NHLSession.SetBaseUrl("https://api-web.nhle.com/v1/")

	resp, _ := nhl.NHLSession.NHLGetRequest(fmt.Sprintf("player/%d/game-log/%s/%d", p.Id, seasonYear, seasonType), nil, "", nil)

	marsh, err := json.Marshal(resp.Data.(map[string]interface{})["gameLog"])
	if err != nil {
		return err
	}
	if p.PositionCode != "G" {
		err = json.Unmarshal(marsh, &p.GameLogs)
		if err != nil {
			return err
		}
	} else {
		err = json.Unmarshal(marsh, &p.GoalieLogs)
		if err != nil {
			return err
		}
	}

	return nil
}
