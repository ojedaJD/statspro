package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CumeStatsPlayerGamesOptions defines the query parameters for the CumeStatsPlayerGames API.
type CumeStatsPlayerGamesOptions struct {
	LeagueID     string
	PlayerID     string
	Season       string
	SeasonType   string
	VsTeamID     string
	VsDivision   string
	VsConference string
	Outcome      string
	Location     string
}

// CumulativeStatsPlayerGames calls the NBA API to retrieve cumulative statistics for a player over multiple games.
//
// Parameters:
// - opts (CumeStatsPlayerGamesOptions): Struct containing all query parameters.
//
// Returns:
// - (*client.NBAResponse): The JSON response from the API, decoded into a struct.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	opts := CumeStatsPlayerGamesOptions{
//	    LeagueID:   "00",
//	    PlayerID:   "2544",
//	    Season:     "2019-20",
//	    SeasonType: "Regular Season",
//	}
//	stats, err := CumeStatsPlayerGames(opts)
//	if err != nil {
//	    log.Fatal("Error fetching cumulative player game stats:", err)
//	}
//	fmt.Println(stats)
func CumulativeStatsPlayerGames(opts CumeStatsPlayerGamesOptions) (*client.NBAResponse, error) {
	if err := validateCumulativeStatsPlayerGamesParams(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"LeagueID":   opts.LeagueID,
		"PlayerID":   opts.PlayerID,
		"Season":     opts.Season,
		"SeasonType": opts.SeasonType,
	}

	// Include optional parameters if they are not empty
	if opts.VsTeamID != "" {
		params["VsTeamID"] = opts.VsTeamID
	}
	if opts.VsDivision != "" {
		params["VsDivision"] = opts.VsDivision
	}
	if opts.VsConference != "" {
		params["VsConference"] = opts.VsConference
	}
	if opts.Outcome != "" {
		params["Outcome"] = opts.Outcome
	}
	if opts.Location != "" {
		params["Location"] = opts.Location
	}

	return client.NBASession.NBAGetRequest(endpoints.CumulativeStatsPlayerGames, params, "", nil)
}

// validateCumulativeStatsPlayerGamesParams ensures all input parameters are valid.
func validateCumulativeStatsPlayerGamesParams(opts CumeStatsPlayerGamesOptions) error {
	// Validate LeagueID
	if valid, err := helpers.ValidateLeagueID(opts.LeagueID); !valid {
		return err
	}

	// Validate PlayerID
	if valid, err := helpers.ValidatePlayerID(opts.PlayerID); !valid {
		return err
	}

	// Validate Season
	if valid, err := helpers.ValidateSeason(opts.Season); !valid {
		return err
	}

	// Validate SeasonType
	if valid, err := helpers.ValidateSeasonType(opts.SeasonType); !valid {
		return err
	}

	return nil
}
