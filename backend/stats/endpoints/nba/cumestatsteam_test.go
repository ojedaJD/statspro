package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCumulativeStatsTeam(t *testing.T) {
	validGameIDs := "0021700807"
	validLeagueID := "00"
	validSeason := "2019-20"
	validSeasonType := "Regular Season"
	validTeamID := "1610612739"
	invalidLeagueID := "99"
	invalidSeason := "19-20"
	invalidSeasonType := "Spring League"

	tests := []struct {
		gameIDs    string
		leagueID   string
		season     string
		seasonType string
		teamID     string
		expectErr  bool
	}{
		{validGameIDs, validLeagueID, validSeason, validSeasonType, validTeamID, false},  // Valid request
		{validGameIDs, invalidLeagueID, validSeason, validSeasonType, validTeamID, true}, // Invalid LeagueID
		{validGameIDs, validLeagueID, invalidSeason, validSeasonType, validTeamID, true}, // Invalid Season format
		{validGameIDs, validLeagueID, validSeason, invalidSeasonType, validTeamID, true}, // Invalid SeasonType
		{"", validLeagueID, validSeason, validSeasonType, validTeamID, true},             // Missing GameIDs
	}

	for _, test := range tests {
		resp, err := CumulativeStatsTeam(test.gameIDs, test.leagueID, test.season, test.seasonType, test.teamID)

		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test)
			} else {
				fmt.Printf("Expected error received for input: %+v\n", test)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v for input: %+v", err, test)
			} else if resp == nil {
				t.Errorf("Received nil response from API for input: %+v", test)
			} else if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected response code: got %d, expected %d for input: %+v", resp.StatusCode, http.StatusOK, test)
			} else {
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test)
			}
		}
	}
}
