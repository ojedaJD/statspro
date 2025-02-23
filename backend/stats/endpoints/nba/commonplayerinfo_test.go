package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCommonPlayerInfo_ActualCall(t *testing.T) {
	leagueNBA := "00"
	leagueWNBA := "10"
	invalidLeague := "99"

	tests := []struct {
		playerID  string
		leagueID  *string
		expectErr bool
	}{
		{"2544", &leagueNBA, false},     // Valid request (LeBron James, NBA)
		{"1629027", &leagueWNBA, false}, // Valid request (WNBA player)
		{"2544", nil, false},            // Valid request with no LeagueID
		{"", &leagueNBA, true},          // Missing playerID
		{"2544", &invalidLeague, true},  // Invalid LeagueID
	}

	for _, test := range tests {
		resp, err := CommonPlayerInfo(test.playerID, test.leagueID)

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
