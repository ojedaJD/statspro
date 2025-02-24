package nba

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestScoreboardV2(t *testing.T) {
	validLeagueID := "00"
	validDayOffset := -19

	gameDate1, _ := time.Parse("2006-01-02", "2023-01-01")
	gameDate2, _ := time.Parse("2006-01-02", "2023-12-31")
	gameDate3, _ := time.Parse("2006-01-02", "2022-06-15")
	gameDate4, _ := time.Parse("2006-01-02", "2020-08-16")
	gameDate5, _ := time.Parse("2006-01-02", "2018-02-10")

	// 5 valid test cases
	tests := []struct {
		dayOffset int
		gameDate  *time.Time
		leagueID  string
	}{
		{validDayOffset, &gameDate1, validLeagueID}, // Valid case 1
		{validDayOffset, &gameDate2, validLeagueID}, // Valid case 2
		{validDayOffset, &gameDate3, validLeagueID}, // Valid case 3
		{validDayOffset, &gameDate4, validLeagueID}, // Valid case 4
		{validDayOffset, &gameDate5, validLeagueID}, // Valid case 5
	}

	for _, test := range tests {
		resp, err := ScoreboardV2(test.dayOffset, test.gameDate, test.leagueID)

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
