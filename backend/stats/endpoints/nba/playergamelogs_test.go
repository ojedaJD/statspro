package nba

import (
	"fmt"
	"testing"
)

type BaseStatGameLog struct {
	AST                int     `json:"AST"`
	ASTRANK            int     `json:"AST_RANK"`
	AVAILABLEFLAG      int     `json:"AVAILABLE_FLAG"`
	BLK                int     `json:"BLK"`
	BLKA               int     `json:"BLKA"`
	BLKARANK           int     `json:"BLKA_RANK"`
	BLKRANK            int     `json:"BLK_RANK"`
	DD2                int     `json:"DD2"`
	DD2RANK            int     `json:"DD2_RANK"`
	DREB               int     `json:"DREB"`
	DREBRANK           int     `json:"DREB_RANK"`
	FG3A               int     `json:"FG3A"`
	FG3ARANK           int     `json:"FG3A_RANK"`
	FG3M               int     `json:"FG3M"`
	FG3MRANK           int     `json:"FG3M_RANK"`
	FG3PCT             float64 `json:"FG3_PCT"`
	FG3PCTRANK         int     `json:"FG3_PCT_RANK"`
	FGA                int     `json:"FGA"`
	FGARANK            int     `json:"FGA_RANK"`
	FGM                int     `json:"FGM"`
	FGMRANK            int     `json:"FGM_RANK"`
	FGPCT              float64 `json:"FG_PCT"`
	FGPCTRANK          int     `json:"FG_PCT_RANK"`
	FTA                int     `json:"FTA"`
	FTARANK            int     `json:"FTA_RANK"`
	FTM                int     `json:"FTM"`
	FTMRANK            int     `json:"FTM_RANK"`
	FTPCT              int     `json:"FT_PCT"`
	FTPCTRANK          int     `json:"FT_PCT_RANK"`
	GAMEDATE           string  `json:"GAME_DATE"`
	GAMEID             string  `json:"GAME_ID"`
	GPRANK             int     `json:"GP_RANK"`
	LRANK              int     `json:"L_RANK"`
	MATCHUP            string  `json:"MATCHUP"`
	MIN                float64 `json:"MIN"`
	MINRANK            int     `json:"MIN_RANK"`
	MINSEC             string  `json:"MIN_SEC"`
	NBAFANTASYPTS      float64 `json:"NBA_FANTASY_PTS"`
	NBAFANTASYPTSRANK  int     `json:"NBA_FANTASY_PTS_RANK"`
	NICKNAME           string  `json:"NICKNAME"`
	OREB               int     `json:"OREB"`
	OREBRANK           int     `json:"OREB_RANK"`
	PF                 int     `json:"PF"`
	PFD                int     `json:"PFD"`
	PFDRANK            int     `json:"PFD_RANK"`
	PFRANK             int     `json:"PF_RANK"`
	PLAYERID           int     `json:"PLAYER_ID"`
	PLAYERNAME         string  `json:"PLAYER_NAME"`
	PLUSMINUS          int     `json:"PLUS_MINUS"`
	PLUSMINUSRANK      int     `json:"PLUS_MINUS_RANK"`
	PTS                int     `json:"PTS"`
	PTSRANK            int     `json:"PTS_RANK"`
	REB                int     `json:"REB"`
	REBRANK            int     `json:"REB_RANK"`
	SEASONYEAR         string  `json:"SEASON_YEAR"`
	STL                int     `json:"STL"`
	STLRANK            int     `json:"STL_RANK"`
	TD3                int     `json:"TD3"`
	TD3RANK            int     `json:"TD3_RANK"`
	TEAMABBREVIATION   string  `json:"TEAM_ABBREVIATION"`
	TEAMID             int     `json:"TEAM_ID"`
	TEAMNAME           string  `json:"TEAM_NAME"`
	TOV                int     `json:"TOV"`
	TOVRANK            int     `json:"TOV_RANK"`
	WL                 string  `json:"WL"`
	WNBAFANTASYPTS     int     `json:"WNBA_FANTASY_PTS"`
	WNBAFANTASYPTSRANK int     `json:"WNBA_FANTASY_PTS_RANK"`
	WPCTRANK           int     `json:"W_PCT_RANK"`
	WRANK              int     `json:"W_RANK"`
}

func TestGetAllNBAPlayerStats(t *testing.T) {
	response := GetNBAPlayerStatsSecondHalf()

	if response == nil {
		t.Fatal("Expected response, got nil")
	}

	// Additional checks depending on response structure
	if len(response) == 0 {
		t.Errorf("Expected non-empty data in response, got empty")
	}

	fmt.Println(response.GetPlayerGameLog(2544))

}
