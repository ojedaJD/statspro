package nba

import (
	client "sports_api/globals/nba"
	endpoints "sports_api/urls/nba"
)

// DraftCombineSpotShooting calls the NBA API and retrieves spot shooting drill results.
func DraftCombineSpotShooting(leagueID, seasonYear string) (*client.NBAResponse, error) {
	// Validate parameters
	if err := validateDraftCombineParams(leagueID, seasonYear); err != nil {
		return nil, err
	}

	// Construct query parameters
	params := map[string]string{
		"LeagueID":   leagueID,
		"SeasonYear": seasonYear,
	}

	// Make API request
	return client.NBASession.NBAGetRequest(endpoints.DraftCombineSpotShooting, params, "", nil)
}
