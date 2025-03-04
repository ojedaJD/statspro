package mlb

import (
	"fmt"
	"testing"
)

func TestGetMLBRoster(t *testing.T) {

	resp := GetAndParseMLBTeams()

	if resp == nil {
		t.Fatal("Expected a response, got nil")
	}
	fmt.Println(resp[0].Roster[0].Person)

}
