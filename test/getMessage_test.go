package test

import (
    "testing"
    "github.com/federicolivarez/challengeMeli/utils"
	"github.com/federicolivarez/challengeMeli/config"
	// "fmt"
)

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