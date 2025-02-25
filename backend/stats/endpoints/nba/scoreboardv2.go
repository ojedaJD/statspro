package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
	"strconv"
	"time"
)

// ScoreboardV2 retrieves NBA game data for a specific date.
func ScoreboardV2(dayOffset int, gameDate *time.Time, leagueID string) (*client.NBAResponse, error) {
	// Validate required parameters
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return nil, err
	}

	// Format GameDate
	gameDateStr := helpers.FormatDateToString(gameDate)

	params := map[string]string{
		"DayOffset": strconv.Itoa(dayOffset),
		"GameDate":  gameDateStr,
		"LeagueID":  leagueID,
	}

	return client.NBASession.NBAGetRequest(endpoints.ScoreboardV2, params, "", nil)
}

func GetNBAGamesToday() []map[string]interface{} {
	gameDate, _ := time.Parse("2006-01-02", "2025-02-25")
	scoreBoard, err := ScoreboardV2(0, &gameDate, "00")

	dict2, err := scoreBoard.GetNormalizedDict2()
	if err != nil {
		return nil
	}

	return dict2["GameHeader"]

}
