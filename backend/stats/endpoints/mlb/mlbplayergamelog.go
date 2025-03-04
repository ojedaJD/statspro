package mlb

import (
	"encoding/json"
	"fmt"
	"sports_api/globals/mlb"
)

type HittingStats struct {
	Date string `json:"date"`
	Game struct {
		Content struct {
			Link string `json:"link"`
		} `json:"content"`
		DayNight   string `json:"dayNight"`
		GameNumber int    `json:"gameNumber"`
		GamePk     int    `json:"gamePk"`
		Link       string `json:"link"`
	} `json:"game"`
	GameType string `json:"gameType"`
	IsHome   bool   `json:"isHome"`
	IsWin    bool   `json:"isWin"`
	League   struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"league"`
	Opponent struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"opponent"`
	PositionsPlayed []struct {
		Abbreviation string `json:"abbreviation"`
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"positionsPlayed"`
	Season string `json:"season"`
	Stat   struct {
		AtBats               int    `json:"atBats"`
		Avg                  string `json:"avg"`
		BaseOnBalls          int    `json:"baseOnBalls"`
		CatchersInterference int    `json:"catchersInterference"`
		Doubles              int    `json:"doubles"`
		Hits                 int    `json:"hits"`
		HomeRuns             int    `json:"homeRuns"`
		IntentionalWalks     int    `json:"intentionalWalks"`
		NumberOfPitches      int    `json:"numberOfPitches"`
		PlateAppearances     int    `json:"plateAppearances"`
		Rbi                  int    `json:"rbi"`
		Runs                 int    `json:"runs"`
		StolenBases          int    `json:"stolenBases"`
		StrikeOuts           int    `json:"strikeOuts"`
		TotalBases           int    `json:"totalBases"`
		Triples              int    `json:"triples"`
	} `json:"stat"`
	Team struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"team"`
}

type PitchingStats struct {
	Date string `json:"date"`
	Game struct {
		Content struct {
			Link string `json:"link"`
		} `json:"content"`
		DayNight   string `json:"dayNight"`
		GameNumber int    `json:"gameNumber"`
		GamePk     int    `json:"gamePk"`
		Link       string `json:"link"`
	} `json:"game"`
	GameType string `json:"gameType"`
	IsHome   bool   `json:"isHome"`
	IsWin    bool   `json:"isWin"`
	League   struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"league"`
	Opponent struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"opponent"`
	Player struct {
		FullName string `json:"fullName"`
		Id       int    `json:"id"`
		Link     string `json:"link"`
	} `json:"player"`
	Season string `json:"season"`
	Sport  struct {
		Abbreviation string `json:"abbreviation"`
		Id           int    `json:"id"`
		Link         string `json:"link"`
	} `json:"sport"`
	Stat struct {
		AirOuts                int    `json:"airOuts"`
		AtBats                 int    `json:"atBats"`
		Avg                    string `json:"avg"`
		Balks                  int    `json:"balks"`
		BaseOnBalls            int    `json:"baseOnBalls"`
		BattersFaced           int    `json:"battersFaced"`
		BlownSaves             int    `json:"blownSaves"`
		CatchersInterference   int    `json:"catchersInterference"`
		CaughtStealing         int    `json:"caughtStealing"`
		CompleteGames          int    `json:"completeGames"`
		Doubles                int    `json:"doubles"`
		EarnedRuns             int    `json:"earnedRuns"`
		Era                    string `json:"era"`
		FlyOuts                int    `json:"flyOuts"`
		GamesFinished          int    `json:"gamesFinished"`
		GamesPitched           int    `json:"gamesPitched"`
		GamesPlayed            int    `json:"gamesPlayed"`
		GamesStarted           int    `json:"gamesStarted"`
		GroundIntoDoublePlay   int    `json:"groundIntoDoublePlay"`
		GroundOuts             int    `json:"groundOuts"`
		GroundOutsToAirouts    string `json:"groundOutsToAirouts"`
		HitBatsmen             int    `json:"hitBatsmen"`
		HitByPitch             int    `json:"hitByPitch"`
		Hits                   int    `json:"hits"`
		HitsPer9Inn            string `json:"hitsPer9Inn"`
		Holds                  int    `json:"holds"`
		HomeRuns               int    `json:"homeRuns"`
		HomeRunsPer9           string `json:"homeRunsPer9"`
		InheritedRunners       int    `json:"inheritedRunners"`
		InheritedRunnersScored int    `json:"inheritedRunnersScored"`
		InningsPitched         string `json:"inningsPitched"`
		IntentionalWalks       int    `json:"intentionalWalks"`
		Losses                 int    `json:"losses"`
		NumberOfPitches        int    `json:"numberOfPitches"`
		Obp                    string `json:"obp"`
		Ops                    string `json:"ops"`
		Outs                   int    `json:"outs"`
		Pickoffs               int    `json:"pickoffs"`
		PitchesPerInning       string `json:"pitchesPerInning"`
		Runs                   int    `json:"runs"`
		RunsScoredPer9         string `json:"runsScoredPer9"`
		SacBunts               int    `json:"sacBunts"`
		SacFlies               int    `json:"sacFlies"`
		SaveOpportunities      int    `json:"saveOpportunities"`
		Saves                  int    `json:"saves"`
		Shutouts               int    `json:"shutouts"`
		Slg                    string `json:"slg"`
		StolenBasePercentage   string `json:"stolenBasePercentage"`
		StolenBases            int    `json:"stolenBases"`
		StrikeOuts             int    `json:"strikeOuts"`
		StrikePercentage       string `json:"strikePercentage"`
		StrikeoutWalkRatio     string `json:"strikeoutWalkRatio"`
		StrikeoutsPer9Inn      string `json:"strikeoutsPer9Inn"`
		Strikes                int    `json:"strikes"`
		Summary                string `json:"summary"`
		TotalBases             int    `json:"totalBases"`
		Triples                int    `json:"triples"`
		WalksPer9Inn           string `json:"walksPer9Inn"`
		Whip                   string `json:"whip"`
		WildPitches            int    `json:"wildPitches"`
		WinPercentage          string `json:"winPercentage"`
		Wins                   int    `json:"wins"`
	} `json:"stat"`
	Team struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"team"`
}

func (player *MLBPlayer) SetStats(stats interface{}) {
	switch v := stats.(type) {
	case []HittingStats:
		player.Hitting = v
	case []PitchingStats:
		player.Pitching = v
	default:
		fmt.Println("Unsupported stats type")
	}
}

func (player *MLBPlayer) GetGameLog(year, gameType string) error {
	fmt.Println(player.Position.Name)

	params := map[string]string{
		"stats":    "gameLog",
		"season":   year,
		"gameType": gameType, // 'S' for Spring Training
		"language": "en",
	}

	url := player.Person.Link + "/stats"

	resp, err := mlb.MLBSession.MLBGetRequest(url, params, "", nil)
	if err != nil {
		return err
	}

	if len(resp.Data.(map[string]interface{})["stats"].([]interface{})) == 0 {

		return fmt.Errorf("no Result")
	}
	i := resp.Data.(map[string]interface{})["stats"].([]interface{})[0]
	displayName := i.(map[string]interface{})["group"].(map[string]interface{})["displayName"].(string)

	var pitching []PitchingStats
	var hitting []HittingStats
	switch displayName {
	case "pitching":
		bytes, err := json.Marshal(i.(map[string]interface{})["splits"])
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &pitching)
		if err != nil {
			return err
		}
		player.SetStats(pitching)
	case "hitting":
		bytes, err := json.Marshal(i.(map[string]interface{})["splits"])
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &hitting)
		if err != nil {
			return err
		}
		player.SetStats(hitting)
	default:
		return fmt.Errorf("unexpected format")

	}

	return nil
}

type GameLogMetaInformation struct {
	Group struct {
		DisplayName string `json:"displayName"`
	} `json:"group"`
}

func (i *GameLogMetaInformation) isNull() bool {
	return i == nil
}
