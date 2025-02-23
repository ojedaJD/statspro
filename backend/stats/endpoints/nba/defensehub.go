package nba

import (
	"fmt"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// DefenseHubOptions defines query parameters for the DefenseHub API.
type DefenseHubOptions struct {
	GameScope    string
	LeagueID     string
	PlayerOrTeam string
	PlayerScope  string
	Season       string
	SeasonType   string
}

// Deprecated: to be implemented
// DefenseHub calls the NBA API and retrieves defensive statistics.
//
// Example Usage:
//
//	data, err := DefenseHub(DefenseHubOptions{
//	    GameScope:   "Season",
//	    LeagueID:    "00",
//	    PlayerOrTeam: "Team",
//	    PlayerScope: "All Players",
//	    Season:      "2019-20",
//	    SeasonType:  "Regular Season",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(data)
func DefenseHub(opts DefenseHubOptions) (*client.NBAResponse, error) {
	// Validate parameters
	if err := validateDefenseHubParams(opts); err != nil {
		return nil, err
	}

	// Construct query parameters
	params := map[string]string{
		"GameScope":    opts.GameScope,
		"LeagueID":     opts.LeagueID,
		"PlayerOrTeam": opts.PlayerOrTeam,
		"PlayerScope":  opts.PlayerScope,
		"Season":       opts.Season,
		"SeasonType":   opts.SeasonType,
	}
	request, err := client.NBASession.NBAGetRequest(endpoints.DefenseHub, params, "", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(request.ParseResponse())
	// Make API request
	return client.NBASession.NBAGetRequest(endpoints.DefenseHub, params, "", nil)
}

// validateDefenseHubParams ensures all input parameters are valid.
func validateDefenseHubParams(opts DefenseHubOptions) error {
	// Validate GameScope
	if _, err := helpers.ValidateGameScope(opts.GameScope); err != nil {
		return err
	}

	// Validate LeagueID
	if _, err := helpers.ValidateLeagueID(opts.LeagueID); err != nil {
		return err
	}

	// Validate PlayerOrTeam
	if _, err := helpers.ValidatePlayerOrTeam(opts.PlayerOrTeam); err != nil {
		return err
	}

	// Validate PlayerScope
	if _, err := helpers.ValidatePlayerScope(opts.PlayerScope); err != nil {
		return err
	}

	// Validate Season format
	if _, err := helpers.ValidateSeason(opts.Season); err != nil {
		return err
	}

	// Validate SeasonType
	if _, err := helpers.ValidateSeasonType(opts.SeasonType); err != nil {
		return err
	}

	return nil
}
