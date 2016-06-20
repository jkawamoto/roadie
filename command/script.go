package command

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jkawamoto/roadie-cli/util"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

type script struct {
	filename     string
	instanceName string
	body         map[interface{}]interface{}
}

// Load a given script file and apply arguments.
func loadScript(filename string, args []string) (*script, error) {

	// Define function map to replace place holders.
	funcs := template.FuncMap{}
	for _, v := range args {
		sp := strings.Split(v, "=")
		if len(sp) >= 2 {
			funcs[sp[0]] = func() string {
				return sp[1]
			}
		}
	}

	// Load YAML config file.
	conf, err := template.New(filename).Funcs(funcs).ParseFiles(filename)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			return nil, err
		default:
			return nil, fmt.Errorf("Cannot apply variables to the place holders in %s", filename)
		}
	}

	// Replace place holders with given args.
	buf := &bytes.Buffer{}
	if err := conf.Execute(buf, nil); err != nil {
		return nil, err
	}

	// Unmarshal YAML file.
	body := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(buf.Bytes(), body); err != nil {
		return nil, err
	}

	return &script{
		filename:     filename,
		instanceName: util.Basename(filename) + time.Now().Format("20060102150405"),
		body:         body,
	}, nil

}

// Set a git repository to source section.
func (s *script) setGitSource(repo string) {
	if _, ok := s.body[source]; ok {
		log.Printf(
			chalk.Red.Color("%s has source section but a Git repository is given. The source section will be overwritten to '%s'."),
			s.filename, repo)
	}
	s.body[source] = repo
}

// Set a URL to source section.
func (s *script) setURLSource(url string) {
	if _, ok := s.body[source]; ok {
		log.Printf(
			chalk.Red.Color("%s has source section but a repository URL is given. The source section will be overwritten to '%s'."),
			s.filename, url)
	}
	s.body[source] = url
}

// Upload source files and set that location to source section.
func (s *script) setLocalSource(path, project, bucket string) error {
	if _, ok := s.body[source]; ok {
		log.Printf(
			chalk.Red.Color("%s has source section but a path for source codes is given. The source section will be overwritten."),
			s.filename)
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	var arcPath string
	var location *url.URL
	if info.IsDir() {

		filename := s.instanceName + ".tar.gz"
		arcPath = filepath.Join(os.TempDir(), filename)
		log.Printf("Create an archived file %s", arcPath)
		if err := util.Archive(path, arcPath, nil); err != nil {
			return err
		}
		location = util.CreateURL(bucket, "source", filename)

	} else {

		arcPath = path
		location = util.CreateURL(bucket, "source", util.Basename(path))

	}

	log.Printf("Uploading to %s", location)
	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return err
	}
	if err := storage.Upload(arcPath, location); err != nil {
		return err
	}

	s.body[source] = location.String()
	return nil

}

// Set result section with a given bucket name.
func (s *script) setResult(bucket string) {

	location := util.CreateURL(bucket, "result", s.instanceName)
	s.body["result"] = location

}

// Convert to string.
func (s *script) String() string {
	res, _ := yaml.Marshal(s.body)
	return string(res)
}
