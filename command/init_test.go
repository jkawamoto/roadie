//
// command/init_test.go
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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/config"
	yaml "gopkg.in/yaml.v2"
)

// readConfigFile reads and parses a given config file.
func readConfigFile(filename string) (cfg *config.Config, err error) {

	var data []byte
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	cfg = new(config.Config)
	err = yaml.Unmarshal(data, cfg)
	return

}

func TestCheckNotEmpty(t *testing.T) {

	cases := []struct {
		input string
		err   bool
	}{
		{"aaa", false},
		{"", true},
	}
	for _, c := range cases {

		t.Run(fmt.Sprintf("Input=%q", c.input), func(t *testing.T) {
			err := checkNotEmpty(c.input)
			if (!c.err && err != nil) || (c.err && err == nil) {
				t.Errorf("checkNotEmpty(%q) returns %v, expected %v", c.input, err, c.err)
			}
		})

	}

}

func TestCheckOption(t *testing.T) {

	options := []string{"1", "2", "3"}
	cases := []struct {
		input string
		err   bool
	}{
		{"1", false},
		{"", true},
		{"4", true},
	}
	f := checkOption(options...)
	for _, c := range cases {
		t.Run(fmt.Sprintf("Input=%q", c.input), func(t *testing.T) {
			err := f(c.input)
			if (c.err && err == nil) || (!c.err && err != nil) {
				t.Errorf("checkOption(%q) returns %v, expected %v", c.input, err, c.err)
			}
		})
	}

}

func TestCmdInit(t *testing.T) {

	var err error
	var output bytes.Buffer
	m := testMetadata(&output, nil)
	testID := "test-id"

	t.Run("init for GCP", func(t *testing.T) {

		var tmp string
		tmp, err = ioutil.TempDir("", "")
		if err != nil {
			t.Fatalf("cannot create a temporary directory: %v", err)
		}
		defer os.RemoveAll(tmp)

		var wd string
		wd, err = os.Getwd()
		if err != nil {
			t.Fatalf("Getwd returns an error: %v", err)
		}
		err = os.Chdir(tmp)
		if err != nil {
			t.Fatalf("cannot change directory: %v", err)
		}
		defer os.Chdir(wd)

		m.Stdin = strings.NewReader("g\n" + testID + "\n")
		err = cmdInit(m)
		if err != nil {
			t.Fatalf("cmdInit returns an error: %v", err)
		}
		if m.Config.GcpConfig.Project != testID {
			t.Errorf("project ID is set %q, want %v", m.Config.GcpConfig.Project, testID)
		}

		var cfg *config.Config
		cfg, err = readConfigFile(m.Config.FileName)
		if err != nil {
			t.Fatalf("readConfigFile returns an error: %v", err)
		}

		if cfg.GcpConfig.Project != testID {
			t.Errorf("stored config has a wrong project ID %q, want %v", cfg.GcpConfig.Project, testID)
		}

	})

	t.Run("init for Azure", func(t *testing.T) {

		var tmp string
		tmp, err = ioutil.TempDir("", "")
		if err != nil {
			t.Fatalf("cannot create a temporary directory: %v", err)
		}
		defer os.RemoveAll(tmp)

		var wd string
		wd, err = os.Getwd()
		if err != nil {
			t.Fatalf("Getwd returns an error: %v", err)
		}
		err = os.Chdir(tmp)
		if err != nil {
			t.Fatalf("cannot change directory: %v", err)
		}
		defer os.Chdir(wd)

		testSubscriptionID := "subscription"
		testProjectID := "myproject"
		m.Stdin = strings.NewReader(
			strings.Join([]string{"a", testID, testSubscriptionID, testProjectID}, "\n") + "\n")
		err = cmdInit(m)
		if err != nil {
			t.Fatalf("cmdInit returns an error: %v", err)
		}
		if m.Config.AzureConfig.TenantID != testID {
			t.Errorf("tennant ID is %q, want %v", m.Config.AzureConfig.TenantID, testID)
		}
		if m.Config.AzureConfig.SubscriptionID != testSubscriptionID {
			t.Errorf("subscription ID is %q, want %v", m.Config.AzureConfig.SubscriptionID, testSubscriptionID)
		}
		if m.Config.AzureConfig.ProjectID != testProjectID {
			t.Errorf("project ID is %q, want %v", m.Config.AzureConfig.ProjectID, testProjectID)
		}

		var cfg *config.Config
		cfg, err = readConfigFile(m.Config.FileName)
		if err != nil {
			t.Fatalf("readConfigFile returns an error: %v", err)
		}

		if cfg.AzureConfig.TenantID != testID {
			t.Errorf("stored config has a wrong tennant ID %q, want %v", cfg.AzureConfig.TenantID, testID)
		}
		if cfg.AzureConfig.SubscriptionID != testSubscriptionID {
			t.Errorf("stored config has a wrong tennant ID %q, want %v", cfg.AzureConfig.SubscriptionID, testSubscriptionID)
		}

	})

	t.Run("input nothing", func(t *testing.T) {

		m.Stdin = strings.NewReader("\n\n")
		err = cmdInit(m)
		if err == nil {
			t.Error("input nothing but no errors are returned")
		}

	})

	t.Run("config file is given", func(t *testing.T) {

		var tmp string
		tmp, err = ioutil.TempDir("", "")
		if err != nil {
			t.Fatalf("cannot create a temporary directory: %v", err)
		}
		defer os.RemoveAll(tmp)
		filename := filepath.Join(tmp, ConfigFile)

		m.Config.FileName = filename
		m.Stdin = strings.NewReader("g\n" + testID + "\n")
		err = cmdInit(m)
		if err != nil {
			t.Fatalf("cmdInit returns an error: %v", err)
		}
		if m.Config.GcpConfig.Project != testID {
			t.Errorf("project ID is set %q, want %v", m.Config.GcpConfig.Project, testID)
		}

		var cfg *config.Config
		cfg, err = readConfigFile(filename)
		if err != nil {
			t.Fatalf("readConfigFile returns an error: %v", err)
		}

		if cfg.GcpConfig.Project != testID {
			t.Errorf("stored config has a wrong project ID %q, want %v", cfg.GcpConfig.Project, testID)
		}

	})

}
