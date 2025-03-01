package nba

import (
	"sort"
	"sync"
	"time"
)

type NBABaseGameLog struct {
	AST              int     `json:"AST"`
	AVAILABLEFLAG    int     `json:"AVAILABLE_FLAG"`
	BLK              int     `json:"BLK"`
	BLKA             int     `json:"BLKA"`
	DD2              int     `json:"DD2"`
	DREB             int     `json:"DREB"`
	FG3A             int     `json:"FG3A"`
	FG3M             int     `json:"FG3M"`
	FG3PCT           float64 `json:"FG3_PCT"`
	FGA              int     `json:"FGA"`
	FGM              int     `json:"FGM"`
	FGPCT            float64 `json:"FG_PCT"`
	FTA              int     `json:"FTA"`
	FTM              int     `json:"FTM"`
	FTPCT            float64 `json:"FT_PCT"`
	GAMEDATE         string  `json:"GAME_DATE"`
	GAMEID           string  `json:"GAME_ID"`
	MATCHUP          string  `json:"MATCHUP"`
	MIN              float64 `json:"MIN"`
	MINSEC           string  `json:"MIN_SEC"`
	NBAFANTASYPTS    float64 `json:"NBA_FANTASY_PTS"`
	NICKNAME         string  `json:"NICKNAME"`
	OREB             int     `json:"OREB"`
	PF               int     `json:"PF"`
	PFD              int     `json:"PFD"`
	PLAYERID         int     `json:"PLAYER_ID"`
	PLAYERNAME       string  `json:"PLAYER_NAME"`
	PLUSMINUS        int     `json:"PLUS_MINUS"`
	PTS              int     `json:"PTS"`
	REB              int     `json:"REB"`
	SEASONYEAR       string  `json:"SEASON_YEAR"`
	STL              int     `json:"STL"`
	TD3              int     `json:"TD3"`
	TEAMABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAMID           int     `json:"TEAM_ID"`
	TEAMNAME         string  `json:"TEAM_NAME"`
	TOV              int     `json:"TOV"`
	WL               string  `json:"WL"`
}

type BaseGameLogSlice []NBABaseGameLog

func (receiver BaseGameLogSlice) GetPlayerGameLog(playerID int) BaseGameLogSlice {
	var wg sync.WaitGroup
	gameLogChannel := make(chan NBABaseGameLog, len(receiver)) // Buffered channel to collect results

	// Goroutine to filter logs
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, log := range receiver {
			if log.PLAYERID == playerID {
				gameLogChannel <- log
			}
		}
		close(gameLogChannel) // Close the channel after filtering
	}()

	wg.Wait() // Wait for all goroutines to finish

	// Collect results from the channel
	var filteredLogs []NBABaseGameLog
	for log := range gameLogChannel {
		filteredLogs = append(filteredLogs, log)
	}
	// Sort by GAME_DATE (newest first, oldest last)
	sort.Slice(filteredLogs, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02", filteredLogs[i].GAMEDATE) // YYYY-MM-DD format
		dateJ, _ := time.Parse("2006-01-02", filteredLogs[j].GAMEDATE)
		return dateI.Before(dateJ) // Sort oldest first
	})

	return filteredLogs
}
