package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CommonPlayoffSeries calls the NBA API to retrieve playoff series details based on the provided filters.
//
// Parameters:
// - leagueID (string): The league identifier (required).
//   - "00" → NBA (default)
//   - "10" → WNBA
//   - "20" → G-League
//
// - season (string): The NBA season format as "YYYY-YY" (e.g., "2019-20").
//
// - seriesID (*string): The series identifier (nullable). If nil, retrieves all available playoff series.
//
// Returns:
// - (*client.NBAResponse): The JSON response from the API, decoded into a struct.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	seriesID := "0041900401"
//	playoffData, err := CommonPlayoffSeries("00", "2019-20", &seriesID)
//	if err != nil {
//	    log.Fatal("Error fetching playoff series:", err)
//	}
//	fmt.Println(playoffData)
func CommonPlayoffSeries(leagueID, season string, seriesID *string) (*client.NBAResponse, error) {
	if err := validateCommonPlayoffSeriesParams(leagueID, season); err != nil {
		return nil, err
	}

	params := map[string]string{
		"LeagueID": leagueID,
		"Season":   season,
	}

	// Include SeriesID in params only if it is not nil
	if seriesID != nil {
		params["SeriesID"] = *seriesID
	}

	return client.NBASession.NBAGetRequest(endpoints.CommonPlayoffSeries, params, "", nil)
}

// validateCommonPlayoffSeriesParams ensures all input parameters are valid.
func validateCommonPlayoffSeriesParams(leagueID, season string) error {
	// Validate LeagueID using helper function
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return err
	}

	// Validate Season using helper function
	if valid, err := helpers.ValidateSeason(season); !valid {
		return err
	}

	return nil
}
