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
	"github.com/jkawamoto/roadie/config"
	"github.com/jkawamoto/roadie/resource"
	"github.com/urfave/cli"
)

// RoadieSchemePrefix is the prefix of roadie scheme URLs.
const RoadieSchemePrefix = "roadie://"

// runOpt manages all arguments and flags defined in run command.
type runOpt struct {
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

	conf := config.FromCliContext(c)
	opt := runOpt{
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
	if err := cmdRun(conf, &opt); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	if c.Bool("follow") {
		return cmdLog(&logOpt{
			Context:      util.GetContext(c),
			InstanceName: opt.InstanceName,
			Timestamp:    true,
			Follow:       true,
			Output:       os.Stdout,
		})
	}
	return nil

}

// cmdRun implements the main logic of run command.
func cmdRun(conf *config.Config, opt *runOpt) (err error) {

	if conf.Project == "" {
		return fmt.Errorf("project ID must be given")
	}
	if conf.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Project)
		conf.Bucket = conf.Project
	}

	script, err := resource.NewScript(opt.ScriptFile, opt.ScriptArgs)
	if err != nil {
		return
	}
	if err = replaceURLScheme(conf, script); err != nil {
		return
	}

	// Update instance name.
	// If an instance name is not given, use the default name.
	if opt.InstanceName != "" {
		script.InstanceName = strings.ToLower(opt.InstanceName)
	}

	// Prepare a context.
	ctx, cancel := context.WithCancel(config.NewContext(context.Background(), conf))
	defer cancel()

	// Check a specified bucket exists and create it if not.
	storage := cloud.NewStorage(ctx)
	if err = storage.PrepareBucket(); err != nil {
		return
	}

	// Check source section.
	switch {
	case opt.Git != "":
		if script.Source != "" {
			fmt.Printf(
				chalk.Red.Color("The source section of %s will be overwritten to '%s' since a Git repository is given.\n"),
				script.Filename, opt.Git)
		}
		if err = setGitSource(script, opt.Git); err != nil {
			return
		}

	case opt.URL != "":
		if script.Source != "" {
			fmt.Printf(
				chalk.Red.Color("The source section of %s will be overwritten to '%s' since a repository URL is given.\n"),
				script.Filename, opt.URL)
		}
		script.Source = opt.URL

	case opt.Local != "":
		if err = setLocalSource(ctx, storage, script, opt.Local, opt.Exclude, opt.Dry); err != nil {
			return
		}

	case opt.Source != "":
		setSource(conf, script, opt.Source)

	case script.Source == "":
		fmt.Println(chalk.Red.Color("No source section and source flags are given."))
	}

	// Check result section.
	if script.Result == "" || opt.OverWriteResultSection {
		location := util.CreateURL(conf.Bucket, ResultPrefix, script.InstanceName)
		script.Result = location.String()
	} else {
		fmt.Printf(
			chalk.Red.Color("Since result section is given in %s, all outputs will be stored in %s.\n"), script.Filename, script.Result)
		fmt.Println(
			chalk.Red.Color("Those buckets might not be retrieved from this program and manually downloading results is required."))
		fmt.Println(
			chalk.Red.Color("To manage outputs by this program, delete result section or set --overwrite-result-section flag."))
	}

	// Debugging info.
	fmt.Printf("Script to be run:\n%s\n", script.String())

	if len(opt.Queue) == 0 {
		// If queue flag is not given, execute the script in one instance.

		// Prepare startup script.
		options := " "
		if opt.NoShutdown {
			options = "--no-shutdown"
		}
		if opt.Retry <= 0 {
			opt.Retry = 10
		}

		var startup string
		startup, err = resource.Startup(&resource.StartupOpt{
			Name:    script.InstanceName,
			Script:  script.String(),
			Options: options,
			Image:   opt.Image,
			Retry:   opt.Retry,
		})

		if opt.Dry {
			// If dry flag is set, just print the startup script.
			fmt.Printf("Startup script:\n%s\n", startup)
		} else {
			err = createInstance(ctx, script.InstanceName, startup, opt.DiskSize, os.Stderr)
		}

	} else {
		// If queue name is given, the script will be enqueued in the queue.
		// If there are no instances working with the queue,
		// one instance should be created.
		task := resource.Task{
			InstanceName: script.InstanceName,
			Image:        opt.Image,
			Body:         script.ScriptBody,
			QueueName:    opt.Queue,
		}

		store := cloud.NewDatastore(ctx)
		store.Insert(time.Now().Unix(), &task)

		var worker bool
		var instances map[string]struct{}
		instances, err = runningInstances(ctx)
		if err != nil {
			return
		}
		for name := range instances {
			if strings.HasPrefix(name, task.QueueName) {
				worker = true
			}
		}
		if !worker {
			// If there are no instance working to the queue, creating it.
			// Name of the new instance must start with queue name and
			// current UNIX time should follow it.
			var startup string
			instanceName := fmt.Sprintf("%s-%d", opt.Queue, time.Now().Unix())
			startup, err = resource.WorkerStartup(&resource.WorkerStartupOpt{
				ProjectID:    conf.Project,
				Name:         opt.Queue,
				InstanceName: instanceName,
				Version:      QueueManagerVersion,
			})
			if err != nil {
				return err
			}
			err = createInstance(ctx, instanceName, startup, opt.DiskSize, os.Stderr)
		}
	}

	return
}

