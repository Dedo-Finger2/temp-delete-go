package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	const REGULAR_TEMP_PATH = "C:\\Windows\\Temp"

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	percentTempPath := fmt.Sprintf("%s\\AppData\\Local\\Temp", u.HomeDir)

	rtEntries, err := os.ReadDir(REGULAR_TEMP_PATH)
	if err != nil {
		log.Fatal(err)
	}
	if err := DeleteFolderEntries(REGULAR_TEMP_PATH, rtEntries); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Regular temp cleaned!")

	ptEntries, err := os.ReadDir(percentTempPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := DeleteFolderEntries(percentTempPath, ptEntries); err != nil {
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
