package files

import "os"

func CheckFolder(path string) bool {

	var exist bool = false
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		exist = true
	}
	return exist
}

func CreateFolder(path string) error {

	return os.MkdirAll(path, os.ModePerm)

}

func DeleteFolder(path string) error {

	return os.RemoveAll(path)

}

func CleanFolder(path string) error {

	if err := DeleteFolder(path); err == nil {
		return CreateFolder(path)
	} else {
		return err
	}

}
