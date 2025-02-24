package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGameRotation(t *testing.T) {
	validGameID := "0021700807"
	validLeagueID := "00"
	invalidGameID := "123"  // Incorrect length
	invalidLeagueID := "99" // Invalid LeagueID
	missingGameID := ""     // Empty GameID

	tests := []struct {
		gameID    string
		leagueID  *string
		expectErr bool
	}{
		{validGameID, &validLeagueID, false},  // Valid request
		{validGameID, nil, false},             // Valid request without LeagueID
		{invalidGameID, &validLeagueID, true}, // Invalid GameID
		{missingGameID, &validLeagueID, true}, // Missing GameID
		{validGameID, &invalidLeagueID, true}, // Invalid LeagueID
	}

	for _, test := range tests {
		resp, err := GameRotation(test.gameID, test.leagueID)

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
