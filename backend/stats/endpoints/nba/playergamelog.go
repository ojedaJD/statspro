package nba

import (
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
