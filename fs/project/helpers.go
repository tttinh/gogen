package project

import "os"

func createDir(path string) error {
	// check if AbsolutePath exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(path, 0754); err != nil {
			return err
		}
	}

	return nil
}
