//
// command/run.go
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/script"
	"github.com/urfave/cli"
)

// runOpt manages all arguments and flags defined in run command.
type runOpt struct {
	*Metadata

	// Git repository URL which will be cloned as source codes.
	Git string
	// URL where source codes are stored.
	URL string
	// Directory or file path which will be used as source codes.
	Local string
	// File patterns which will be excluded from source codes in a given
	// directory. This flag will only work when `local` flag is given with
	// a directory.
	Exclude []string
	// Filename in source directory in GCS.
	Source string
	// Path for the script file to be run.
	ScriptFile string
	// Arguments for the script.
	ScriptArgs []string
	// Instance name. If not set, named by script file name and current time.
	InstanceName string
	// Base docker image name.
	Image string
	// Specify disk size of new instance.
	DiskSize int64
	// If true, result section will be overwritten so that roadie can manage
	// result data. Otherwise, users require to manage them by their self.
	OverWriteResultSection bool
	// If true, created instance will not shutdown automatically. So, users
	// require to do it by their self. This flag can be useful for debugging.
	NoShutdown bool
	// If true, do not create any instances but show startup script.
	// This flag is for debugging.
	Dry bool
	// The number of times retry roadie-gcp container when GCP's error happens.
	Retry int64
	// Queue name. If specified, the given script will be enqueued to the queue.
	Queue string
}

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m := getMetadata(c)
	opt := &runOpt{
		Metadata:               m,
		Git:                    c.String("git"),
		URL:                    c.String("url"),
		Local:                  c.String("local"),
		Exclude:                c.StringSlice("exclude"),
		Source:                 c.String("source"),
		ScriptFile:             c.Args().First(),
		ScriptArgs:             c.StringSlice("e"),
		InstanceName:           c.String("name"),
		Image:                  c.String("image"),
		DiskSize:               c.Int64("disk-size"),
		OverWriteResultSection: c.Bool("overwrite-result-section"),
		NoShutdown:             c.Bool("no-shutdown"),
		Dry:                    c.Bool("dry"),
		Retry:                  c.Int64("retry") + 1,
		Queue:                  c.String("queue"),
	}
	if err := cmdRun(opt); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	if c.Bool("follow") {
		return cmdLog(&optLog{
			Metadata:     m,
			InstanceName: opt.InstanceName,
			Timestamp:    true,
			Follow:       true,
			SleepTime:    DefaultSleepTime,
		})
	}
	return nil

}

// cmdRun implements the main logic of run command.
func cmdRun(opt *runOpt) (err error) {

	s, err := script.NewScript(opt.ScriptFile, opt.ScriptArgs)
	if err != nil {
		return
	}

	// Update instance name.
	// If an instance name is not given, use the default name.
	if opt.InstanceName != "" {
		s.InstanceName = strings.ToLower(opt.InstanceName)
	}

	// Prepare a context.
	ctx, cancel := context.WithCancel(opt.Context)
	defer cancel()

	// Check a specified bucket exists and create it if not.
	service, err := opt.StorageManager()
	if err != nil {
		return err
	}
	storage := cloud.NewStorage(service, nil)

	// Check source section.
	switch {
	case opt.Git != "":
		if s.Source != "" {
			fmt.Printf(
				chalk.Red.Color("The source section of %s will be overwritten to '%s' since a Git repository is given.\n"),
				opt.ScriptFile, opt.Git)
		}
		if err = setGitSource(s, opt.Git); err != nil {
			return
		}

	case opt.URL != "":
		if s.Source != "" {
			fmt.Printf(
				chalk.Red.Color("The source section of %s will be overwritten to '%s' since a repository URL is given.\n"),
				opt.ScriptFile, opt.URL)
		}
		s.Source = opt.URL

	case opt.Local != "":
		if err = setLocalSource(ctx, storage, s, opt.Local, opt.Exclude, opt.Dry); err != nil {
			return
		}

	case opt.Source != "":
		setSource(s, opt.Source)

	case s.Source == "":
		fmt.Println(chalk.Red.Color("No source section and source flags are given."))
	}

	// Check result section.
	if s.Result == "" || opt.OverWriteResultSection {
		s.Result = script.RoadieSchemePrefix + filepath.Join(script.ResultPrefix, s.InstanceName)
	} else {
		fmt.Printf(
			chalk.Red.Color("Since result section is given in %s, all outputs will be stored in %s.\n"), opt.ScriptFile, s.Result)
		fmt.Println(
			chalk.Red.Color("Those buckets might not be retrieved from this program and manually downloading results is required."))
		fmt.Println(
			chalk.Red.Color("To manage outputs by this program, delete result section or set --overwrite-result-section flag."))
	}

	// Debugging info.
	opt.Logger.Printf("Script to be run:\n%s\n", s.String())

	// if len(opt.Queue) == 0 {
	// If queue flag is not given, execute the script in one instance.

	// Prepare options.
	if opt.NoShutdown {
		s.Options = append(s.Options, "no-shutdown")
	}
	if opt.Retry <= 0 {
		opt.Retry = 10
	}
	s.Options = append(s.Options, fmt.Sprintf("retry:%d", opt.Retry))

	err = createInstance(opt.Metadata, s, os.Stderr)

	// } else {
	// 	// If queue name is given, the script will be enqueued in the queue.
	// 	// If there are no instances working with the queue,
	// 	// one instance should be created.
	// 	var queueManager *gce.QueueService
	// 	queueManager, err = gce.NewQueueService(ctx, &cfg.GcpConfig, ioutil.Discard)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer queueManager.Close()
	//
	// 	s.Image = opt.Image
	// 	err = queueManager.Enqueue(ctx, opt.Queue, s)
	// 	if err != nil {
	// 		return
	// 	}
	//
	// 	var workerExist bool
	// 	err = queueManager.Workers(ctx, opt.Queue, func(name string) error {
	// 		workerExist = true
	// 		return nil
	// 	})
	// 	if err != nil {
	// 		return
	// 	}
	//
	// 	if !workerExist {
	// 		err = queueManager.CreateWorkers(ctx, opt.Queue, opt.DiskSize, 1, func(name string) error {
	// 			return nil
	// 		})
	// 		if err != nil {
	// 			return
	// 		}
	// 	}
	//
	// }

	return

}

