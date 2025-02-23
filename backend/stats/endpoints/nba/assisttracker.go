package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// AssistTrackerOptions defines the query parameters for the AssistTracker API.
type AssistTrackerOptions struct {
	Weight           string // Player's weight, typically in pounds.
	VsDivision       string // Division of the opponent team (e.g., "Atlantic", "Central", "Northwest", "Pacific", "Southeast", "Southwest", "East", "West").
	VsConference     string // Conference of the opponent team (e.g., "East" or "West").
	TeamID           string // ID of the team to filter by. If empty, retrieves stats for all teams.
	StarterBench     string // Filter for starters or bench players. Accepts "Starters" or "Bench".
	SeasonType       string // The type of season: "Regular Season", "Pre Season", "Playoffs", or "All Star".
	SeasonSegment    string // Specifies a segment of the season: "Post All-Star" or "Pre All-Star".
	Season           string // The NBA season format as "YYYY-YY" (e.g., "2023-24").
	PlayerPosition   string // Filter by player position: "F", "C", "G", "C-F", "F-C", "F-G", "G-F".
	PlayerExperience string // Filter by experience level: "Rookie", "Sophomore", or "Veteran".
	PerMode          string // Data aggregation method: "Totals" (cumulative stats) or "PerGame" (averaged stats).
	PORound          string // Playoff round filter. Empty for regular season stats.
	Outcome          string // Game outcome filter: "W" (win) or "L" (loss).
	OpponentTeamID   string // ID of the opposing team. If empty, retrieves stats against all teams.
	Month            string // Filter by month (1-12). If empty, retrieves stats for the full season.
	Location         string // Game location filter: "Home" or "Road".
	LeagueID         string // League identifier: "00" (NBA), "10" (WNBA), "20" (G-League).
	LastNGames       string // Filter for last N games (e.g., "5" for last 5 games).
	Height           string // Player height filter in inches (e.g., "6-10" for 6'10").
	GameScope        string // Time-based filter: "Yesterday" or "Last 10".
	DraftYear        string // Filter players by draft year (e.g., "2020").
	DraftPick        string // Filter by draft pick position (e.g., "1" for first pick).
	Division         string // Filter for player's division: "Atlantic", "Central", "Northwest", "Pacific", "Southeast", "Southwest".
	DateTo           string // End date for the data query (format: YYYY-MM-DD).
	DateFrom         string // Start date for the data query (format: YYYY-MM-DD).
	Country          string // Filter players by country of origin (e.g., "USA", "Canada").
	Conference       string // Filter by player's conference: "East" or "West".
	College          string // Filter by college attended (e.g., "Duke", "Kentucky").
}

func AssistTracker(opts *AssistTrackerOptions) (*client.NBAResponse, error) {
	if err := validateAssistTrackerParams(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"Weight":           opts.Weight,
		"VsDivision":       opts.VsDivision,
		"VsConference":     opts.VsConference,
		"TeamID":           opts.TeamID,
		"StarterBench":     opts.StarterBench,
		"SeasonType":       opts.SeasonType,
		"SeasonSegment":    opts.SeasonSegment,
		"Season":           opts.Season,
		"PlayerPosition":   opts.PlayerPosition,
		"PlayerExperience": opts.PlayerExperience,
		"PerMode":          opts.PerMode,
		"PORound":          opts.PORound,
		"Outcome":          opts.Outcome,
		"OpponentTeamID":   opts.OpponentTeamID,
		"Month":            opts.Month,
		"Location":         opts.Location,
		"LeagueID":         opts.LeagueID,
		"LastNGames":       opts.LastNGames,
		"Height":           opts.Height,
		"GameScope":        opts.GameScope,
		"DraftYear":        opts.DraftYear,
		"DraftPick":        opts.DraftPick,
		"Division":         opts.Division,
		"DateTo":           opts.DateTo,
		"DateFrom":         opts.DateFrom,
		"Country":          opts.Country,
		"Conference":       opts.Conference,
		"College":          opts.College,
	}

	return client.NBASession.NBAGetRequest(endpoints.AssistTracker, params, "", nil)
}

func validateAssistTrackerParams(opts *AssistTrackerOptions) error {
	if opts.SeasonType != "" {
		if valid, err := helpers.ValidateSeasonType(opts.SeasonType); !valid {
			return err
		}
	}

	if opts.PerMode != "" {
		if valid, err := helpers.ValidatePerMode(opts.PerMode); !valid {
			return err
		}
	}

	if opts.Location != "" {
		if valid, err := helpers.ValidateLocation(opts.Location); !valid {
			return err
		}
	}

	if opts.Conference != "" {
		if valid, err := helpers.ValidateConference(opts.Conference); !valid {
			return err
		}
	}

	if opts.VsConference != "" {
		if valid, err := helpers.ValidateConference(opts.VsConference); !valid {
			return err
		}
	}

	if opts.Division != "" {
		if valid, err := helpers.ValidateDivision(opts.Division); !valid {
			return err
		}
	}

	if opts.VsDivision != "" {
		if valid, err := helpers.ValidateDivision(opts.VsDivision); !valid {
			return err
		}
	}

	if opts.StarterBench != "" {
		if valid, err := helpers.ValidateStarterBench(opts.StarterBench); !valid {
			return err
		}
	}

	if opts.Outcome != "" {
		if valid, err := helpers.ValidateOutcome(opts.Outcome); !valid {
			return err
		}
	}

	if opts.SeasonSegment != "" {
		if valid, err := helpers.ValidateSeasonSegment(opts.SeasonSegment); !valid {
			return err
		}
	}

	if opts.PlayerPosition != "" {
		if valid, err := helpers.ValidatePlayerPosition(opts.PlayerPosition); !valid {
			return err
		}
	}

	if opts.PlayerExperience != "" {
		if valid, err := helpers.ValidatePlayerExperience(opts.PlayerExperience); !valid {
			return err
		}
	}

	if opts.GameScope != "" {
		if valid, err := helpers.ValidateGameScope(opts.GameScope); !valid {
			return err
		}
	}

	// No validation needed for fields that accept any string: Weight, TeamID, Season, PORound, OpponentTeamID, Month,
	// LeagueID, LastNGames, Height, DraftYear, DraftPick, DateTo, DateFrom, Country, College.

	return nil
}
