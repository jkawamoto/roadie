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
	"bytes"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

// uploadDummyFiles uploads given files to a storage specified in given
// metadata.
func uploadDummyFiles(m *Metadata, files []string) (err error) {

	s, err := m.StorageManager()
	if err != nil {
		return
	}

	var loc *url.URL
	for _, f := range files {
		loc, err = url.Parse(f)
		if err != nil {
			return
		}
		err = s.Upload(m.Context, loc, strings.NewReader(f))
		if err != nil {
			return
		}
	}
	return

}

// TestUploadDummyFiles tests uploadDummyFiles.
func TestUploadDummyFiles(t *testing.T) {

	var err error
	var output bytes.Buffer
	m := testMetadata(&output)
	s, err := m.StorageManager()
	if err != nil {
		t.Fatalf("cannot get a storage manager: %v", err)
	}

	files := []string{
		"roadie://test/instance1/stdout1.txt",
		"roadie://another/instance1/stdout3.txt",
	}

	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	var loc *url.URL
	for _, f := range files {
		loc, err = url.Parse(f)
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}

		var buf bytes.Buffer
		err = s.Download(m.Context, loc, &buf)
		if err != nil {
			t.Fatalf("Download returns an error: %v", err)
		}

		if buf.String() != f {
			t.Errorf("uploaded file is %v, want %v", buf, f)
		}

	}

}

func TestCreateURL(t *testing.T) {

	cases := []struct {
		container string
		path      string
		expect    string
	}{
		{"container", "", script.RoadieSchemePrefix + "container/"},
		{"container", "file", script.RoadieSchemePrefix + "container/file"},
		{"container", "dir/", script.RoadieSchemePrefix + "container/dir/"},
	}

	for _, c := range cases {
		loc, err := createURL(c.container, c.path)
		if err != nil {
			t.Fatalf("createURL(%q, %q) returns an error: %v", c.container, c.path, err)
		}
		if loc.String() != c.expect {
			t.Errorf("createURL(%q, %q) = %q, want %v", c.container, c.path, loc, c.expect)
		}
	}

}

// TestCmdGet tests cmdGet.
func TestCmdGet(t *testing.T) {

	var err error
	var output bytes.Buffer
	files := []string{
		"roadie://container1/stdout1.txt",
		"roadie://container1/stdout2.txt",
		"roadie://container1/fig3.png",
		"roadie://container2/stdout1.txt",
		"roadie://container2/stdout2.txt",
	}

	m := testMetadata(&output)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	cases := []struct {
		container string
		queries   []string
		expected  []string
	}{
		{"container1", []string{"stdout*"}, files[:2]},
		{"container1", []string{"stdout*", "fig*"}, files[:3]},
	}

	for _, c := range cases {

		tmp, err := ioutil.TempDir("", "")
		if err != nil {
			t.Fatalf("TempDir returns an error: %v", err)
		}
		defer os.RemoveAll(tmp)

		err = cmdGet(m, c.container, c.queries, tmp)
		if err != nil {
			t.Fatalf("cmdGet returns an error: %v", err)
		}

		var data []byte
		for _, f := range c.expected {
			data, err = ioutil.ReadFile(path.Join(tmp, path.Base(f)))
			if err != nil {
				t.Fatalf("cannot read a expected file %v: %v", f, err)
			}
			if string(data) != f {
				t.Errorf("downloaded file is %v, want %v", string(data), f)
			}

		}

	}

}

// TestCmdDelete tests cmdDelete.
func TestCmdDelete(t *testing.T) {

	var err error
	var output bytes.Buffer
	files := []string{
		"roadie://container1/stdout1.txt",
		"roadie://container1/stdout2.txt",
		"roadie://container1/fig3.png",
		"roadie://container2/stdout1.txt",
		"roadie://container2/stdout2.txt",
	}

	m := testMetadata(&output)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	s, err := m.StorageManager()
	if err != nil {
		t.Fatalf("cannot get a storage manager: %v", err)
	}

	loc, err := url.Parse("roadie://")
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	cases := []struct {
		container string
		queries   []string
		reminings []string
	}{
		{"container1", []string{"stdout*"}, files[2:]},
		{"container2", []string{"*"}, files[2:3]},
	}

	for _, c := range cases {

		err = cmdDelete(m, c.container, c.queries)
		if err != nil {
			t.Fatalf("cmdDelete returns an error: %v", err)
		}

		reminings := make(map[string]struct{})
		err = s.List(m.Context, loc, func(info *cloud.FileInfo) error {
			if !strings.HasSuffix(info.URL.Path, "/") {
				reminings[info.URL.String()] = struct{}{}
			}
			return nil
		})
		if err != nil {
			t.Fatalf("List returns an error: %v", err)
		}

		if len(reminings) != len(c.reminings) {
			t.Errorf("remining data files %v, want %v", len(reminings), len(c.reminings))
		}
		for _, f := range c.reminings {
			if _, exist := reminings[f]; !exist {
				t.Errorf("%v is not found", f)
			}
		}

	}

}

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
