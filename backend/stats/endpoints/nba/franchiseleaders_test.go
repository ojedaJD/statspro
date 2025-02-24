package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFranchiseLeaders(t *testing.T) {
	validLeagueID := "00"
	validTeamID := "1610612739"
	invalidLeagueID := "99"
	invalidTeamID := ""

	tests := []struct {
		teamID    string
		leagueID  *string
		expectErr bool
	}{
		{validTeamID, &validLeagueID, false},  // Valid request
		{validTeamID, nil, false},             // Valid without LeagueID
		{validTeamID, &invalidLeagueID, true}, // Invalid LeagueID
		{invalidTeamID, &validLeagueID, true}, // Missing TeamID
	}

	for _, test := range tests {
		resp, err := FranchiseLeaders(test.teamID, test.leagueID)

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
