package nba

import (
	"encoding/json"
	"errors"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// PlayerGameLog retrieves game log statistics for a specific player.
func PlayerGameLog(playerID, season, seasonType string, leagueID *string) (*client.NBAResponse, error) {
	// Validate required parameters
	if playerID == "" {
		return nil, errors.New("PlayerID is required")
	}
	if _, e := helpers.ValidateSeason(season); e != nil {
		return nil, e
	}
	if valid, err := helpers.ValidateSeasonType(seasonType); !valid {
		return nil, err
	}

	// Validate optional LeagueID if provided
	if leagueID != nil {
		if valid, err := helpers.ValidateLeagueID(*leagueID); !valid {
			return nil, err
		}
	}

	params := map[string]string{
		"PlayerID":   playerID,
		"Season":     season,
		"SeasonType": seasonType,
	}

	if leagueID != nil {
		params["LeagueID"] = *leagueID
	}

	return client.NBASession.NBAGetRequest(endpoints.PlayerGameLog, params, "", nil)
}

type GameLog struct {
	AST       int     `json:"AST"`
	BLK       int     `json:"BLK"`
	DREB      int     `json:"DREB"`
	FG3A      int     `json:"FG3A"`
	FG3M      int     `json:"FG3M"`
	FG3PCT    float64 `json:"FG3_PCT"`
	FGA       int     `json:"FGA"`
	FGM       int     `json:"FGM"`
	FGPCT     float64 `json:"FG_PCT"`
	FTA       int     `json:"FTA"`
	FTM       int     `json:"FTM"`
	FTPCT     float64 `json:"FT_PCT"`
	GAMEDATE  string  `json:"GAME_DATE"`
	GameID    string  `json:"Game_ID"`
	MATCHUP   string  `json:"MATCHUP"`
	MIN       int     `json:"MIN"`
	OREB      int     `json:"OREB"`
	PLUSMINUS int     `json:"PLUS_MINUS"`
	PTS       int     `json:"PTS"`
	REB       int     `json:"REB"`
	STL       int     `json:"STL"`
	TOV       int     `json:"TOV"`
	WL        string  `json:"WL"`
}

func GetPlayerGameLog(playerID, season, seasonType string, leagueID *string) []GameLog {
	log, err := PlayerGameLog(playerID, season, seasonType, leagueID)
	if err != nil {
		return nil
	}
	dict2, err := log.GetNormalizedDict2()
	marshal, err := json.Marshal(dict2["PlayerGameLog"])
	if err != nil {
		return nil
	}

	var GameLogs []GameLog

	err = json.Unmarshal(marshal, &GameLogs)
	if err != nil {
		return nil
	}
	return GameLogs

}
