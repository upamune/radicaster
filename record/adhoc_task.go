package record

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/xid"
)

type AdHocTaskStatus string

const (
	TaskPending   AdHocTaskStatus = "pending"
	TaskRecording AdHocTaskStatus = "recording"
	TaskCompleted AdHocTaskStatus = "completed"
	TaskFailed    AdHocTaskStatus = "failed"
)

type AdHocTask struct {
	ID           string
	StationID    string
	From         time.Time
	AreaID       string
	Status       AdHocTaskStatus
	Error        string
	FilePath     string
	ProgramTitle string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AdHocTaskManager struct {
	mu    sync.RWMutex
	tasks map[string]*AdHocTask
}

func NewAdHocTaskManager() *AdHocTaskManager {
	return &AdHocTaskManager{
		tasks: make(map[string]*AdHocTask),
	}
}

func (m *AdHocTaskManager) Create(stationID string, from time.Time, areaID string) *AdHocTask {
	m.mu.Lock()
	defer m.mu.Unlock()

	taskID := generateTaskID(stationID, from)
	task := &AdHocTask{
		ID:        taskID,
		StationID: stationID,
		From:      from,
		AreaID:    areaID,
		Status:    TaskPending,
		Error:     "",
		FilePath:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.tasks[taskID] = task
	return task
}

func (m *AdHocTaskManager) Update(taskID string, fn func(*AdHocTask)) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if task, ok := m.tasks[taskID]; ok {
		fn(task)
		task.UpdatedAt = time.Now()
	}
}

func (m *AdHocTaskManager) Get(taskID string) (*AdHocTask, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	task, ok := m.tasks[taskID]
	return task, ok
}

func (m *AdHocTaskManager) List(taskIDs []string) []*AdHocTask {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(taskIDs) == 0 {
		// 全タスクを返す
		tasks := make([]*AdHocTask, 0, len(m.tasks))
		for _, task := range m.tasks {
			tasks = append(tasks, task)
		}
		return tasks
	}

	// 指定されたタスクIDのみを返す
	tasks := make([]*AdHocTask, 0, len(taskIDs))
	for _, id := range taskIDs {
		if task, ok := m.tasks[id]; ok {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// CleanupOldTasks は1時間以上前の完了タスクを削除
func (m *AdHocTaskManager) CleanupOldTasks() {
	m.mu.Lock()
	defer m.mu.Unlock()

	cutoff := time.Now().Add(-1 * time.Hour)
	for id, task := range m.tasks {
		if (task.Status == TaskCompleted || task.Status == TaskFailed) &&
			task.UpdatedAt.Before(cutoff) {
			delete(m.tasks, id)
		}
	}
}

func generateTaskID(stationID string, from time.Time) string {
	return fmt.Sprintf("adhoc_%s_%s_%s",
		stationID,
		from.Format("20060102150405"),
		xid.New().String(),
	)
}
