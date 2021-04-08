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

func TestGetBadMessage(t *testing.T) {
	config.GetConfiguration("../")

	var msg1 = []string{"este","","","",""}
	var msg2 = []string{"","es","","mensaje","",""}
	var msg3 = []string{"este","","un","","hola",}

	MessageCompleted := utils.GetMessage(msg1,msg2,msg3)
	expectedMsj := "este es un mensaje"

	if MessageCompleted == expectedMsj {
		t.Errorf("Expected: %v, got: %v", expectedMsj, MessageCompleted)
	} else{
		t.Logf("Expected: %v, got: %v", expectedMsj, MessageCompleted)
	}
}

func TestGetOkMessage(t *testing.T) {
	config.GetConfiguration("../")

	var msg1 = []string{"este","","","",""}
	var msg2 = []string{"","es","","mensaje",""}
	var msg3 = []string{"este","","un","","secreto"}

	MessageCompleted := utils.GetMessage(msg1,msg2,msg3)
	expectedMsj := "este es un mensaje secreto"

	if MessageCompleted == expectedMsj {
		t.Logf("Expected: %v, got: %v", expectedMsj, MessageCompleted)
	} else{
		t.Errorf("Expected: %v, got: %v", expectedMsj, MessageCompleted)
	}
}