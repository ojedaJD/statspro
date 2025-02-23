package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCommonAllPlayers_ActualCall(t *testing.T) {
	tests := []struct {
		isOnlyCurrentSeason int
		leagueID            string
		season              string
		expectErr           bool
	}{
		{1, "00", "2023-24", false}, // Valid request (Current season, NBA)
		{0, "10", "2022-23", false}, // Valid request (All players, WNBA)
		{1, "20", "2021-22", false}, // Valid request (Current season, G-League)
		{2, "00", "2023-24", true},  // Invalid isOnlyCurrentSeason (should be 0 or 1)
		{1, "99", "2023-24", true},  // Invalid leagueID
		{1, "00", "23-24", true},    // Invalid season format
	}

	for _, test := range tests {
		resp, err := CommonAllPlayers(test.isOnlyCurrentSeason, test.leagueID, test.season)

		if test.expectErr {
			// Expecting an error, ensure an error is returned
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test)
			} else {
				fmt.Printf("Expected error received for input: %+v\n", test)
			}
		} else {
			// Expecting success, check the response
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
