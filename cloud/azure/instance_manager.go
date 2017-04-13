package azure

import (
	"context"
	"log"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

// InstanceManager implements cloud.InstanceManager interface to run a script
// on Azure.
type InstanceManager struct {
	service *BatchService
	Config  *AzureConfig
	Logger  *log.Logger
}

// NewInstanceManager creates a new instance manager.
func NewInstanceManager(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (m *InstanceManager, err error) {

	service, err := NewBatchService(ctx, cfg, logger)
	if err != nil {
		return
	}

	m = &InstanceManager{
		service: service,
		Config:  cfg,
		Logger:  logger,
	}
	return

}

// CreateInstance creates an instance which has a given name.
// Note that Azure doesn't support custome HDD size; argument disksize is,
// hence, ignored in this method.
func (m *InstanceManager) CreateInstance(ctx context.Context, name string, task *script.Script, disksize int64) (err error) {

	err = m.service.CreateJob(ctx, name)
	if err != nil {
		return
	}
	return m.service.CreateTask(ctx, name, task)

}

// DeleteInstance deletes the given named instance.
func (m *InstanceManager) DeleteInstance(ctx context.Context, name string) error {

	return m.service.DeleteJob(ctx, name)

}

// Instances returns a list of running instances
func (m *InstanceManager) Instances(ctx context.Context) (instances map[string]struct{}, err error) {

	jobs, err := m.service.Jobs(ctx)
	if err != nil {
		return
	}

	instances = make(map[string]struct{})
	for name := range jobs {
		instances[name] = struct{}{}
	}
	return

}

// AvailableRegions returns a list of available regions.
func (m *InstanceManager) AvailableRegions(ctx context.Context) (regions []cloud.Region, err error) {

	m.Logger.Println("Retrieving available regions")
	regions, err = Locations(ctx, &m.Config.Token, m.Config.SubscriptionID)
	if err != nil {
		m.Logger.Println("Cannot retrieve available regions")
	} else {
		m.Logger.Println("Retrieved available regions")
	}
	return

}

// AvailableMachineTypes returns a list of available machine types.
func (m *InstanceManager) AvailableMachineTypes(ctx context.Context) (types []cloud.MachineType, err error) {
	return m.service.AvailableMachineTypes(ctx)
}
