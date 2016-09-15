//
// command/table.go
//
// Copyright (c) 2016 Junpei Kawamoto
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/command/util"
)

// AddRecorder is a callback to add file information to a table.
type AddRecorder func(table *uitable.Table, info *util.FileInfo, quiet bool)

// PrintFileList prints a list of files having a given prefix.
func PrintFileList(project, bucket, prefix string, url, quiet bool) (err error) {

	var headers []string
	if url {
		headers = []string{"FILE NAME", "SIZE", "TIME CREATED", "URL"}
	} else {
		headers = []string{"FILE NAME", "SIZE", "TIME CREATED"}
	}

	return printList(
		project, bucket, prefix, quiet, headers,
		func(table *uitable.Table, info *util.FileInfo, quiet bool) {

			if info.Name != "" {
				if quiet {
					table.AddRow(info.Name)
				} else if url {
					table.AddRow(info.Name, fmt.Sprintf(
						"%dKB", info.Size/1024), info.TimeCreated.Format(PrintTimeFormat),
						fmt.Sprintf("gs://%s/%s", bucket, info.Path))
				} else {
					table.AddRow(info.Name, fmt.Sprintf(
						"%dKB", info.Size/1024), info.TimeCreated.Format(PrintTimeFormat))
				}
			}

		})
}

// PrintDirList prints a list of directoris in a given prefix.
func PrintDirList(project, bucket, prefix string, url, quiet bool) (err error) {

	var headers []string
	if url {
		headers = []string{"INSTANCE NAME", "TIME CREATED", "URL"}
	} else {
		headers = []string{"INSTANCE NAME", "TIME CREATED"}
	}

	// Storing previous folder name.
	prev := ""

	return printList(
		project, bucket, prefix, quiet, headers,
		func(table *uitable.Table, info *util.FileInfo, quiet bool) {

			rel, _ := filepath.Rel(prefix, info.Path)
			rel = filepath.Dir(rel)

			if rel != "." && rel != prev {
				if quiet {
					table.AddRow(rel)
				} else if url {
					table.AddRow(
						rel, info.TimeCreated.Format(PrintTimeFormat),
						fmt.Sprintf("gs://%s/%s", bucket, rel))
				} else {
					table.AddRow(rel, info.TimeCreated.Format(PrintTimeFormat))
				}
				prev = rel
			}

		})
}

func printList(project, bucket, prefix string, quiet bool, headers []string, addRecorder AddRecorder) (err error) {

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.FinalMSG = fmt.Sprintf("\n%s\r", strings.Repeat(" ", len(s.Prefix)+2))
	s.Start()

	table := uitable.New()
	if !quiet {
		rawHeaders := make([]interface{}, len(headers))
		for i, v := range headers {
			rawHeaders[i] = v
		}
		table.AddRow(rawHeaders...)
	}

	err = util.ListupFiles(context.Background(), project, bucket, prefix, func(storage *util.Storage, info *util.FileInfo) error {

		addRecorder(table, info, quiet)
		return nil

	})

	s.Stop()
	if err == nil {
		fmt.Println(table.String())
	}
	return

}
