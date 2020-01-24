package tools

import (
	"os"
)

func WriteWorld(worldInBytes []byte) error {
	file, err := os.Create("worldxout.txt")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(worldInBytes)
	if err != nil {
		return err
	}

	return nil
}
