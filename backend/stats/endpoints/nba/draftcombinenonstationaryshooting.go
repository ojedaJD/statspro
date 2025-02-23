package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// DraftCombineNonStationaryShooting calls the NBA API and retrieves non-stationary shooting drill results.
//
// Example Usage:
//
//	data, err := DraftCombineNonStationaryShooting("00", "2019")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(data)
func DraftCombineNonStationaryShooting(leagueID, seasonYear string) (*client.NBAResponse, error) {
	// Validate parameters
	if err := validateDraftCombineNonStationaryShootingParams(leagueID, seasonYear); err != nil {
		return nil, err
	}

	// Construct query parameters
	params := map[string]string{
		"LeagueID":   leagueID,
		"SeasonYear": seasonYear,
	}

	// Make API request
	return client.NBASession.NBAGetRequest(endpoints.DraftCombineNonStationaryShooting, params, "", nil)
}

// validateDraftCombineNonStationaryShootingParams ensures all input parameters are valid.
func validateDraftCombineNonStationaryShootingParams(leagueID, seasonYear string) error {
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
