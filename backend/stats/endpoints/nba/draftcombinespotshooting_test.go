package nba

import (
	"fmt"
	"testing"
)

// TestDraftCombineSpotShooting validates correct and incorrect parameters.
func TestDraftCombineSpotShooting(t *testing.T) {
	tests := []struct {
		leagueID   string
		seasonYear string
		expectErr  bool
	}{
		{"00", "2019", false},
		{"10", "2020", false},
		{"99", "2019", true},
		{"00", "abcd", true},
		{"", "2019", true},
		{"00", "", true},
	}

	for _, test := range tests {
		resp, err := DraftCombineSpotShooting(test.leagueID, test.seasonYear)
		fmt.Println(resp.GetNormalizedDict())
		fmt.Println(resp.GetNormalizedDict())
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
