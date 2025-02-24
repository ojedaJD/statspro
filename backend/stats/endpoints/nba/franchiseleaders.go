package nba

import (
	"errors"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// FranchiseLeaders retrieves the all-time statistical leaders for a given franchise.
func FranchiseLeaders(teamID string, leagueID *string) (*client.NBAResponse, error) {
	// Validate input parameters
	if teamID == "" {
		return nil, errors.New("TeamID is required")
	}

	if leagueID != nil {
		if valid, err := helpers.ValidateLeagueID(*leagueID); !valid {
			return nil, err
		}
	}

	params := map[string]string{
		"TeamID": teamID,
	}

	if leagueID != nil {
		params["LeagueID"] = *leagueID
	}

	return client.NBASession.NBAGetRequest(endpoints.FranchiseLeaders, params, "", nil)
}
