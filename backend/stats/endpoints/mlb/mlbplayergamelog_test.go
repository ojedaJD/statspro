package mlb

import (
	"fmt"
	"testing"
)

func TestSpringTraingGameLog(t *testing.T) {

	resp := GetAndParseMLBTeams()
	if resp == nil {
		t.Fatal("Expected a response, got nil")
	}
	for _, team := range resp {
		for _, player := range team.Roster {
			err := player.GetGameLog("2024", "R")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(player.Hitting, player.Pitching)

		}
	}

}
