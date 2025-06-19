package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	ut, err := GetUserTempFolders()
	if err != nil {
		log.Fatal(err)
	}

	rtEntries, err := os.ReadDir(ut[0])
	if err != nil {
		log.Fatal(err)
	}
	if err := DeleteFolderEntries(ut[0], rtEntries); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Regular temp cleaned!")

	ptEntries, err := os.ReadDir(ut[1])
	if err != nil {
		log.Fatal(err)
	}
	if err := DeleteFolderEntries(ut[1], ptEntries); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Percent temp cleaned!")

	fmt.Scanln()
}

func DeleteFolderEntries(basePath string, entries []os.DirEntry) error {
	for _, e := range entries {
		if err := os.RemoveAll(basePath + "\\" + e.Name()); err != nil {
			return err
		}
	}
	return nil
}

func GetUserTempFolders() ([]string, error) {
	const REGULAR_TEMP_PATH = "C:\\Windows\\Temp"
	u, err := user.Current()
	if err != nil {
		return []string{}, err
	}
	percentTempPath := fmt.Sprintf("%s\\AppData\\Local\\Temp", u.HomeDir)
	return []string{REGULAR_TEMP_PATH, percentTempPath}, nil
}
