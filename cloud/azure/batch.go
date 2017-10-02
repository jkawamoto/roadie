//
// cloud/azure/batch.go
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

// This source file is associated with Azure's Batch API of which Swagger's
// clients are stored in `batch` directory.

package azure

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/azure/batch/client"
	"github.com/jkawamoto/roadie/cloud/azure/batch/client/accounts"
	"github.com/jkawamoto/roadie/cloud/azure/batch/client/compute_nodes"
	"github.com/jkawamoto/roadie/cloud/azure/batch/client/jobs"
	"github.com/jkawamoto/roadie/cloud/azure/batch/client/pools"
	"github.com/jkawamoto/roadie/cloud/azure/batch/client/tasks"
	"github.com/jkawamoto/roadie/cloud/azure/batch/models"
	"github.com/jkawamoto/roadie/script"
)

const (
	// BatchAPIVersion defines API version of batch service.
	BatchAPIVersion = "2016-07-01.3.1"
	// RoadieAzureArchiveName is an archive name of roarie-azure command.
	RoadieAzureArchiveName = "roadie-azure_linux_amd64.tar.gz"

	// JobManagerURL = ""
)

// MinimalJSONProducer is a provider which marshals message bodies as a JSON
// object and remove null values from it.
type MinimalJSONProducer struct {
	regexp *regexp.Regexp
	blank  []byte
}

// NewMinimalJSONProducer creates a new MinimalJSONProducer.
func NewMinimalJSONProducer() *MinimalJSONProducer {

	return &MinimalJSONProducer{
		regexp: regexp.MustCompile("(\"[^\"]+?\":null,?|,\"[^\"]+\":null)"),
		blank:  []byte(""),
	}

}

// Produce creates a message body from a given request message.
func (p *MinimalJSONProducer) Produce(out io.Writer, msg interface{}) (err error) {

	data, err := json.Marshal(msg)
	if err != nil {
		return
	}
	data = p.regexp.ReplaceAllLiteral(data, p.blank)

	_, err = out.Write(data)
	return

}

// AuthorizedTransporter is a transporter which adds authentication information
// to each request before transporting it.
type AuthorizedTransporter struct {
	http.RoundTripper
	account string
	key     []byte
}

// NewAuthorizedTransporter creates a new authorized transporter with a given
// account name and shared key.
func NewAuthorizedTransporter(transport http.RoundTripper, account string, key []byte) *AuthorizedTransporter {

	return &AuthorizedTransporter{
		RoundTripper: transport,
		account:      account,
		key:          key,
	}

}

// RoundTrip computes a shared key from a given request and add Authorization
// header to the request.
func (t *AuthorizedTransporter) RoundTrip(req *http.Request) (*http.Response, error) {

	var ocpHeaderKeys []string
	for key := range req.Header {
		if strings.HasPrefix(key, "Ocp-") || strings.HasPrefix(key, "ocp-") {
			ocpHeaderKeys = append(ocpHeaderKeys, key)
		}
	}
	sort.Strings(ocpHeaderKeys)

	var ocpHeaders []string
	for _, key := range ocpHeaderKeys {
		ocpHeaders = append(ocpHeaders, fmt.Sprintf("%s:%s", strings.ToLower(key), req.Header.Get(key)))
	}

	canonicalizedHeaders := strings.Join(ocpHeaders, "\n")
	canonicalizedResource := "/" + t.account + req.URL.Path

	var queryKeys []string
	for key := range req.URL.Query() {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)
	for _, key := range queryKeys {
		v, _ := url.QueryUnescape(req.URL.Query().Get(key))
		canonicalizedResource += "\n" + strings.ToLower(key) + ":" + v
	}

	stringToSign := strings.Join([]string{
		strings.ToUpper(req.Method),
		req.Header.Get("Content-Encoding"),
		req.Header.Get("Content-Language"),
		getContentLength(req),
		req.Header.Get("Content-MD5"),
		req.Header.Get("Content-Type"),
		req.Header.Get("Date"),
		req.Header.Get("If-Modified-Since"),
		req.Header.Get("If-Match"),
		req.Header.Get("If-None-Match"),
		req.Header.Get("If-Unmodified-Since"),
		req.Header.Get("Range"),
		canonicalizedHeaders,
		canonicalizedResource,
	}, "\n")

	if apiAccessDebugMode {
		fmt.Println("StringToSign:", stringToSign)
	}

	hash := hmac.New(sha256.New, t.key)
	hash.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	req.Header.Add("Authorization", fmt.Sprintf("SharedKey %v:%v", t.account, signature))

	return t.RoundTripper.RoundTrip(req)

}

