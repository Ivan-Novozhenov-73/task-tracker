package task_tracker_cli

import (
	"encoding/json"
	"fmt"
	"os"
)

func uploadFromFile() (map[int]Task, error) {
	tasks := make(map[int]Task)

	file, err := os.OpenFile("tasks.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return tasks, err
	}
	defer file.Close()

	temp := make(map[string]Task)
	json.NewDecoder(file).Decode(&temp)

	for strID, task := range temp {
		var intID int
		fmt.Sscan(strID, &intID)
		tasks[intID] = task
	}

	return tasks, nil
}

func loadToFile(tasks map[int]Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(tasks)
	return nil
}
