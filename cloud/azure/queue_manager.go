package azure

import (
	"context"
	"fmt"
	"log"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

// QueueManager implements cloud.QueueManager interface to run a script
// on Azure.
type QueueManager struct {
	service *BatchService
	Config  *AzureConfig
	Logger  *log.Logger
}

// NewQueueManager creates a new queue manager.
func NewQueueManager(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (m *QueueManager, err error) {

	service, err := NewBatchService(ctx, cfg, logger)
	if err != nil {
		return
	}

	m = &QueueManager{
		service: service,
		Config:  cfg,
		Logger:  logger,
	}
	return

}

// Enqueue a new task to a given named queue.
func (m *QueueManager) Enqueue(ctx context.Context, queue string, task *script.Script) (err error) {

	err = m.service.CreateJob(ctx, queue)
	if err != nil {
		return
	}
	return m.service.CreateTask(ctx, queue, task)

}

// Tasks retrieves tasks in a given names queue.
func (m *QueueManager) Tasks(ctx context.Context, queue string, handler cloud.QueueManagerTaskHandler) (err error) {

	// set, err := m.service.Tasks(ctx, queue)
	// for _, v := range set {
	//   v.ResourceFiles
	//
	// }
	return fmt.Errorf("Not implemented")

}

// Queues retrieves existing queue names.
func (m *QueueManager) Queues(ctx context.Context, handler cloud.QueueManagerNameHandler) (err error) {

	set, err := m.service.Jobs(ctx)
	if err != nil {
		return
	}

	for name := range set {
		err = handler(name)
		if err != nil {
			return
		}
	}
	return

}

// Stop executing tasks in a given named queue.
func (m *QueueManager) Stop(ctx context.Context, queue string) error {

	return m.service.DisableJob(ctx, queue)

}

// Restart executing tasks in a given names queue.
func (m *QueueManager) Restart(ctx context.Context, queue string) error {

	return m.service.EnableJob(ctx, queue)

}

// CreateWorkers creates worker instances working for a given named queue.
func (m *QueueManager) CreateWorkers(ctx context.Context, queue string, diskSize int64, n int, handler cloud.QueueManagerNameHandler) (err error) {

	jobInfo, err := m.service.GetJobInfo(ctx, queue)
	if err != nil {
		return
	}
	pool := jobInfo.PoolInfo.PoolID
	poolInfo, err := m.service.GetPoolInfo(ctx, pool)
	if err != nil {
		return
	}
	return m.service.UpdatePoolSize(ctx, pool, poolInfo.TargetDedicated+int32(n))

}

// Workers retrieves worker instance names for a given queue.
func (m *QueueManager) Workers(ctx context.Context, queue string, handler cloud.QueueManagerNameHandler) (err error) {

	jobInfo, err := m.service.GetJobInfo(ctx, queue)
	if err != nil {
		return
	}
	pool := jobInfo.PoolInfo.PoolID
	nodes, err := m.service.Nodes(ctx, pool)
	if err != nil {
		return
	}

	for _, v := range nodes {
		err = handler(v.ID)
		if err != nil {
			break
		}
	}
	return

}
