package command

import "testing"

func TestLoadScript(t *testing.T) {

	_, err := loadScript("../test.yml", []string{"method=test"})
	if err != nil {
		t.Error(err.Error())
	}

}

func TestSetGitSource(t *testing.T) {

	s, err := loadScript("../test.yml", []string{"method=test"})
	if err != nil {
		t.Error(err.Error())
	}

	s.body.Source = ""
	s.setGitSource("https://github.com/jkawamoto/roadie-gcp.git")
	if s.body.Source != "https://github.com/jkawamoto/roadie-gcp.git" {
		t.Errorf("setGitSource doesn't work: %s", s.body.Source)
	}

}
