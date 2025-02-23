package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAssistTracker_ActualCall(t *testing.T) {
	tests := []struct {
		opts      AssistTrackerOptions
		expectErr bool
	}{
		{
			opts: AssistTrackerOptions{
				Season:     "2023-24",
				SeasonType: "Regular Season",
				PerMode:    "PerGame",
				Conference: "East",
				Division:   "Atlantic",
				Location:   "Home",
			},
			expectErr: false, // Valid request
		},
		{
			opts: AssistTrackerOptions{
				Season:     "2023-24",
				SeasonType: "Playoffs",
				PerMode:    "Totals",
				Conference: "West",
				Division:   "Pacific",
			},
			expectErr: false, // Valid request
		},
		{
			opts: AssistTrackerOptions{
				Season:     "2023-24",
				SeasonType: "InvalidSeasonType",
				PerMode:    "PerGame",
			},
			expectErr: true, // Invalid SeasonType
		},
		{
			opts: AssistTrackerOptions{
				Season:     "2023-24",
				SeasonType: "Regular Season",
				PerMode:    "InvalidMode",
			},
			expectErr: true, // Invalid PerMode
		},
		{
			opts: AssistTrackerOptions{
				Season:     "InvalidSeason",
				SeasonType: "Regular Season",
				PerMode:    "PerGame",
			},
			expectErr: true, // Invalid Season format
		},
	}

	for _, test := range tests {
		resp, err := AssistTracker(&test.opts)

		if test.expectErr {
			// If we expect an error, ensure an error is returned
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test.opts)
			} else {
				fmt.Printf("Expected error received for input: %+v\n", test.opts)
			}
		} else {
			// If no error is expected, check the response
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
