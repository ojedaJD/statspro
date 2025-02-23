package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCommonTeamRoster(t *testing.T) {
	leagueNBA := "00"
	validSeason := "2019-20"
	invalidSeason := "19-20"
	teamID := "1610612739" // Cleveland Cavaliers
	invalidTeamID := ""

	tests := []struct {
		teamID    string
		season    string
		leagueID  *string
		expectErr bool
	}{
		{teamID, validSeason, &leagueNBA, false},       // Valid request (NBA, 2019-20, valid team)
		{teamID, validSeason, nil, false},              // Valid request without LeagueID
		{invalidTeamID, validSeason, &leagueNBA, true}, // Invalid TeamID
		{teamID, invalidSeason, &leagueNBA, true},      // Invalid Season format
		{teamID, validSeason, nil, false},              // Valid request without LeagueID
	}

	for _, test := range tests {
		resp, err := CommonTeamRoster(test.teamID, test.season, test.leagueID)

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
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test)
			}
		}
	}
}
