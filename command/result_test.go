//
// command/result_test.go
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
	"bufio"
	"bytes"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
)

// TestCmdResultList tests cmdResultList.
func TestCmdResultList(t *testing.T) {

	var err error
	var output bytes.Buffer
	files := []string{
		"roadie://result/instance1/stdout1.txt",
		"roadie://result/instance1/stdout2.txt",
		"roadie://result/instance1/stdout3.txt",
		"roadie://result/instance11/stdout11.txt",
		"roadie://result/instance11/stdout12.txt",
	}

	m := testMetadata(&output, nil)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	cases := []struct {
		instance string
		expected []string
	}{
		{"instance1", []string{"stdout1.txt", "stdout2.txt", "stdout3.txt"}},
		{"", []string{"instance1", "instance11"}},
	}

	for _, c := range cases {

		err = cmdResultList(m, c.instance, false, true)
		if err != nil {
			t.Fatalf("cmdResultList returns an error: %v", err)
		}

		res := make(map[string]struct{})
		scanner := bufio.NewScanner(&output)
		for scanner.Scan() {
			res[strings.TrimSpace(strings.Split(scanner.Text(), "\t")[0])] = struct{}{}
		}
		if len(res) != len(c.expected) {
			t.Errorf("%v items are shown, want %v items", len(res), len(c.expected))
		}

		for _, item := range c.expected {
			if _, exist := res[item]; !exist {
				t.Errorf("%v isn't shown", item)
			}
		}
		output.Reset()

	}

}

func TestCmdResultShow(t *testing.T) {

	var err error
	var output bytes.Buffer
	files := []string{
		"roadie://result/instance1/stdout1.txt",
		"roadie://result/instance1/stdout2.txt",
		"roadie://result/instance1/fig1.png",
		"roadie://result/instance11/stdout11.txt",
		"roadie://result/instance11/stdout12.txt",
	}

	m := testMetadata(&output, nil)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	cases := []struct {
		instance string
		filename string
		expected []string
	}{
		{"instance1", "", []string{
			"*** stdout1.txt ***" + files[0] + "*** stdout2.txt ***" + files[1],
			"*** stdout2.txt ***" + files[1] + "*** stdout1.txt ***" + files[0],
		}},
		{"instance1", "stdout", []string{
			files[0] + files[1],
			files[1] + files[0],
		}},
		{"instance1", "fig1.png", []string{files[2]}},
	}

	for _, c := range cases {
		err = cmdResultShow(m, c.instance, c.filename)
		if err != nil {
			t.Fatalf("cmdResultShow returns an error: %v", err)
		}
		res := strings.Replace(output.String(), "\n", "", -1)

		var check bool
		for _, e := range c.expected {
			if res == e {
				check = true
			}
		}
		if !check {
			t.Errorf("outptted %v, want %v", res, c.expected)
		}
		output.Reset()
	}

}

func TestCmdResultGet(t *testing.T) {

	var err error
	files := []string{
		"roadie://result/instance1/stdout1.txt",
		"roadie://result/instance1/stdout2.txt",
		"roadie://result/instance1/fig1.png",
		"roadie://result/instance11/stdout11.txt",
		"roadie://result/instance11/stdout12.txt",
	}

	m := testMetadata(nil, nil)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	cases := []struct {
		instance string
		queries  []string
		expected []string
	}{
		{"instance1", nil, files[:3]},
		{"instance1", []string{"stdout*"}, files[:2]},
		{"instance1", []string{"stdout*", "fig*"}, files[:3]},
	}

	for _, c := range cases {

		tmp, err := ioutil.TempDir("", "TestCmdResultGet")
		if err != nil {
			t.Fatalf("TempDir returns an error: %v", err)
		}
		defer os.RemoveAll(tmp)

		err = cmdResultGet(m, c.instance, c.queries, tmp)
		if err != nil {
			t.Fatalf("cmdResultGet returns an error: %v", err)
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

func TestCmdResultDelete(t *testing.T) {

	var err error
	files := []string{
		"roadie://result/instance1/stdout1.txt",
		"roadie://result/instance1/stdout2.txt",
		"roadie://result/instance1/fig1.png",
		"roadie://result/instance11/stdout11.txt",
		"roadie://result/instance11/stdout12.txt",
	}

	m := testMetadata(nil, nil)
	err = uploadDummyFiles(m, files)
	if err != nil {
		t.Fatalf("uploadDummyFiles returns an error: %v", err)
	}

	s, err := m.StorageManager()
	if err != nil {
		t.Fatalf("cannot get a storage manager: %v", err)
	}

	loc, err := url.Parse("roadie://result/")
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	cases := []struct {
		instance  string
		queries   []string
		reminings []string
	}{
		{"instance1", []string{"stdout*"}, files[2:]},
		{"instance11", nil, files[2:3]},
	}

	for _, c := range cases {

		err = cmdResultDelete(m, c.instance, c.queries)
		if err != nil {
			t.Fatalf("cmdResultDelete returns an error: %v", err)
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
