package nba

import (
	"testing"
)

func TestCumulativeStatsPlayer(t *testing.T) {
	validGameID := "0021700807"
	validLeagueID := "00"
	validPlayerID := "2544" // LeBron James
	validSeason := "2019-20"
	validSeasonType := "Regular Season"

	invalidGameID := "123"      // Incorrect format
	invalidLeagueID := "999"    // Invalid LeagueID
	invalidSeason := "19-20"    // Wrong format
	invalidSeasonType := "Test" // Invalid season type

	tests := []struct {
		gameIDs    string
		leagueID   string
		playerID   string
		season     string
		seasonType string
		expectErr  bool
	}{
		{validGameID, validLeagueID, validPlayerID, validSeason, validSeasonType, false},  // Valid request
		{invalidGameID, validLeagueID, validPlayerID, validSeason, validSeasonType, true}, // Invalid GameIDs
		{validGameID, invalidLeagueID, validPlayerID, validSeason, validSeasonType, true}, // Invalid LeagueID
		{validGameID, validLeagueID, validPlayerID, invalidSeason, validSeasonType, true}, // Invalid Season
		{validGameID, validLeagueID, validPlayerID, validSeason, invalidSeasonType, true}, // Invalid SeasonType
	}

	for _, test := range tests {
		_, err := CumulativeStatsPlayer(test.gameIDs, test.leagueID, test.playerID, test.season, test.seasonType)

		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test)
			}
		} else if err != nil {
			t.Errorf("Unexpected error: %v for input: %+v", err, test)
		}
	}
}
