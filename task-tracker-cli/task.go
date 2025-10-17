package task_tracker_cli

import (
	"fmt"
	"time"
)

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

func postTask(description string) (int, error) {
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

	tasks[id] = newTask(description)

	err = loadToFile(tasks)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getTasks(status taskStatus) (map[int]*Task, error) {
	tasks, err := uploadFromFile()
	if err != nil {
		return tasks, nil
	}

	switch status {
	case TASK_STATUS_TODO:
		temp := make(map[int]*Task)
		for id, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				temp[id] = task
			}
		}
		return temp, nil
	case TASK_STATUS_IN_PROGRESS:
		temp := make(map[int]*Task)
		for id, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				temp[id] = task
			}
		}
		return temp, nil
	case TASK_STATUS_DONE:
		temp := make(map[int]*Task)
		for id, task := range tasks {
			if task.Status == TASK_STATUS_DONE {
				temp[id] = task
			}
		}
		return temp, nil
	default:
		return tasks, nil
	}
}

func (task Task) statusToString() string {
	switch task.Status {
	case TASK_STATUS_IN_PROGRESS:
		return "выполняется"
	case TASK_STATUS_DONE:
		return "выполнено"
	default:
		return "надо сделать"
	}
}

func deleteTask(id int) error {
	tasks, err := uploadFromFile()
	if err != nil {
		return nil
	}

	if _, ok := tasks[id]; !ok {
		return fmt.Errorf("задачи с ID(%d) нет", id)
	}
	delete(tasks, id)

	err = loadToFile(tasks)
	if err != nil {
		return err
	}

	return nil
}

func putTask(id int, description string) error {
	tasks, err := uploadFromFile()
	if err != nil {
		return err
	}

	if _, ok := tasks[id]; !ok {
		return fmt.Errorf("задачи с ID(%d) нет", id)
	}

	tasks[id].Description = description
	tasks[id].UpdatedAt = time.Now()

	err = loadToFile(tasks)
	if err != nil {
		return err
	}

	return nil
}
