package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CommonTeamRoster calls the NBA API to retrieve the roster for a given team in a specific season.
//
// Parameters:
// - teamID (string): The ID of the team to retrieve the roster for.
// - season (string): The season in "YYYY-YY" format (e.g., "2019-20").
// - leagueID (*string): The league identifier (nullable).
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
//	league := "00"
//	roster, err := CommonTeamRoster("1610612739", "2019-20", &league)
//	if err != nil {
//	    log.Fatal("Error fetching team roster:", err)
//	}
//	fmt.Println(roster)
func CommonTeamRoster(teamID, season string, leagueID *string) (*client.NBAResponse, error) {
	if err := validateCommonTeamRosterParams(teamID, season, leagueID); err != nil {
		return nil, err
	}

	params := map[string]string{
		"TeamID": teamID,
		"Season": season,
	}

	// Include LeagueID in params only if it is not nil
	if leagueID != nil {
		params["LeagueID"] = *leagueID
	}

	return client.NBASession.NBAGetRequest(endpoints.CommonTeamRoster, params, "", nil)
}

// validateCommonTeamRosterParams ensures all input parameters are valid.
func validateCommonTeamRosterParams(teamID, season string, leagueID *string) error {
	// Validate teamID using helper function
	if valid, err := helpers.ValidateTeamID(teamID); !valid {
		return err
	}

	// Validate Season using helper function
	if valid, err := helpers.ValidateSeason(season); !valid {
		return err
	}

	// Validate LeagueID only if it's provided (not nil)
	if leagueID != nil {
		if valid, err := helpers.ValidateLeagueID(*leagueID); !valid {
			return err
		}
	}

	return nil
}
