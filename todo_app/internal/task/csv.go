package task

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

const fileName = "tasks.csv"

var csvHeader = []string{
	"ID",
	"Description",
	"CreatedAt",
	"IsCompleted",
}

func loadFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open the file: %v", err)
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func readCSV() ([]Task, error) {
	f, err := loadFile(fileName)
	if err != nil {
		return nil, err
	}

	defer closeFile(f)

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) <= 1 {
		return []Task{}, nil
	}

	var tasks []Task

	for i, record := range records {
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		createdAt, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			return nil, err
		}
		isCompleted, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, Task{
			ID:          id,
			Description: record[1],
			CreatedAt:   createdAt,
			IsComplete:  isCompleted,
		})

	}

	return tasks, nil
}

func writeCSV(tasks []Task) error {
	f, err := loadFile(fileName)
	if err != nil {
		return err
	}

	defer closeFile(f)

	if err := f.Truncate(0); err != nil {
		return nil
	}

	if _, err := f.Seek(0, 0); err != nil {
		return nil
	}

	w := csv.NewWriter(f)

	if err := w.Write(csvHeader); err != nil {
		return err
	}

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(task.IsComplete),
		}

		if err := w.Write(record); err != nil {
			return nil
		}
	}

	w.Flush()
	return w.Error()

}
