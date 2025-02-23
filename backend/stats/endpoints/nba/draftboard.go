package nba

import (
	"errors"
	"regexp"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// DraftBoardOptions defines query parameters for the DraftBoard API.
type DraftBoardOptions struct {
	LeagueID    string
	Season      string
	TopX        *string
	TeamID      *string
	RoundPick   *string
	RoundNum    *string
	OverallPick *string
	College     *string
}

// DraftBoard calls the NBA API and retrieves draft board information.
//
// Example Usage:
//
//	data, err := DraftBoard(DraftBoardOptions{
//	    LeagueID: "00",
//	    Season:   "2019",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(data)
func DraftBoard(opts DraftBoardOptions) (*client.NBAResponse, error) {
	// Validate parameters
	if err := validateDraftBoardParams(opts); err != nil {
		return nil, err
	}

	// Construct query parameters
	params := map[string]string{
		"LeagueID": opts.LeagueID,
		"Season":   opts.Season,
	}

	// Include nullable parameters only if they are set
	if opts.TopX != nil {
		params["TopX"] = *opts.TopX
	}
	if opts.TeamID != nil {
		params["TeamID"] = *opts.TeamID
	}
	if opts.RoundPick != nil {
		params["RoundPick"] = *opts.RoundPick
	}
	if opts.RoundNum != nil {
		params["RoundNum"] = *opts.RoundNum
	}
	if opts.OverallPick != nil {
		params["OverallPick"] = *opts.OverallPick
	}
	if opts.College != nil {
		params["College"] = *opts.College
	}

	// Make API request
	return client.NBASession.NBAGetRequest(endpoints.DraftBoard, params, "", nil)
}

// validateDraftBoardParams ensures all input parameters are valid.
func validateDraftBoardParams(opts DraftBoardOptions) error {
	// Validate LeagueID
	if _, err := helpers.ValidateLeagueID(opts.LeagueID); err != nil {
		return err
	}

	// Validate Season format (must be "YYYY")
	seasonPattern := `^\d{4}$`
	if matched, err := regexp.MatchString(seasonPattern, opts.Season); err != nil || !matched {
		return errors.New("invalid Season format: must be 'YYYY' (e.g., '2019')")
	}

	return nil
}
