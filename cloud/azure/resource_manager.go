package azure

import (
	"context"
	"log"

	"github.com/jkawamoto/roadie/cloud"
)

// ResourceManager implements cloud.ResourceManager.
type ResourceManager struct {
	Config *Config
	Logger *log.Logger
}

// NewResourceManager creates a new ResourceManager.
func NewResourceManager(cfg *Config, logger *log.Logger) *ResourceManager {
	return &ResourceManager{
		Config: cfg,
		Logger: logger,
	}
}

// GetProjectID returns an ID of the current project.
func (m *ResourceManager) GetProjectID() string {
	return m.Config.ResourceGroupName
}

// SetProjectID sets an ID to the current project.
func (m *ResourceManager) SetProjectID(id string) {
	m.Config.ResourceGroupName = id
}

// GetMachineType returns a machine type the current project uses by default.
func (m *ResourceManager) GetMachineType() string {
	return m.Config.MachineType
}

// SetMachineType sets a machine type as the default one.
func (m *ResourceManager) SetMachineType(t string) {
	m.Config.MachineType = t
}

// MachineTypes returns a set of available machine types.
func (m *ResourceManager) MachineTypes(ctx context.Context) ([]cloud.MachineType, error) {
	s, err := NewBatchService(ctx, m.Config, m.Logger)
	if err != nil {
		return nil, err
	}
	return s.AvailableMachineTypes(ctx)
}

// GetRegion returns a region name the current project working on.
func (m *ResourceManager) GetRegion() string {
	return m.Config.Location
}

// SetRegion sets a region to the current project.
func (m *ResourceManager) SetRegion(region string) {
	m.Config.Location = region
}

// Regions returns a set of available regions.
func (m *ResourceManager) Regions(ctx context.Context) ([]cloud.Region, error) {

	m.Logger.Println("Retrieving available regions")
	regions, err := Locations(ctx, &m.Config.Token, m.Config.SubscriptionID)
	if err != nil {
		m.Logger.Println("Cannot retrieve available regions")
	} else {
		m.Logger.Println("Retrieved available regions")
	}
	return regions, err

}
