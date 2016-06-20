package command

import "testing"

func TestCmdRun(t *testing.T) {

	// CmdRun(cli.NewContext())

}

func TestBasename(t *testing.T) {

	if res := basename("somefile.go"); res != "somefile" {
		t.Errorf("%s does not match 'somefile'", res)
	}

	if res := basename("noext"); res != "noext" {
		t.Errorf("%s does not match 'noext'", res)
	}

	if res := basename("/path/to/somefile.go"); res != "somefile" {
		t.Errorf("%s does not match 'somefile'", res)
	}

}
