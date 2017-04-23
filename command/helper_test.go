//
// command/helper_test.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This file is part of Roadie.
//
// Roadie is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Roadie is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"context"
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/cloud/mock"
	"github.com/jkawamoto/roadie/script"
)

// TestSetGitSource checks setGitSource sets correct repository URL.
func TestSetGitSource(t *testing.T) {

	var s script.Script

	s = script.Script{}
	setGitSource(&s, "https://github.com/jkawamoto/roadie.git")
	if s.Source != "https://github.com/jkawamoto/roadie.git" {
		t.Errorf("source section is not correct: %s", s.Source)
	}

	s = script.Script{}
	setGitSource(&s, "git@github.com:jkawamoto/roadie.git")
	if s.Source != "https://github.com/jkawamoto/roadie.git" {
		t.Errorf("source section is not correct: %s", s.Source)
	}

	s = script.Script{}
	setGitSource(&s, "github.com/jkawamoto/roadie")
	if s.Source != "https://github.com/jkawamoto/roadie.git" {
		t.Errorf("source section is not correct: %s", s.Source)
	}

}

// TestSetLocalSource checks setLocalSource sets correct url with any directories,
// and file paths. This test doesn't check excludes parameters since those parameters
// are tested in tests for util.Archive.
func TestSetLocalSource(t *testing.T) {

	provider := mock.NewProvider()
	m := &Metadata{
		Context:  context.Background(),
		provider: provider,
		Logger:   log.New(ioutil.Discard, "", log.LstdFlags),
		Spinner:  spinner.New(spinner.CharSets[14], 100*time.Millisecond),
	}
	m.Spinner.Writer = ioutil.Discard

	var s script.Script
	var err error

	// Test with directories.
	for _, target := range []string{".", "../command"} {

		s = script.Script{
			InstanceName: "test",
		}

		t.Logf("Trying target %s", target)
		if err = setLocalSource(m, &s, target, nil); err != nil {
			t.Error(err.Error())
		}
		if !strings.HasSuffix(s.Source, "test.tar.gz") {
			t.Error("source section is not correct:", s.Source)
		}

	}

	// Test with a file.
	s = script.Script{
		InstanceName: "test",
	}
	if err = setLocalSource(m, &s, "run.go", nil); err != nil {
		t.Error(err.Error())
	}
	if !strings.HasSuffix(s.Source, "run.go") {
		t.Error("source section is not correct:", s.Source)
	}

	// Test with unexisting file.
	if err = setLocalSource(m, &s, "abcd.efg", nil); err == nil {
		t.Error("Give an unexisting path but no error occurs.")
	}
	t.Logf("Give an unexisting path to setLocalSource and got an error: %s", err.Error())

}

// TestSetSource checks setSource sets correct url from a given filename.
// Since all source files are archived by tar.gz, if the given filename doesn't
// have extension .tar.gz, it should be added.
func TestSetSource(t *testing.T) {

	var s script.Script

	s = script.Script{}
	setSource(&s, "abc.tar.gz")
	if s.Source != script.RoadieSchemePrefix+"source/abc.tar.gz" {
		t.Errorf("source section is not correct: %s", s.Source)
	}

	s = script.Script{}
	setSource(&s, "abc")
	if s.Source != script.RoadieSchemePrefix+"source/abc.tar.gz" {
		t.Errorf("source section is not correct: %s", s.Source)
	}

}

func TestUploadFiles(t *testing.T) {

	provider := mock.NewProvider()
	m := Metadata{
		Context:  context.Background(),
		provider: provider,
		Logger:   log.New(ioutil.Discard, "", log.LstdFlags),
		Spinner:  spinner.New(spinner.CharSets[14], 100*time.Millisecond),
	}
	m.Spinner.Writer = ioutil.Discard

	var err error
	// Test uploading a directory without renaming.
	if _, err = uploadFiles(&m, ".", "", nil); err != nil {
		t.Error(err.Error())
	} else if _, exist := provider.MockStorageManager.Storage["command.tar.gz"]; !exist {
		t.Error("Failed to upload a directory")
	}

	// Test uploading a directory with a file name.
	if _, err = uploadFiles(&m, ".", "dir", nil); err != nil {
		t.Error(err.Error())
	} else if _, exist := provider.MockStorageManager.Storage["dir.tar.gz"]; !exist {
		t.Error("Faild to upload a directory with a file name")
	}

	// Test uploading a file without renaming.
	if _, err = uploadFiles(&m, "source.go", "", nil); err != nil {
		t.Error(err.Error())
	} else if _, exist := provider.MockStorageManager.Storage["source.go"]; !exist {
		t.Error("Faild tp upload a file")
	}

	// Test uploading a file with renaming.
	if _, err = uploadFiles(&m, "helper_test.go", "another_test.go", nil); err != nil {
		t.Error(err.Error())
	} else if _, exist := provider.MockStorageManager.Storage["another_test.go"]; !exist {
		t.Error("Faild to upload a file with renaming")
	}

}
