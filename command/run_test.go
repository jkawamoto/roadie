//
// command/run_test.go
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
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/gce"
	"github.com/jkawamoto/roadie/config"
	"github.com/jkawamoto/roadie/script"
)

type MockStorageServicer struct{}

// Upload a given stream with a given file name; returned string represents
// a URI assosiated with the uploaded file.
func (s *MockStorageServicer) Upload(ctx context.Context, filename string, in io.Reader) (string, error) {
	return script.RoadieSchemePrefix + filename, nil
}

// Download a file associated with a given file name and write it to a given
// writer.
func (s *MockStorageServicer) Download(ctx context.Context, filename string, out io.Writer) error {
	return nil
}

// GetFileInfo gets file information of a given filename.
func (s *MockStorageServicer) GetFileInfo(ctx context.Context, filename string) (*cloud.FileInfo, error) {
	return nil, nil
}

// List up files matching a given prefix.
// It takes a handler; information of found files are sent to it.
func (s *MockStorageServicer) List(ctx context.Context, prefix string, handler cloud.FileInfoHandler) error {
	return nil
}

// Delete a given file.
func (s *MockStorageServicer) Delete(ctx context.Context, filename string) error {
	return nil
}

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

	conf := &config.Config{
		GcpConfig: gce.GcpConfig{
			Bucket: "somebucket",
		},
	}
	ctx := config.NewContext(context.Background(), conf)
	storage := cloud.NewStorage(&MockStorageServicer{}, ioutil.Discard)

	var s script.Script
	var err error

	// Test with directories.
	for _, target := range []string{".", "../command"} {

		s = script.Script{
			InstanceName: "test",
		}

		t.Logf("Trying target %s", target)
		if err = setLocalSource(ctx, storage, &s, target, nil, true); err != nil {
			t.Error(err.Error())
		}
		if !strings.HasSuffix(s.Source, "test.tar.gz") {
			t.Errorf("source section is not correct: %s", s.Source)
		}

	}

	// Test with a file.
	s = script.Script{
		InstanceName: "test",
	}
	if err = setLocalSource(ctx, storage, &s, "run.go", nil, true); err != nil {
		t.Error(err.Error())
	}
	if !strings.HasSuffix(s.Source, "run.go") {
		t.Errorf("source section is not correct: %s", s.Source)
	}

	// Test with unexisting file.
	if err = setLocalSource(ctx, storage, &s, "abcd.efg", nil, true); err == nil {
		t.Error("Give an unexisting path but no error occurs.")
	}
	t.Logf("Give an unexisting path to setLocalSource and got an error: %s", err.Error())

}

// TestSetSource checks setSource sets correct url from a given filename.
// Since all source files are archived by tar.gz, if the given filename doesn't
// have extension .tar.gz, it should be added.
func TestSetSource(t *testing.T) {

	var s script.Script
	conf := &config.Config{
		GcpConfig: gce.GcpConfig{
			Bucket: "somebucket",
		},
	}

	s = script.Script{}
	setSource(conf, &s, "abc.tar.gz")
	if s.Source != script.RoadieSchemePrefix+"abc.tar.gz" {
		t.Errorf("source section is not correct: %s", s.Source)
	}

	s = script.Script{}
	setSource(conf, &s, "abc")
	if s.Source != script.RoadieSchemePrefix+"abc.tar.gz" {
		t.Errorf("source section is not correct: %s", s.Source)
	}

}
