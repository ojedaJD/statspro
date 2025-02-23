package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// DraftCombineStats calls the NBA API and retrieves combine stats for a given season.
func DraftCombineStats(leagueID, seasonYear string) (*client.NBAResponse, error) {
	// Validate parameters
	if err := validateDraftCombineStatsParams(leagueID, seasonYear); err != nil {
		return nil, err
	}

	// Construct query parameters
	params := map[string]string{
		"LeagueID":   leagueID,
		"SeasonYear": seasonYear,
	}

	// Make API request
	return client.NBASession.NBAGetRequest(endpoints.DraftCombineStats, params, "", nil)
}

// validateDraftCombineStatsParams ensures leagueID and seasonYear are valid for DraftCombineStats.
func validateDraftCombineStatsParams(leagueID, seasonYear string) error {
	// Validate LeagueID
	if _, err := helpers.ValidateLeagueID(leagueID); err != nil {
		return err
	}

	// Validate SeasonYear (must be either "YYYY-YY" format or "All Time")
	if seasonYear != "All Time" {
		if _, err := helpers.ValidateSeason(seasonYear); err != nil {
			return err
		}
	}

	return nil
}
