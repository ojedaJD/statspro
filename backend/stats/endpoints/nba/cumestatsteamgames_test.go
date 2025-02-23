package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCumulativeStatsTeamGames_ActualCall(t *testing.T) {
	validOpts := CumulativeStatsTeamGamesOptions{
		LeagueID:   "00",
		Season:     "2020-21",
		SeasonType: "Playoffs",
		TeamID:     "1610612744",
	}
	invalidOpts := CumulativeStatsTeamGamesOptions{
		LeagueID:   "99", // Invalid LeagueID
		Season:     "2019-20",
		SeasonType: "Regular Season",
		TeamID:     "1610612739",
	}

	tests := []struct {
		opts      CumulativeStatsTeamGamesOptions
		expectErr bool
	}{
		{validOpts, false},  // Valid request
		{invalidOpts, true}, // Invalid LeagueID
	}

	for _, test := range tests {
		resp, err := CumulativeStatsTeamGames(test.opts)

		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test.opts)
			} else {
				fmt.Printf("Expected error received for input: %+v\n", test.opts)
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
