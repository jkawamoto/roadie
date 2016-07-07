package command

import (
	"testing"

	"github.com/jkawamoto/roadie/config"
)

// TestCmdDataPut checks if wrong patterns are given, cmdDataPut returns error,
// and if empty pattern is given, it do nothing.
func TestCmdDataPut(t *testing.T) {

	conf := config.Config{}

	// Test for wrong pattern.
	if err := cmdDataPut(&conf, "[b-a", ""); err == nil {
		t.Error("Give a wrong pattern but no errors occur.")
	} else {
		t.Logf("Wrong patter makes an error: %s", err.Error())
	}

	// Test for empty pattern.
	if err := cmdDataPut(&conf, "", ""); err != nil {
		t.Errorf(err.Error())
	}

}
