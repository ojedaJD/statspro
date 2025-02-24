package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// GameRotation retrieves game rotation data for a given game.
func GameRotation(gameID string, leagueID *string) (*client.NBAResponse, error) {
	// Validate required parameters
	if valid, err := helpers.ValidateGameID(gameID); !valid {
		return nil, err
	}

	// Validate LeagueID if provided
	if leagueID != nil {
		if valid, err := helpers.ValidateLeagueID(*leagueID); !valid {
			return nil, err
		}
	}

	params := map[string]string{
		"GameID": gameID,
	}

	if leagueID != nil {
		params["LeagueID"] = *leagueID
	}

	return client.NBASession.NBAGetRequest(endpoints.GameRotation, params, "", nil)
}
