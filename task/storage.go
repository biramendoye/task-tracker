package task

import (
	"encoding/json"
	"fmt"
	"os"
)

func (list *TaskList) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (list *TaskList) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s does not exist, starting with an empty task list.\n", filename)
			return nil
		}
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(list)
	if err != nil {
		return err
	}

	return nil
}
