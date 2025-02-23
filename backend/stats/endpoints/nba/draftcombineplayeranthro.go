package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// DraftCombinePlayerAnthro calls the NBA API and retrieves player anthropometric data.
func DraftCombinePlayerAnthro(leagueID, seasonYear string) (*client.NBAResponse, error) {
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
	return client.NBASession.NBAGetRequest(endpoints.DraftCombinePlayerAnthro, params, "", nil)
}

// validateDraftCombineParams ensures leagueID and seasonYear are valid.
func validateDraftCombineParams(leagueID, seasonYear string) error {
	// Validate LeagueID
	if _, err := helpers.ValidateLeagueID(leagueID); err != nil {
		return err
	}

	// Validate SeasonYear
	if _, err := helpers.ValidateSeasonYear(seasonYear); err != nil {
		return err
	}

	return nil
}
