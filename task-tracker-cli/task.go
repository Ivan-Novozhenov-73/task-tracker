package task_tracker_cli

import "time"

type taskStatus string

const (
	TASK_STATUS_TODO        taskStatus = "todo"
	TASK_STATUS_IN_PROGRESS taskStatus = "in-progress"
	TASK_STATUS_DONE        taskStatus = "done"
)

type Task struct {
	Description string
	Status      taskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func newTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func createTask(description string) (int, error) {
	tasks, err := uploadFromFile()
	if err != nil {
		return 0, err
	}

	id := len(tasks) + 1

	for {
		if _, ok := tasks[id]; !ok {
			break
		}
		id++
	}

	tasks[id] = *newTask(description)

	err = loadToFile(tasks)
	if err != nil {
		delete(tasks, id)
		return 0, err
	}

	return id, nil
}
