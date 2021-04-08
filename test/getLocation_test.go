package test

import (
    "testing"
    "github.com/federicolivarez/challengeMeli/utils"
	"github.com/federicolivarez/challengeMeli/config"
	// "fmt"
)

func TestGetLocationX(t *testing.T) {
	config.GetConfiguration("../")
	Positions := utils.GetLocation( 100.0,115.5, 142.7)
	expectedx := -312.1124559649421

	if Positions.Ejex == expectedx {
		t.Logf("Expected: %v, got: %v", expectedx, Positions.Ejex)
	} else{
		t.Errorf("Expected: %v, got: %v", expectedx, Positions.Ejex)
	}
}

func TestGetLocationY(t *testing.T) {
	config.GetConfiguration("../")
	Positions := utils.GetLocation( 100.0,115.5, 142.7)
	expectedy := 154.5355105513955

	if Positions.Ejey == expectedy {
		t.Logf("Expected: %v, got: %v", expectedy, Positions.Ejey)
	} else{
		t.Errorf("Expected: %v, got: %v", expectedy, Positions.Ejey)
	}
}