package nba

import (
	"fmt"
	"testing"
)

// TestDraftBoard tests the DraftBoard API function.
func TestDraftBoard(t *testing.T) {
	validLeague := "00"
	validSeason := "2019"
	invalidLeague := "99"
	invalidSeason := "19"

	topX := "10"
	teamID := "1610612739"
	roundPick := "5"
	roundNum := "2"
	overallPick := "15"
	college := "Duke"

	tests := []struct {
		opts      DraftBoardOptions
		expectErr bool
	}{
		// Valid cases
		{DraftBoardOptions{LeagueID: validLeague, Season: validSeason}, false},
		{DraftBoardOptions{LeagueID: validLeague, Season: validSeason, TopX: &topX}, false},
		{DraftBoardOptions{LeagueID: validLeague, Season: validSeason, TeamID: &teamID}, false},
		{DraftBoardOptions{LeagueID: validLeague, Season: validSeason, RoundPick: &roundPick}, false},
		{DraftBoardOptions{LeagueID: validLeague, Season: validSeason, RoundNum: &roundNum}, false},
		{DraftBoardOptions{LeagueID: validLeague, Season: validSeason, OverallPick: &overallPick}, false},
		{DraftBoardOptions{LeagueID: validLeague, Season: validSeason, College: &college}, false},

		// Invalid cases
		{DraftBoardOptions{LeagueID: invalidLeague, Season: validSeason}, true}, // Invalid LeagueID
		{DraftBoardOptions{LeagueID: validLeague, Season: invalidSeason}, true}, // Invalid Season
		{DraftBoardOptions{LeagueID: "", Season: validSeason}, true},            // Missing LeagueID
		{DraftBoardOptions{LeagueID: validLeague, Season: ""}, true},            // Missing Season
	}

	for _, test := range tests {
		resp, err := DraftBoard(test.opts)
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
			} else {
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test.opts)
			}
		}
	}
}
