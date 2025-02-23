package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CommonTeamYears calls the NBA API to retrieve available team years based on the league.
//
// Parameters:
// - leagueID (string): The league identifier (must be a two-digit number).
//   - "00" → NBA (default)
//   - "10" → WNBA
//   - "20" → G-League
//
// Returns:
// - (*client.NBAResponse): The JSON response from the API, decoded into a struct.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	years, err := CommonTeamYears("00")
//	if err != nil {
//	    log.Fatal("Error fetching team years:", err)
//	}
//	fmt.Println(years)
func CommonTeamYears(leagueID string) (*client.NBAResponse, error) {
	if err := validateCommonTeamYearsParams(leagueID); err != nil {
		return nil, err
	}

	params := map[string]string{
		"LeagueID": leagueID,
	}

	return client.NBASession.NBAGetRequest(endpoints.CommonTeamYears, params, "", nil)
}

// validateCommonTeamYearsParams ensures the leagueID is valid.
func validateCommonTeamYearsParams(leagueID string) error {
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return err
	}
	return nil
}
