package task

import (
	"fmt"
	"time"
)

func LoadTasks() ([]Task, error) {
	return readCSV()
}

func SaveTasks(tasks []Task) error {
	return writeCSV(tasks)
}

func NextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}

	return max + 1
}

func AddTask(description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:          NextID(tasks),
		Description: description,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

	tasks = append(tasks, newTask)

	return SaveTasks(tasks)
}

func CompleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].IsComplete = true
			return SaveTasks(tasks)
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}

func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	found := false
	var newTask []Task

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTask = append(newTask, t)
	}

	if !found {
		return fmt.Errorf("task for the id %d was not found", id)
	}

	return SaveTasks(newTask)
}

func ListTasks() ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil

}
