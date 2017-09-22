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
	"net/url"
	"testing"

	"github.com/jkawamoto/roadie/script"
)

// TestSetGitSource checks setGitSource sets correct repository URL.
// setGitSource takes several types of URLs for a git repository such as
// - https://github.com/user/repos.git (via https)
// - git@github.com:user/repos.git (via ssl)
// - github.com/user/repos (ambiguous)
// It parses all of the above formats and sets a https URL to the given script.
func TestSetGitSource(t *testing.T) {

	var err error
	var s script.Script
	cases := []string{
		"https://github.com/jkawamoto/roadie.git",
		"git@github.com:jkawamoto/roadie.git",
		"github.com/jkawamoto/roadie",
	}
	expect := "https://github.com/jkawamoto/roadie.git"

	for _, v := range cases {
		s = script.Script{}
		err = setGitSource(&s, v)
		if err != nil {
			t.Fatalf("error with setGitSource(%v, %v): %v", s, v, err)
		}
		if s.Source != expect {
			t.Errorf("Source = %v, want %v", s.Source, expect)
		}
	}

}

// TestSetLocalSource checks setLocalSource sets correct URLs.
// setLocalSource takes a directory path or a file path.
// If a directory path is given, it makes a tarball named the script name
// with `.tar.gz` suffix and uploads the tarball to the cloud storage.
// If a file path is given, setLocalSource uploads it to the cloud storage.
// setLocalSource, then, sets a URL for the uploaded file to the source section.
// The URL must have `roadie://source` prefix so that cloud providers can modify
// it to their specific forms.
//
// Note that this test doesn't check excludes parameters since those parameters
// are tested in tests for util.Archive.
func TestSetLocalSource(t *testing.T) {

	var err error
	m := testMetadata(nil)
	storage, err := m.StorageManager()
	if err != nil {
		t.Fatalf("cannot get a storage manager: %v", err)
	}

	var s script.Script
	var loc *url.URL
	expected := script.RoadieSchemePrefix + script.SourcePrefix + "/test.tar.gz"
	for _, target := range []string{".", "../command", "run.go"} {

		s = script.Script{
			Name: "test",
		}

		err = setLocalSource(m, &s, target, nil)
		if err != nil {
			t.Fatalf("error with setLocalSource of %v: %v", target, err)
		}
		if s.Source != expected {
			t.Errorf("Source = %v, want %v", s.Source, expected)
		}

		loc, err = url.Parse(s.Source)
		if err != nil {
			t.Fatalf("cannot parse %q: %v", s.Source, err)
		}
		_, err = storage.GetFileInfo(m.Context, loc)
		if err != nil {
			t.Errorf("%v doesn't exist", s.Source)
		}

	}

	// Test with an unexisting file.
	err = setLocalSource(m, &s, "abcd.efg", nil)
	if err == nil {
		t.Error("setLocalSource with an unexisting path doesn't return any errors")
	}

}

// TestSetUploadedSource checks setSource sets correct url from a given filename.
// Since all source files are archived by tar.gz, if the given filename doesn't
// have extension .tar.gz, it should be added.
func TestSetUploadedSource(t *testing.T) {

	var s script.Script
	cases := []string{"abc.tar.gz", "abc"}
	expected := script.RoadieSchemePrefix + script.SourcePrefix + "/abc.tar.gz"

	for _, c := range cases {
		s = script.Script{}
		setUploadedSource(&s, c)
		if s.Source != expected {
			t.Errorf("Source = %v, want %v", s.Source, expected)
		}
	}

}
