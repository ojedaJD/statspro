package nba

import (
	"errors"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
	"strconv"
)

// DraftHistoryOptions defines the query parameters for retrieving NBA draft history.
type DraftHistoryOptions struct {
	LeagueID    string // The league ID (e.g., "00" for NBA, "10" for WNBA, "20" for G-League).
	TopX        int    // The number of top draft picks to retrieve (must be a non-negative integer).
	TeamID      string // The ID of the team to filter the draft history (optional).
	Season      string // The draft season in "YYYY" format (e.g., "2019") (optional).
	RoundPick   int    // The specific pick within a round (must be a non-negative integer) (optional).
	RoundNum    int    // The round number in the draft (must be a non-negative integer) (optional).
	OverallPick int    // The overall draft pick number (must be a non-negative integer) (optional).
	College     int    // The ID of the college associated with the player (must be a non-negative integer) (optional).
}

// ValidateDraftHistory checks for negative values in DraftHistoryOptions.
func (opts *DraftHistoryOptions) ValidateDraftHistory() error {
	if _, err := helpers.ValidateLeagueID(opts.LeagueID); err != nil {
		return err
	}

	if opts.Season != "" {
		if _, err := helpers.ValidateSeasonYear(opts.Season); err != nil {
			return err
		}
	}

	// Ensure no negative values for numerical fields
	if opts.TopX < 0 {
		return errors.New("TopX must be a non-negative integer")
	}
	if opts.RoundPick < 0 {
		return errors.New("RoundPick must be a non-negative integer")
	}
	if opts.RoundNum < 0 {
		return errors.New("RoundNum must be a non-negative integer")
	}
	if opts.OverallPick < 0 {
		return errors.New("OverallPick must be a non-negative integer")
	}
	if opts.College < 0 {
		return errors.New("college must be a non-negative integer")
	}

	return nil
}

// DraftHistory retrieves NBA draft history based on the provided filters.
func DraftHistory(opts DraftHistoryOptions) (*client.NBAResponse, error) {
	if err := opts.ValidateDraftHistory(); err != nil {
		return nil, err
	}

	if opts.Season != "" {
		if _, err := helpers.ValidateSeasonYear(opts.Season); err != nil {
			return nil, err
		}
	}

	params := map[string]string{
		"LeagueID": opts.LeagueID,
	}

	if opts.TopX > 0 {
		params["TopX"] = strconv.Itoa(opts.TopX)
	}
	if opts.TeamID != "" {
		params["TeamID"] = opts.TeamID
	}
	if opts.Season != "" {
		params["Season"] = opts.Season
	}
	if opts.RoundPick > 0 {
		params["RoundPick"] = strconv.Itoa(opts.RoundPick)
	}
	if opts.RoundNum > 0 {
		params["RoundNum"] = strconv.Itoa(opts.RoundNum)
	}
	if opts.OverallPick > 0 {
		params["OverallPick"] = strconv.Itoa(opts.OverallPick)
	}
	if opts.College > 0 {
		params["College"] = strconv.Itoa(opts.College)
	}

	return client.NBASession.NBAGetRequest(endpoints.DraftHistory, params, "", nil)

}
