package mlb

import (
	"fmt"
	"testing"
)

func TestGetMLBTeams(t *testing.T) {
	resp := GetAndParseMLBTeams()
	fmt.Println(resp)

	if resp == nil {
		t.Fatal("Expected a response, got nil")
	}

}
