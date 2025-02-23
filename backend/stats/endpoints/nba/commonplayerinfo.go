package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CommonPlayerInfo calls the NBA API to retrieve detailed information about a player.
//
// Parameters:
// - playerID (string): The ID of the player to retrieve information for.
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
//	playerData, err := CommonPlayerInfo("2544", &league)
//	if err != nil {
//	    log.Fatal("Error fetching player info:", err)
//	}
//	fmt.Println(playerData)
func CommonPlayerInfo(playerID string, leagueID *string) (*client.NBAResponse, error) {
	if err := validateCommonPlayerInfoParams(playerID, leagueID); err != nil {
		return nil, err
	}

	params := map[string]string{
		"PlayerID": playerID,
	}

	// Include LeagueID in params only if it is not nil
	if leagueID != nil {
		params["LeagueID"] = *leagueID
	}

	return client.NBASession.NBAGetRequest(endpoints.CommonPlayerInfo, params, "", nil)
}

// validateCommonPlayerInfoParams ensures all input parameters are valid.
func validateCommonPlayerInfoParams(playerID string, leagueID *string) error {
	// Validate playerID using helper function
	if valid, err := helpers.ValidatePlayerID(playerID); !valid {
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
