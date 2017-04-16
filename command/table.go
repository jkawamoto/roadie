//
// command/table.go
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
	"fmt"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/cloud"
)

// AddRecorder is a callback to add file information to a table.
type AddRecorder func(table *uitable.Table, info *cloud.FileInfo, quiet bool)

// PrintFileList prints a list of files having a given prefix.
func PrintFileList(m *Metadata, prefix string, url, quiet bool) (err error) {

	var headers []string
	if url {
		headers = []string{"FILE NAME", "SIZE", "TIME CREATED", "URL"}
	} else {
		headers = []string{"FILE NAME", "SIZE", "TIME CREATED"}
	}

	return printList(m, prefix, quiet, headers, func(table *uitable.Table, info *cloud.FileInfo, quiet bool) {

		if info.Name != "" {
			if quiet {
				table.AddRow(info.Name)
			} else if url {
				table.AddRow(info.Name, fmt.Sprintf(
					"%dKB", info.Size/1024), info.TimeCreated.Format(PrintTimeFormat), info.Path)
			} else {
				table.AddRow(info.Name, fmt.Sprintf(
					"%dKB", info.Size/1024), info.TimeCreated.Format(PrintTimeFormat))
			}
		}

	})
}

// PrintDirList prints a list of directoris in a given prefix.
func PrintDirList(m *Metadata, prefix string, url, quiet bool) (err error) {

	var headers []string
	if url {
		headers = []string{"INSTANCE NAME", "TIME CREATED", "URL"}
	} else {
		headers = []string{"INSTANCE NAME", "TIME CREATED"}
	}

	// Storing previous folder name.
	prev := ""

	return printList(m, prefix, quiet, headers, func(table *uitable.Table, info *cloud.FileInfo, quiet bool) {

		rel, _ := filepath.Rel(prefix, info.Path)
		rel = filepath.Dir(rel)

		if rel != "." && rel != prev {
			if quiet {
				table.AddRow(rel)
			} else if url {
				table.AddRow(rel, info.TimeCreated.Format(PrintTimeFormat), rel)
			} else {
				table.AddRow(rel, info.TimeCreated.Format(PrintTimeFormat))
			}
			prev = rel
		}

	})
}

func printList(m *Metadata, prefix string, quiet bool, headers []string, addRecorder AddRecorder) (err error) {

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.Start()
	defer s.Stop()

	table := uitable.New()
	if !quiet {
		rawHeaders := make([]interface{}, len(headers))
		for i, v := range headers {
			rawHeaders[i] = v
		}
		table.AddRow(rawHeaders...)
	}

	service, err := m.StorageManager()
	if err != nil {
		return err
	}

	storage := cloud.NewStorage(service, nil)
	err = storage.ListupFiles(m.Context, prefix, "", func(info *cloud.FileInfo) error {
		addRecorder(table, info, quiet)
		return nil
	})
	if err == nil {
		s.FinalMSG += table.String() + "\n"
	}
	return

}