// getContentLength returns a string representing the context length of a given
// request for calculating a shared key associated with the request.
// If the request doesn't have Content-Length header, it returns an empty
// string.
func getContentLength(r *http.Request) string {
	if r.ContentLength <= 0 {
		return ""
	}
	return fmt.Sprintf("%v", r.ContentLength)
}

// BatchService provides an interface for Azure's batch service.
type BatchService struct {
	client    *client.BatchService
	storage   *StorageService
	gmt       *time.Location
	Config    *AzureConfig
	Logger    *log.Logger
	SleepTime time.Duration
}

// JobSet represents a set of jobs.
type JobSet map[string]*models.CloudJob

// TaskSet represents a set of tasks.
type TaskSet map[string]*models.CloudTask

// NewBatchService creates a new batch service interface assosiated with
// a given config; to authorize a authentication token
// is required.
func NewBatchService(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (service *BatchService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	// Create a resource group if not exist.
	err = CreateResourceGroupIfNotExist(ctx, cfg, logger)
	if err != nil {
		return
	}

	// Create a management client.
	mcli, err := NewBatchManagementService(ctx, cfg, logger)
	var accounts BatchAccountSet
	if err != nil {
		return
	}
	if accounts, err = mcli.BatchAccounts(ctx); err != nil {
		return
	} else if _, exists := accounts[cfg.BatchAccount]; !exists {
		err = mcli.CreateBatchAccount(ctx)
		if err != nil {
			return
		}
	}
	key, err := mcli.GetKey(ctx)
	if err != nil {
		return
	}

	// Create a service client.
	scli := client.NewHTTPClient(strfmt.NewFormats())

	switch transport := scli.Transport.(type) {
	case *httptransport.Runtime:
		// Update the host to {account-name}.{region-id}.batch.azure.com
		transport.Host = fmt.Sprintf("%v.%v.batch.azure.com", cfg.BatchAccount, cfg.Location)
		transport.Debug = apiAccessDebugMode
		transport.Transport = NewAuthorizedTransporter(transport.Transport, cfg.BatchAccount, key)
		transport.Producers["application/json; odata=minimalmetadata"] = NewMinimalJSONProducer()
		scli.Accounts.SetTransport(transport)
		scli.Jobs.SetTransport(transport)

	}

	// Create a storage service client.
	storage, err := NewStorageService(ctx, cfg, logger)
	if err != nil {
		return
	}

	gmt, _ := time.LoadLocation("GMT")
	service = &BatchService{
		client:    scli,
		storage:   storage,
		gmt:       gmt,
		Config:    cfg,
		Logger:    logger,
		SleepTime: DefaultSleepTime,
	}
	return

}

// Nodes retrieves information of compute nodes in a given named pool.
func (s *BatchService) Nodes(ctx context.Context, pool string) (nodes []*models.ComputeNode, err error) {

	s.Logger.Println("Retrieving compute nodes in pool", pool)
	res, err := s.client.ComputeNodes.ComputeNodeList(
		compute_nodes.NewComputeNodeListParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithPoolID(pool).
			WithOcpDate(s.getOcpDate()))
	if err != nil {
		return
	}
	nodes = res.Payload.Value

	s.Logger.Println("Finished retrieving compute nodes in pool", pool)
	return

}

// GetPoolInfo retrieves information of a given named pool.
func (s *BatchService) GetPoolInfo(ctx context.Context, name string) (info *models.CloudPool, err error) {

	s.Logger.Println("Retrieving information of pool", name)
	res, err := s.client.Pools.PoolGet(
		pools.NewPoolGetParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithPoolID(name).
			WithOcpDate(s.getOcpDate()))
	if err != nil {
		return
	}

	info = res.Payload
	s.Logger.Println("Finished retrieving information of pool", name)
	return

}

// UpdatePoolSize requests updating the size of the given named pool to size.
// Note that: resizing pool size is an asynchronous operation.
func (s *BatchService) UpdatePoolSize(ctx context.Context, name string, size int32) (err error) {

	s.Logger.Println("Updating the size of pool", name)
	_, err = s.client.Pools.PoolResize(
		pools.NewPoolResizeParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithPoolID(name).
			WithOcpDate(s.getOcpDate()).
			WithPoolResizeParameter(&models.PoolResizeParameter{
				TargetDedicated: &size,
			}))
	if err != nil {
		return
	}
	s.Logger.Println("The update request is accepted")
	return

}

// CreateJob creates a job which has a given name.
func (s *BatchService) CreateJob(ctx context.Context, name string) (err error) {

	// TODO:
	// 1. Check metadata, if error returns, it means no app exists, then upload.
	// 2. If version metadata is old or snapshot, upload new version.
	// 3. otherwise create url and use it.
	var execURL string
	expired := true

	bin, err := url.Parse(script.RoadieSchemePrefix + path.Join(BinContainer, RoadieAzureArchiveName))
	if err != nil {
		return
	}
	_, err = s.storage.GetFileInfo(ctx, bin)
	if err == nil && !expired {
		execURL = s.storage.getFileURL(BinContainer, RoadieAzureArchiveName)

	} else {
		s.Logger.Println("Job management program is not found")

		var fp *os.File
		fp, err = os.Open(filepath.Join(os.Getenv("GOPATH"), "src/github.com/jkawamoto/roadie-azure/pkg/snapshot/roadie-azure_linux_amd64.tar.gz"))
		if err != nil {
			return
		}
		defer fp.Close()

		err = s.storage.UploadWithMetadata(ctx, BinContainer, RoadieAzureArchiveName, fp, map[string]string{
			"Content-Type": "application/tar+gzip",
			"Version":      "snapshot",
		})
		// TODO: The above logic should be replaced to download the archive from
		// GitHub directory.
		// var res *http.Response
		// res, err = ctxhttp.Get(ctx, nil, JobManagerURL)
		// if err != nil {
		// 	return
		// }
		// defer res.Body.Close()
		// s.blobClient.GetBlobURL(group, filename)
		// execURL, err = s.storage.Upload(ctx, BinContainer, "roadie-azure", res.Body)
		if err != nil {
			return
		}
		execURL = s.storage.getFileURL(BinContainer, RoadieAzureArchiveName)
		s.Logger.Println("Job management program is uploaded at", execURL)

	}

	// Upload the config file.
	configFilename := fmt.Sprintf("%v%v-init.cfg", name, time.Now().Unix())
	configString, err := s.Config.String()
	if err != nil {
		return
	}
	err = s.storage.UploadWithMetadata(ctx, StartupContainer, configFilename, strings.NewReader(configString), map[string]string{
		"Content-Type": "text/yaml",
	})
	if err != nil {
		return
	}
	configURL := s.storage.getFileURL(StartupContainer, configFilename)

	s.Logger.Println("Creating job", name)
	_, err = s.client.Jobs.JobAdd(
		jobs.NewJobAddParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithOcpDate(s.getOcpDate()).
			WithJob(&models.JobAddParameter{
				ID: &name,
				PoolInfo: &models.PoolInformation{
					AutoPoolSpecification: &models.AutoPoolSpecification{
						AutoPoolIDPrefix: "roadie",
						KeepAlive:        false,
						Pool: &models.PoolSpecification{
							VMSize: &s.Config.MachineType,
							VirtualMachineConfiguration: &models.VirtualMachineConfiguration{
								ImageReference: &models.ImageReference{
									Publisher: &s.Config.OS.PublisherName,
									Offer:     &s.Config.OS.Offer,
									Sku:       &s.Config.OS.Skus,
									Version:   s.Config.OS.Version,
								},
								NodeAgentSKUID: toPtr("batch.node.ubuntu 16.04"),
							},
							TargetDedicated: 1,
						},
						PoolLifetimeOption: toPtr(models.AutoPoolSpecificationPoolLifetimeOptionJob),
					},
				},
				JobPreparationTask: &models.JobPreparationTask{
					CommandLine: toPtr(fmt.Sprintf(
						`sh -c "tar -zxvf %v -C ${AZ_BATCH_NODE_SHARED_DIR} --strip-components=1 && sudo ${AZ_BATCH_NODE_SHARED_DIR}/roadie-azure init %v %v"`,
						RoadieAzureArchiveName, configFilename, name)),
					ResourceFiles: []*models.ResourceFile{
						&models.ResourceFile{
							BlobSource: &execURL,
							FilePath:   toPtr(RoadieAzureArchiveName),
						},
						&models.ResourceFile{
							BlobSource: &configURL,
							FilePath:   &configFilename,
						},
					},
					RunElevated: true,
				},
				OnAllTasksComplete: models.CloudJobOnAllTasksCompleteTerminateJob,
				OnTaskFailure:      models.CloudJobOnTaskFailurePerformExitOptionsJobAction,
			}))
	if err != nil {
		err = NewAPIError(err)

	} else {
		var set JobSet
		for {
			if set, err = s.Jobs(ctx); err != nil {
				break
			} else if _, exist := set[name]; exist {
				s.Logger.Println("Created job", name)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	}

	s.Logger.Println("Cannot create a job:", err.Error())
	return

}

// EnableJob starts a given named job which was stopped.
func (s *BatchService) EnableJob(ctx context.Context, name string) (err error) {

	s.Logger.Println("Enabling job", name)
	_, err = s.client.Jobs.JobEnable(
		jobs.NewJobEnableParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithJobID(name).
			WithOcpDate(s.getOcpDate()))

	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot enable job", name, ":", err.Error())
	} else {
		s.Logger.Println("Enabled job", name)
	}
	return

}

// DisableJob stops a given named job.
func (s *BatchService) DisableJob(ctx context.Context, name string) (err error) {

	s.Logger.Println("Disabling job", name)
	_, err = s.client.Jobs.JobDisable(
		jobs.NewJobDisableParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithJobID(name).
			WithOcpDate(s.getOcpDate()).
			WithJobDisableParameter(&models.JobDisableParameter{
				DisableTasks: toPtr(models.JobDisableParameterDisableTasksRequeue),
			}))

	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot disable job", name, ":", err.Error())
	} else {
		s.Logger.Println("Disabled job", name)
	}
	return

}

// Jobs retrieves a set of jobs defined in the batch account specified in
// the configuration.
func (s *BatchService) Jobs(ctx context.Context) (set JobSet, err error) {

	s.Logger.Println("Retriving jobs")
	res, err := s.client.Jobs.JobList(
		jobs.NewJobListParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithOcpDate(s.getOcpDate()))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve jobs:", err.Error())
		return
	}

	set = make(JobSet)
	for _, v := range res.Payload.Value {
		set[v.ID] = v
	}
	s.Logger.Println("Retrieved jobs")
	return

}

