//
// command/table_test.go
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
	"fmt"
	"path"
	"strings"
	"testing"
)

func TestPrintFileList(t *testing.T) {

	var err error
	var output bytes.Buffer
	files := []string{
		"roadie://test/instance1/stdout1.txt",
		"roadie://test/instance1/stdout2.txt",
		"roadie://test/instance1/stdout3.txt",
		"roadie://test/instance2/stdout1.txt",
		"roadie://test/instance2/stdout2.txt",
		"roadie://another/instance1/stdout1.txt",
		"roadie://another/instance1/stdout2.txt",
		"roadie://another/instance1/stdout3.txt",
	}

	m := testMetadata(&output)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	expected := files[:3]
	cases := []struct {
		url   bool
		quiet bool
	}{
		{false, false},
		{true, false},
		{false, true},
		{true, true},
	}

	for _, c := range cases {

		err = PrintFileList(m, "test", "instance1", c.url, c.quiet)
		if err != nil {
			t.Fatalf("PrintFileList returns an error: %v", err)
		}

		if c.quiet {

			files := make(map[string]struct{})
			for _, v := range strings.Split(output.String(), "\n") {
				v = strings.TrimSpace(v)
				if v != "" {
					files[v] = struct{}{}
				}
			}

			if len(files) != len(expected) {
				t.Errorf("the number of printed files is %v. want %v", len(files), len(expected))
			}
			for _, e := range expected {
				if _, ok := files[path.Base(e)]; !ok {
					t.Errorf("file %v isn't shown", e)
				}
			}

		} else {

			res := strings.Split(strings.TrimRight(output.String(), "\n"), "\n")
			if len(res)-1 != len(expected) {
				t.Errorf("the number of files is %v, want %v", len(res)-1, len(expected))
			}
			if !strings.HasPrefix(res[0], "FILE NAME") {
				t.Errorf("shown table doesn't have correct header: %v", res[0])
			}

			size := make(map[string]string)
			for _, line := range res[1:] {
				items := strings.Split(line, "\t")
				size[strings.TrimSpace(items[0])] = strings.TrimSpace(items[1])
			}
			for _, e := range expected {
				if size[path.Base(e)] != fmt.Sprintf("%dB", len(e)) {
					t.Errorf("file size of %v is %v, want %vB", e, size[path.Base(e)], len(e))
				}
			}

			if c.url {
				urls := make(map[string]string)
				for _, line := range res[1:] {
					items := strings.Split(line, "\t")
					urls[strings.TrimSpace(items[0])] = strings.TrimSpace(items[len(items)-1])
				}
				for _, e := range expected {
					if urls[path.Base(e)] != e {
						t.Errorf("URL of %v is %v, want %v", e, urls[path.Base(e)], e)
					}
				}
			}

		}
		output.Reset()

	}

}

func TestPrintDirList(t *testing.T) {

	var err error
	var output bytes.Buffer
	files := []string{
		"roadie://test/instance1/stdout1.txt",
		"roadie://test/instance1/stdout2.txt",
		"roadie://test/instance1/stdout3.txt",
		"roadie://test/instance2/stdout1.txt",
		"roadie://test/instance2/stdout2.txt",
		"roadie://another/instance1/stdout1.txt",
		"roadie://another/instance1/stdout2.txt",
		"roadie://another/instance1/stdout3.txt",
	}

	m := testMetadata(&output)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	expected := []string{
		"roadie://test/instance1/",
		"roadie://test/instance2/",
	}
	cases := []struct {
		url   bool
		quiet bool
	}{
		{false, false},
		{true, false},
		{false, true},
		{true, true},
	}

	for _, c := range cases {

		err = PrintDirList(m, "test", "", c.url, c.quiet)
		if err != nil {
			t.Fatalf("PrintFileList returns an error: %v", err)
		}

		if c.quiet {

			files := make(map[string]struct{})
			for _, v := range strings.Split(output.String(), "\n") {
				v = strings.TrimSpace(v)
				if v != "" {
					files[v] = struct{}{}
				}
			}

			if len(files) != len(expected) {
				t.Errorf("the number of printed files is %v. want %v", len(files), len(expected))
			}
			for _, e := range expected {
				if _, ok := files[path.Base(e)]; !ok {
					t.Errorf("directory %v isn't shown", e)
				}
			}

		} else {

			res := strings.Split(strings.TrimRight(output.String(), "\n"), "\n")
			if len(res)-1 != len(expected) {
				t.Errorf("the number of files is %v, want %v", len(res)-1, len(expected))
			}
			if !strings.HasPrefix(res[0], "INSTANCE NAME") {
				t.Errorf("shown table doesn't have correct header: %v", res[0])
			}

			if c.url {
				urls := make(map[string]string)
				for _, line := range res[1:] {
					items := strings.Split(line, "\t")
					urls[strings.TrimSpace(items[0])] = strings.TrimSpace(items[len(items)-1])
				}
				for _, e := range expected {
					if urls[path.Base(e)] != e {
						t.Errorf("URL of %v is %q, want %v", e, urls[path.Base(e)], e)
					}
				}
			}

		}
		output.Reset()

	}

}
