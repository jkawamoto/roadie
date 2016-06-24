package command

import (
	"bytes"
	"fmt"
	"html/template"
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
	body         struct {
		APT    []string `yaml:"apt,omitempty"`
		Source string   `yaml:"source,omitempty"`
		Data   []string `yaml:"data,omitempty"`
		Run    []string `yaml:"run,omitempty"`
		Result string   `yaml:"result,omitempty"`
		Upload []string `yaml:"upload,omitempty"`
	}
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
	conf, err := template.New(filepath.Base(filename)).Funcs(funcs).ParseFiles(filename)
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

	// Construct a script object.
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	} else if strings.Contains(hostname, ".") {
		hostname = strings.Split(hostname, ".")[0]
	}

	res := script{
		filename: filename,
		instanceName: fmt.Sprintf(
			"%s-%s-%s", hostname, util.Basename(filename), time.Now().Format("20060102150405")),
	}

	// Unmarshal YAML file.
	if err := yaml.Unmarshal(buf.Bytes(), &res.body); err != nil {
		return nil, err
	}

	return &res, nil

}

// Set a git repository to source section.
func (s *script) setGitSource(repo string) {
	if s.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("%s has source section but a Git repository is given. The source section will be overwritten to '%s'."),
			s.filename, repo)
	}
	s.body.Source = repo
}

// Set a URL to source section.
func (s *script) setURLSource(url string) {
	if s.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("%s has source section but a repository URL is given. The source section will be overwritten to '%s'."),
			s.filename, url)
	}
	s.body.Source = url
}

// Upload source files and set that location to source section.
func (s *script) setLocalSource(path, project, bucket string) error {
	if s.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("%s has source section but a path for source codes is given. The source section will be overwritten."),
			s.filename)
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	var name string
	var arcPath string
	if info.IsDir() {

		filename := s.instanceName + ".tar.gz"
		arcPath = filepath.Join(os.TempDir(), filename)
		fmt.Printf("Creating an archived file %s", arcPath)
		if err := util.Archive(path, arcPath, nil); err != nil {
			return err
		}
		name = filename

	} else {

		arcPath = path
		name = util.Basename(path)

	}

	location, err := UploadToGCS(project, bucket, SourcePrefix, name, arcPath)
	if err != nil {
		return err
	}
	s.body.Source = location
	return nil

}

// Set result section with a given bucket name.
func (s *script) setResult(bucket string) {

	location := util.CreateURL(bucket, ResultPrefix, s.instanceName)
	s.body.Result = location.String()

}

// Convert to string.
func (s *script) String() string {
	res, _ := yaml.Marshal(s.body)
	return string(res)
}
