package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// DraftCombineStats retrieves draft combine statistics for NBA draft prospects.
func DraftCombineStats(leagueID, seasonYear string) (*client.NBAResponse, error) {
	// Validate input parameters
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return nil, err
	}
	if valid, err := helpers.ValidateSeasonYearOrAllTime(seasonYear); !valid {
		return nil, err
	}

	params := map[string]string{
		"LeagueID":   leagueID,
		"SeasonYear": seasonYear,
	}

	return client.NBASession.NBAGetRequest(endpoints.DraftCombineStats, params, "", nil)
}
