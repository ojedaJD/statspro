package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// FranchiseHistory retrieves the history of NBA franchises.
func FranchiseHistory(leagueID string) (*client.NBAResponse, error) {
	// Validate input parameters
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return nil, err
	}

	params := map[string]string{
		"LeagueID": leagueID,
	}

	return client.NBASession.NBAGetRequest(endpoints.FranchiseHistory, params, "", nil)
}
