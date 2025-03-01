package nba

import (
	"encoding/json"
	"errors"
	"fmt"

	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// PlayerGameLogsOptions defines all the query parameters that can be used
// to request a player's game logs from the NBA Stats API.
type PlayerGameLogsOptions struct {
	VsDivision     string
	VsConference   string
	TeamID         int
	ShotClockRange string
	SeasonType     string
	SeasonSegment  string
	Season         string
	PlayerID       int
	Period         int
	PerMode        string
	PORound        int
	Outcome        string
	OpposingTeamID int
	Month          int
	MeasureType    string
	Location       string
	LeagueID       string
	LastNGames     int
	GameSegment    string
	DateTo         string
	DateFrom       string
}

// PlayerGameLogs retrieves a player's game log statistics from the NBA Stats API
func PlayerGameLogs(opts *PlayerGameLogsOptions) (*client.NBAResponse, error) {
	if opts == nil {
		return nil, errors.New("options must not be nil")
	}

	params := map[string]string{
		"PlayerID":       helpers.IntToString(opts.PlayerID),
		"Season":         opts.Season,
		"SeasonType":     opts.SeasonType,
		"VsDivision":     opts.VsDivision,
		"VsConference":   opts.VsConference,
		"TeamID":         helpers.IntToString(opts.TeamID),
		"ShotClockRange": opts.ShotClockRange,
		"SeasonSegment":  opts.SeasonSegment,
		"Period":         helpers.IntToString(opts.Period),
		"PerMode":        opts.PerMode,
		"PORound":        helpers.IntToString(opts.PORound),
		"Outcome":        opts.Outcome,
		"OpponentTeamID": helpers.IntToString(opts.OpposingTeamID),
		"Month":          helpers.IntToString(opts.Month),
		"MeasureType":    opts.MeasureType,
		"Location":       opts.Location,
		"LeagueID":       opts.LeagueID,
		"LastNGames":     helpers.IntToString(opts.LastNGames),
		"GameSegment":    opts.GameSegment,
		"DateTo":         opts.DateTo,
		"DateFrom":       opts.DateFrom,
	}

	return client.NBASession.NBAGetRequest(endpoints.PlayerGameLogs, params, "", nil)
}

// getNBAPlayerStats is a helper function to get game logs based on GameSegment or Period.
func getNBAPlayerStats(gameSegment string, period int) BaseGameLogSlice {
	t, e := PlayerGameLogs(&PlayerGameLogsOptions{
		MeasureType:    "Base",
		PerMode:        "Totals",
		LeagueID:       "00",
		Season:         "2024-25",
		SeasonType:     "Regular Season",
		PORound:        0,
		TeamID:         0,
		PlayerID:       0,
		Outcome:        "",
		Location:       "",
		Month:          0,
		SeasonSegment:  "",
		DateFrom:       "",
		DateTo:         "",
		OpposingTeamID: 0,
		VsConference:   "",
		VsDivision:     "",
		GameSegment:    gameSegment,
		Period:         period,
		ShotClockRange: "",
		LastNGames:     0,
	})
	if e != nil {
		return nil
	}

	dict2, err := t.GetNormalizedDict2()
	if err != nil {
		return nil
	}

	var GameLogs []NBABaseGameLog

	marshal, err := json.Marshal(dict2["PlayerGameLogs"])
	if err != nil {
		return nil
	}
	err = json.Unmarshal(marshal, &GameLogs)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return GameLogs
}

// GetAllNBAPlayerStatsFullSeason retrieves full game stats for all players in the current season.
func GetAllNBAPlayerStatsFullSeason() BaseGameLogSlice {
	return getNBAPlayerStats("", 0)
}

// GetNBAPlayerStatsByQuarter retrieves game stats for a specific quarter.
func GetNBAPlayerStatsByQuarter(quarter int) BaseGameLogSlice {
	if quarter < 1 || quarter > 4 {
		fmt.Println("Invalid quarter. Please enter a value between 1 and 4.")
		return nil
	}
	return getNBAPlayerStats("", quarter)
}

// GetNBAPlayerStatsFirstHalf retrieves game stats for the first half.
func GetNBAPlayerStatsFirstHalf() BaseGameLogSlice {
	return getNBAPlayerStats("First Half", 0)
}

// GetNBAPlayerStatsSecondHalf retrieves game stats for the second half.
func GetNBAPlayerStatsSecondHalf() BaseGameLogSlice {
	return getNBAPlayerStats("Second Half", 0)
}

// GetStats retrieves NBA player statistics based on the input key
func GetCurrentSeasonStats(period string) BaseGameLogSlice {
	switch period {
	case "Season":
		return GetAllNBAPlayerStatsFullSeason()
	case "1Q":
		return GetNBAPlayerStatsByQuarter(1)
	case "2Q":
		return GetNBAPlayerStatsByQuarter(2)
	case "3Q":
		return GetNBAPlayerStatsByQuarter(3)
	case "4Q":
		return GetNBAPlayerStatsByQuarter(4)
	case "1H":
		return GetNBAPlayerStatsFirstHalf()
	case "2H":
		return GetNBAPlayerStatsSecondHalf()
	default:
		fmt.Println("Invalid key. Use 'full', '1Q', '2Q', '3Q', '4Q', '1H', or '2H'.")
		return nil
	}
}