// setGitSource sets a Git repository `repo` to source section in a given `script`.
// If overwriting source section, it prints warning, too.
func setGitSource(script *resource.Script, repo string) (err error) {

	if strings.HasPrefix(repo, "git@") {
		sp := strings.SplitN(repo[len("git@"):], ":", 2)
		if len(sp) != 2 {
			return fmt.Errorf("Given git repository URL is invalid: %s", repo)
		}
		script.Source = fmt.Sprintf("https://%s/%s", sp[0], sp[1])
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
		script.Source = u.String()

	}
	return

}

// setLocalSource sets a GCS URL to source section in a given `script` under a given context.
// It uploads source codes specified by `path` to GCS and set the URL pointing
// the uploaded files to the source section. If filename patters are given
// by `excludes`, files matching such patters are excluded to upload.
// To upload files to GCS, `conf` is used.
// If dry is true, it does not upload any files but create a temporary file.
func setLocalSource(ctx context.Context, storage *cloud.Storage, script *resource.Script, path string, excludes []string, dry bool) (err error) {

	conf, err := config.FromContext(ctx)
	if err != nil {
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		return
	}

	var filename string      // File name on GCS.
	var uploadingPath string // File path to be uploaded.

	if info.IsDir() { // Directory will be archived.

		filename = script.InstanceName + ".tar.gz"
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

	var location string // URL where the archive is uploaded.
	if dry {
		location = util.CreateURL(conf.Bucket, SourcePrefix, filename).String()
	} else {
		location, err = storage.UploadFile(SourcePrefix, filename, uploadingPath)
		if err != nil {
			return
		}
	}
	script.Source = location
	return nil

}

// setSource sets a URL to a `file` in source directory in GCS to a given `script`.
// Source codes will be downloaded from the URL. To determine bucketname, it requires
// config. If overwriting source section, it prints warning, too.
func setSource(conf *config.Config, script *resource.Script, file string) {

	if !strings.HasSuffix(file, ".tar.gz") {
		file += ".tar.gz"
	}

	url := util.CreateURL(conf.Bucket, SourcePrefix, file).String()
	if script.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s will be overwritten to '%s' since a filename is given.\n"),
			script.Filename, url)
	}
	script.Source = url
}

// replaceURLScheme replaced URLs which start with "roadie://".
// Those URLs are modified to "gs://<bucketname>/.roadie/".
func replaceURLScheme(conf *config.Config, script *resource.Script) error {

	offset := len(RoadieSchemePrefix)

	// Replace source section.
	if strings.HasPrefix(script.Source, RoadieSchemePrefix) {
		script.Source = util.CreateURL(conf.Bucket, SourcePrefix, script.Source[offset:]).String()
	}

	// Replace data section.
	for i, url := range script.Data {
		if strings.HasPrefix(url, RoadieSchemePrefix) {
			script.Data[i] = util.CreateURL(conf.Bucket, DataPrefix, url[offset:]).String()
		}
	}

	// Replace result section.
	if strings.HasPrefix(script.Result, RoadieSchemePrefix) {
		script.Result = util.CreateURL(conf.Bucket, ResultPrefix, script.Result[offset:]).String()
	}

	return nil
}

// createInstance creates an instance under a given context.
// The new instance has a given name and a given startup script.
// It also has a data disk of which volume size is as same as disk.
// Output messages will be outputted to a given writer, output.
func createInstance(ctx context.Context, name, startup string, disk int64, output io.Writer) (err error) {

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Writer = output
	s.Prefix = fmt.Sprintf("Creating an instance named %s...", chalk.Bold.TextStyle(name))
	s.FinalMSG = fmt.Sprintf("\n%s\rInstance created.\n", strings.Repeat(" ", len(s.Prefix)+2))

	s.Start()
	defer s.Stop()

	err = cloud.CreateInstance(ctx, name, []*cloud.MetadataItem{
		&cloud.MetadataItem{
			Key:   "startup-script",
			Value: startup,
		},
	}, disk)
	if err != nil {
		s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create instance.\n"), strings.Repeat(" ", len(s.Prefix)+2))
	}
	return

}
