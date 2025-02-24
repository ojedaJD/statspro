package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCommonTeamYears(t *testing.T) {
	validLeagueID := "00"
	invalidLeagueID := "999" // Not a valid two-digit LeagueID

	tests := []struct {
		leagueID  string
		expectErr bool
	}{
		{validLeagueID, false},  // Valid request (NBA)
		{"10", false},           // Valid request (WNBA)
		{invalidLeagueID, true}, // Invalid LeagueID
		{"abc", true},           // Non-numeric LeagueID
	}

	for _, test := range tests {
		resp, err := CommonTeamYears(test.leagueID)

		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test)
			} else {
				fmt.Printf("Expected error received: %v for input: %+v\n", err, test)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v for input: %+v", err, test)
			} else if resp == nil {
				t.Errorf("Received nil response from API for input: %+v", test)
			} else if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected response code: got %d, expected %d for input: %+v", resp.StatusCode, http.StatusOK, test)
			} else {
				fmt.Println(resp.GetNormalizedDict())
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test)
			}
		}
	}
}