// setGitSource sets a Git repository `repo` to source section in a given `script`.
// If overwriting source section, it prints warning, too.
func setGitSource(s *script.Script, repo string) (err error) {

	if strings.HasPrefix(repo, "git@") {
		sp := strings.SplitN(repo[len("git@"):], ":", 2)
		if len(sp) != 2 {
			return fmt.Errorf("Given git repository URL is invalid: %s", repo)
		}
		s.Source = fmt.Sprintf("https://%s/%s", sp[0], sp[1])
	} else {
		u, err := url.Parse(repo)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			u.Scheme = "https"
		}
		if !strings.HasSuffix(u.Path, ".git") {
			u.Path += ".git"
		}
		s.Source = u.String()

	}
	return

}

// setLocalSource sets a GCS URL to source section in a given `script` under a given context.
// It uploads source codes specified by `path` to GCS and set the URL pointing
// the uploaded files to the source section. If filename patters are given
// by `excludes`, files matching such patters are excluded to upload.
// To upload files to GCS, `conf` is used.
// If dry is true, it does not upload any files but create a temporary file.
func setLocalSource(ctx context.Context, storage *cloud.Storage, s *script.Script, path string, excludes []string, dry bool) (err error) {

	info, err := os.Stat(path)
	if err != nil {
		return
	}

	var filename string      // File name on GCS.
	var uploadingPath string // File path to be uploaded.

	if info.IsDir() { // Directory will be archived.

		filename = s.InstanceName + ".tar.gz"
		uploadingPath = filepath.Join(os.TempDir(), filename)

		spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		spin.Prefix = fmt.Sprintf("Creating an archived file %s...", uploadingPath)
		spin.FinalMSG = fmt.Sprintf("\n%s\rCreating the archived file %s.    \n", strings.Repeat(" ", len(spin.Prefix)+2), uploadingPath)
		spin.Start()

		if err = util.Archive(path, uploadingPath, excludes); err != nil {
			spin.Stop()
			return
		}
		defer os.Remove(uploadingPath)

		spin.Stop()

	} else { // One source file just will be uploaded.

		uploadingPath = path
		filename = filepath.Base(path)

	}

	// URL where the archive is uploaded.
	location, err := storage.UploadFile(ctx, script.SourcePrefix, filename, uploadingPath)
	if err != nil {
		return
	}
	s.Source = location
	return nil

}

// setSource sets a URL to a `file` in source directory to a given `script`.
// Source codes will be downloaded from the URL. If overwriting the source
// section, it prints warning, too.
func setSource(s *script.Script, file string) {

	if !strings.HasSuffix(file, ".tar.gz") {
		file += ".tar.gz"
	}

	url := script.RoadieSchemePrefix + filepath.Join(script.SourcePrefix, file)
	if s.Source != "" {
		fmt.Printf(
			chalk.Red.Color("Source section will be overwritten to '%s' since a filename is given.\n"), url)
	}
	s.Source = url

}

// createInstance creates an instance under a given metadata.
// The new instance has a given startup script.
// Output messages will be outputted to a given writer, output.
func createInstance(m *Metadata, task *script.Script, output io.Writer) (err error) {

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Writer = output
	s.Prefix = fmt.Sprintf("Creating an instance named %s...", chalk.Bold.TextStyle(task.InstanceName))
	s.FinalMSG = fmt.Sprintf("\n%s\rInstance created.\n", strings.Repeat(" ", len(s.Prefix)+2))
	s.Start()
	defer s.Stop()

	service, err := m.InstanceManager()
	if err != nil {
		return
	}

	err = service.CreateInstance(m.Context, task)
	if err != nil {
		s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create instance.\n"), strings.Repeat(" ", len(s.Prefix)+2))
	}
	return

}
