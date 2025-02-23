package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CumulativeStatsTeamGamesOptions defines the query parameters for the API.
type CumulativeStatsTeamGamesOptions struct {
	LeagueID     string
	Season       string
	SeasonType   string
	TeamID       string
	VsTeamID     string
	VsDivision   string
	VsConference string
	SeasonID     string
	Outcome      string
	Location     string
}

// CumulativeStatsTeamGames calls the NBA API and retrieves cumulative team game statistics based on the provided filters.
func CumulativeStatsTeamGames(opts CumulativeStatsTeamGamesOptions) (*client.NBAResponse, error) {
	// Validate input parameters
	if err := validateCumulativeStatsTeamGamesParams(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"LeagueID":   opts.LeagueID,
		"Season":     opts.Season,
		"SeasonType": opts.SeasonType,
		"TeamID":     opts.TeamID,
	}

	// Add optional parameters if provided
	if opts.VsTeamID != "" {
		params["VsTeamID"] = opts.VsTeamID
	}
	if opts.VsDivision != "" {
		params["VsDivision"] = opts.VsDivision
	}
	if opts.VsConference != "" {
		params["VsConference"] = opts.VsConference
	}
	if opts.SeasonID != "" {
		params["SeasonID"] = opts.SeasonID
	}
	if opts.Outcome != "" {
		params["Outcome"] = opts.Outcome
	}
	if opts.Location != "" {
		params["Location"] = opts.Location
	}

	return client.NBASession.NBAGetRequest(endpoints.CumulativeStatsTeamGames, params, "", nil)
}

// validateCumulativeStatsTeamGamesParams ensures all input parameters are valid.
func validateCumulativeStatsTeamGamesParams(opts CumulativeStatsTeamGamesOptions) error {
	// Validate LeagueID
	if valid, err := helpers.ValidateLeagueID(opts.LeagueID); !valid {
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

	// Validate TeamID
	if valid, err := helpers.ValidateTeamID(opts.TeamID); !valid {
		return err
	}

	// Validate optional fields if provided
	if opts.VsTeamID != "" {
		if valid, err := helpers.ValidateTeamID(opts.VsTeamID); !valid {
			return err
		}
	}
	if opts.VsDivision != "" {
		if valid, err := helpers.ValidateDivision(opts.VsDivision); !valid {
			return err
		}
	}
	if opts.VsConference != "" {
		if valid, err := helpers.ValidateConference(opts.VsConference); !valid {
			return err
		}
	}
	if opts.Outcome != "" {
		if valid, err := helpers.ValidateOutcome(opts.Outcome); !valid {
			return err
		}
	}
	if opts.Location != "" {
		if valid, err := helpers.ValidateLocation(opts.Location); !valid {
			return err
		}
	}

	return nil
}
