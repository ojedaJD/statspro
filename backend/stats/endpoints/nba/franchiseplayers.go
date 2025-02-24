package nba

import (
	"errors"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// FranchisePlayers retrieves player statistics for a given franchise.
func FranchisePlayers(leagueID, perMode, seasonType, teamID string) (*client.NBAResponse, error) {
	// Validate required parameters
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return nil, err
	}
	if valid, err := helpers.ValidatePerMode(perMode); !valid {
		return nil, err
	}
	if valid, err := helpers.ValidateSeasonType(seasonType); !valid {
		return nil, err
	}
	if teamID == "" {
		return nil, errors.New("TeamID is required")
	}

	params := map[string]string{
		"LeagueID":   leagueID,
		"PerMode":    perMode,
		"SeasonType": seasonType,
		"TeamID":     teamID,
	}

	return client.NBASession.NBAGetRequest(endpoints.FranchisePlayers, params, "", nil)
}
