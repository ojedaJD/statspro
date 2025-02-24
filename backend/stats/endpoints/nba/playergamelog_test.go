package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPlayerGameLog(t *testing.T) {
	validPlayerID := "2544"
	validSeason := "2022-23"
	validSeasonType := "Regular Season"
	validLeagueID := "00"

	// Successful test cases
	var tests = []struct {
		playerID   string
		season     string
		seasonType string
		leagueID   *string
	}{
		{validPlayerID, validSeason, validSeasonType, &validLeagueID}, // Valid with all parameters
		{validPlayerID, validSeason, validSeasonType, nil},            // Valid without optional fields
		{validPlayerID, validSeason, validSeasonType, &validLeagueID}, // Valid with LeagueID only
		{validPlayerID, validSeason, validSeasonType, nil},            // Valid with DateFrom and DateTo only
		{validPlayerID, validSeason, validSeasonType, &validLeagueID}, // Valid with DateFrom only
	}
	for i, test := range tests {
		fmt.Println(i)
		resp, err := PlayerGameLog(test.playerID, test.season, test.seasonType, test.leagueID)
		if resp != nil {
			fmt.Println(resp.GetNormalizedDict())
		}
		if err != nil {
			t.Errorf("Unexpected error: %v for input: %+v", err, test)
		} else if resp == nil {
			t.Errorf("Received nil response from API for input: %+v", test)
		} else if resp.StatusCode == http.StatusOK {
			fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test)
		} else {
			t.Errorf("Unexpected response code: got %d, expected %d for input: %+v", resp.StatusCode, http.StatusOK, test)
		}
	}
}
