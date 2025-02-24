package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCumulativeStatsPlayerGames(t *testing.T) {
	validOpts := CumeStatsPlayerGamesOptions{
		LeagueID:   "00",
		PlayerID:   "2544", // LeBron James
		Season:     "2019-20",
		SeasonType: "Regular Season",
	}

	invalidSeasonOpts := CumeStatsPlayerGamesOptions{
		LeagueID:   "00",
		PlayerID:   "2544",
		Season:     "19-20", // Invalid format
		SeasonType: "Regular Season",
	}

	invalidSeasonTypeOpts := CumeStatsPlayerGamesOptions{
		LeagueID:   "00",
		PlayerID:   "2544",
		Season:     "2019-20",
		SeasonType: "InvalidType", // Invalid season type
	}

	invalidLeagueOpts := CumeStatsPlayerGamesOptions{
		LeagueID:   "99", // Invalid LeagueID
		PlayerID:   "2544",
		Season:     "2019-20",
		SeasonType: "Regular Season",
	}

	tests := []struct {
		opts      CumeStatsPlayerGamesOptions
		expectErr bool
	}{
		{validOpts, false},            // Valid request
		{invalidSeasonOpts, true},     // Invalid Season format
		{invalidSeasonTypeOpts, true}, // Invalid SeasonType
		{invalidLeagueOpts, true},     // Invalid LeagueID
	}

	for _, test := range tests {
		resp, err := CumulativeStatsPlayerGames(test.opts)
		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test.opts)
			} else {
				fmt.Printf("Expected error received: %v for input: %+v\n", err, test.opts)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v for input: %+v", err, test.opts)
			} else if resp == nil {
				t.Errorf("Received nil response from API for input: %+v", test.opts)
			} else if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected response code: got %d, expected %d for input: %+v", resp.StatusCode, http.StatusOK, test.opts)
			} else {
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test.opts)
			}
		}
	}
}