// DeleteJob deletes a given named job.
func (s *BatchService) DeleteJob(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting job", name)
	_, err = s.client.Jobs.JobDelete(
		jobs.NewJobDeleteParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithOcpDate(s.getOcpDate()).
			WithJobID(name))

	if err != nil {
		err = NewAPIError(err)

	} else {
		var set JobSet
		for {
			if set, err = s.Jobs(ctx); err != nil {
				break
			} else if _, exist := set[name]; !exist {
				s.Logger.Println("Deleted job", name)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	}

	s.Logger.Println("Cannot delete job", name, ":", err.Error())
	return

}

// GetJobInfo retrives the information of the given named job.
func (s *BatchService) GetJobInfo(ctx context.Context, job string) (info *models.CloudJob, err error) {

	s.Logger.Println("Retrieving information of job", job)
	res, err := s.client.Jobs.JobGet(
		jobs.NewJobGetParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithJobID(job).
			WithOcpDate(s.getOcpDate()))
	if err != nil {
		return
	}

	info = res.Payload
	s.Logger.Println("Finished retrieving information of job", job)
	return

}

// CreateTask adds a given task to a given named job.
func (s *BatchService) CreateTask(ctx context.Context, job string, task *script.Script) (err error) {

	var resourceFiles []*models.ResourceFile

	// Update source section.
	// If the URL schema of the source file is `roadie`, remove the URL and put it
	// as a resource file.
	if strings.HasPrefix(task.Source, script.RoadieSchemePrefix) {

		u, err2 := url.Parse(task.Source)
		if err2 != nil {
			return err2
		}
		filename := filepath.Base(u.Path)

		task.Source = fmt.Sprintf("file://./%v", filename)
		resourceFiles = append(resourceFiles, &models.ResourceFile{
			BlobSource: toPtr(s.storage.getFileURL(u.Hostname(), u.Path)),
			FilePath:   &filename,
		})

	}

	// Update data section.
	var newData []string
	for _, v := range task.Data {
		if strings.HasPrefix(v, script.RoadieSchemePrefix) {
			src, dest := parseRenamableURL(v)
			u, err2 := url.Parse(src)
			if err2 != nil {
				return err2
			}
			resourceFiles = append(resourceFiles, &models.ResourceFile{
				BlobSource: toPtr(s.storage.getFileURL(u.Hostname(), u.Path)),
				FilePath:   &dest,
			})

		} else {
			newData = append(newData, v)
		}

	}
	task.Data = newData

	now := time.Now().Unix()
	// Create a startup script and upload it.
	startupFilename := fmt.Sprintf("%v%v.yml", task.Name, now)
	err = s.storage.UploadWithMetadata(ctx, StartupContainer, startupFilename, strings.NewReader(task.String()), map[string]string{
		"Content-Type": "text/yaml",
	})
	if err != nil {
		return
	}
	resourceFiles = append(resourceFiles, &models.ResourceFile{
		BlobSource: toPtr(s.storage.getFileURL(StartupContainer, startupFilename)),
		FilePath:   &startupFilename,
	})

	// Upload the config file.
	configFilename := fmt.Sprintf("%v%v.cfg", task.Name, now)
	configString, err := s.Config.String()
	if err != nil {
		return
	}
	err = s.storage.UploadWithMetadata(ctx, StartupContainer, configFilename, strings.NewReader(configString), map[string]string{
		"Content-Type": "text/yaml",
	})
	if err != nil {
		return
	}
	resourceFiles = append(resourceFiles, &models.ResourceFile{
		BlobSource: toPtr(s.storage.getFileURL(StartupContainer, configFilename)),
		FilePath:   &configFilename,
	})

	// Create an instance.
	s.Logger.Println("Creating a task in job", job)
	_, err = s.client.Tasks.TaskAdd(
		tasks.NewTaskAddParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithJobID(job).
			WithOcpDate(s.getOcpDate()).
			WithTask(&models.TaskAddParameter{
				ID:            &task.Name,
				CommandLine:   toPtr(fmt.Sprintf(`sh -c "sudo ${AZ_BATCH_NODE_SHARED_DIR}/roadie-azure exec %v %v %v"`, configFilename, startupFilename, task.Name)),
				ResourceFiles: resourceFiles,
				RunElevated:   true,
			}))

	if err != nil {
		err = NewAPIError(err)

	} else {
		var set TaskSet
		for {
			if set, err = s.Tasks(ctx, job); err != nil {
				break
			} else if _, exist := set[task.Name]; exist {
				s.Logger.Println("Created task", task.Name)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	}

	s.Logger.Println("Cannot create a task:", err.Error())
	return

}

// Tasks retrieves tasks in a given named job.
func (s *BatchService) Tasks(ctx context.Context, job string) (set TaskSet, err error) {

	s.Logger.Println("Retrieving tasks in job", job)
	res, err := s.client.Tasks.TaskList(
		tasks.NewTaskListParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithJobID(job).
			WithOcpDate(s.getOcpDate()))

	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve tasks:", err.Error())
	}

	set = make(TaskSet)
	for _, v := range res.Payload.Value {
		set[v.ID] = v
	}

	s.Logger.Println("Retrieved tasks in job", job)
	return

}

// DeleteTask deletes a given named task from a given named job.
func (s *BatchService) DeleteTask(ctx context.Context, job, task string) (err error) {
	// TODO: Delete related files, such as script, config, from the storage.

	s.Logger.Println("Deleting task", task)
	_, err = s.client.Tasks.TaskDelete(
		tasks.NewTaskDeleteParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithJobID(job).
			WithOcpDate(s.getOcpDate()).
			WithTaskID(task))
	if err != nil {
		err = NewAPIError(err)

	} else {
		var set TaskSet
		for {
			if set, err = s.Tasks(ctx, job); err != nil {
				break
			} else if _, exist := set[task]; !exist {
				s.Logger.Println("Deleted task", task)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	}

	s.Logger.Println("Cannot delete task", task, ":", err.Error())
	return

}

// AvailableMachineTypes returns a list of supported machine types.
// Batch supports all Azure VM sizes except STANDARD_A0 and those with premium
// storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
func (s *BatchService) AvailableMachineTypes(ctx context.Context) (types []cloud.MachineType, err error) {

	service, err := NewComputeService(ctx, s.Config, s.Logger)
	if err != nil {
		return
	}

	aux, err := service.AvailableMachineTypes(ctx)
	if err != nil {
		return
	}

	for _, v := range aux {
		switch {
		case v.Name == "Standard_A0":
			// Not supported
		case strings.HasPrefix(v.Name, "Standard_GS"):
			// Not supported
		case strings.HasPrefix(v.Name, "Standard_DS"):
			// Not supported
		default:
			types = append(types, v)
		}
	}

	return

}

// AvailableOSImages returns a list of available OS images.
func (s *BatchService) AvailableOSImages(ctx context.Context) (images []string, err error) {

	s.Logger.Println("Retrieving available os images")
	res, err := s.client.Accounts.AccountListNodeAgentSkus(
		accounts.NewAccountListNodeAgentSkusParamsWithContext(ctx).
			WithAPIVersion(BatchAPIVersion).
			WithClientRequestID(&s.Config.ClientID).
			WithOcpDate(s.getOcpDate()))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve available os images")
		return
	}

	for _, v := range res.Payload.Value {
		for _, image := range v.VerifiedImageReferences {
			images = append(
				images,
				fmt.Sprintf("%v:%v:%v:%v", *image.Publisher, *image.Offer, *image.Sku, image.Version))
		}
	}

	return

}

// getOcpDate returns a pointer for the string representing current date and
// time in RFC 1123 format. This pointer will be used with WithOcpDate method.
func (s *BatchService) getOcpDate() *string {
	now := time.Now().In(s.gmt).Format(time.RFC1123)
	return &now
}
