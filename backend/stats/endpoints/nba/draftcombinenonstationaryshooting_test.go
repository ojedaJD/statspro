package nba

import (
	"testing"
)

// TestDraftCombineNonStationaryShooting validates correct and incorrect parameters.
func TestDraftCombineNonStationaryShooting(t *testing.T) {
	tests := []struct {
		leagueID   string
		seasonYear string
		expectErr  bool
	}{
		{"00", "2019", false}, // Valid input
		{"10", "2020", false}, // Valid input (WNBA)
		{"99", "2019", true},  // Invalid LeagueID
		{"00", "abcd", true},  // Invalid SeasonYear
		{"", "2019", true},    // Missing LeagueID
		{"00", "", true},      // Missing SeasonYear
	}

	for _, test := range tests {
		_, err := DraftCombineNonStationaryShooting(test.leagueID, test.seasonYear)
		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input: %+v, error: %v", test, err)
			}
		}
	}
}
