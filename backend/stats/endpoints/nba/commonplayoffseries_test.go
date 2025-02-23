package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCommonPlayoffSeries_ActualCall(t *testing.T) {
	leagueNBA := "00"
	validSeason := "2019-20"
	invalidSeason := "19-20"
	seriesID := "0041900401"
	invalidLeague := "99"

	tests := []struct {
		leagueID  string
		season    string
		seriesID  *string
		expectErr bool
	}{
		{leagueNBA, validSeason, &seriesID, true}, // Valid request (NBA, 2019-20, specific series)
		{leagueNBA, validSeason, nil, false},      // Valid request (NBA, 2019-20, all series)
		{"10", validSeason, nil, false},           // Valid request (WNBA, 2019-20)
		{invalidLeague, validSeason, nil, true},   // Invalid LeagueID
		{leagueNBA, invalidSeason, nil, true},     // Invalid Season format
	}

	for _, test := range tests {
		resp, err := CommonPlayoffSeries(test.leagueID, test.season, test.seriesID)

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
