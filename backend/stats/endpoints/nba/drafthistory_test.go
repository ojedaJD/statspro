package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDraftHistory(t *testing.T) {
	validLeagueID := "00"
	validSeason := "2019"
	validTopX := 10
	validRoundPick := 5
	validRoundNum := 2
	validOverallPick := 1
	validCollege := 1001
	validTeamID := "1610612739"

	invalidLeagueID := "99"
	invalidSeason := "19"
	invalidTopX := -10
	invalidRoundPick := -5
	invalidRoundNum := -2
	invalidOverallPick := -1
	invalidCollege := -1001

	tests := []struct {
		opts      DraftHistoryOptions
		expectErr bool
	}{
		{DraftHistoryOptions{validLeagueID, validTopX, validTeamID, validSeason, validRoundPick, validRoundNum, validOverallPick, validCollege}, false},  // Valid request
		{DraftHistoryOptions{invalidLeagueID, validTopX, validTeamID, validSeason, validRoundPick, validRoundNum, validOverallPick, validCollege}, true}, // Invalid LeagueID
		{DraftHistoryOptions{validLeagueID, validTopX, validTeamID, invalidSeason, validRoundPick, validRoundNum, validOverallPick, validCollege}, true}, // Invalid Season
		{DraftHistoryOptions{validLeagueID, invalidTopX, validTeamID, validSeason, validRoundPick, validRoundNum, validOverallPick, validCollege}, true}, // Invalid TopX
		{DraftHistoryOptions{validLeagueID, validTopX, validTeamID, validSeason, invalidRoundPick, validRoundNum, validOverallPick, validCollege}, true}, // Invalid RoundPick
		{DraftHistoryOptions{validLeagueID, validTopX, validTeamID, validSeason, validRoundPick, invalidRoundNum, validOverallPick, validCollege}, true}, // Invalid RoundNum
		{DraftHistoryOptions{validLeagueID, validTopX, validTeamID, validSeason, validRoundPick, validRoundNum, invalidOverallPick, validCollege}, true}, // Invalid OverallPick
		{DraftHistoryOptions{validLeagueID, validTopX, validTeamID, validSeason, validRoundPick, validRoundNum, validOverallPick, invalidCollege}, true}, // Invalid College
	}

	for i, test := range tests {
		resp, err := DraftHistory(test.opts)
		fmt.Println(i)
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
