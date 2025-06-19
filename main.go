package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"time"
)

func main() {
	ut, err := GetUserTempFolders()
	if err != nil {
		WriteLogFile(err)
	}

	fmt.Printf("Listing items inside '%s' folder...\n", ut[0])
	rtEntries, err := os.ReadDir(ut[0])
	if err != nil {
		WriteLogFile(err)
	}

	if len(rtEntries) == 0 {
		fmt.Printf("Found no item inside '%s'. Next...\n\n", ut[0])
	} else {
		fmt.Printf(
			"Found %d items inside '%s' folder. Starting to clean up...\n",
			len(rtEntries),
			ut[0],
		)

		if err := DeleteFolderEntries(ut[0], rtEntries); err != nil {
			WriteLogFile(err)
		}

		fmt.Printf("'%s' was cleaned! Next is '%s'\n\n", ut[0], ut[1])
	}

	fmt.Printf("Listing items inside '%s' folder...\n", ut[1])
	ptEntries, err := os.ReadDir(ut[1])
	if err != nil {
		WriteLogFile(err)
	}

	if len(ptEntries) == 0 {
		fmt.Printf("Found no item inside '%s'. Next...\n", ut[1])
	} else {
		fmt.Printf(
			"Found %d items inside '%s' folder. Starting to clean up...\n",
			len(ptEntries),
			ut[1],
		)

		if err := DeleteFolderEntries(ut[1], ptEntries); err != nil {
			WriteLogFile(err)
		}

		fmt.Printf("'%s' was cleaned!\n\n", ut[1])
	}

	fmt.Println("You can close this terminal, or press any key to finish it.")
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

func WriteLogFile(inErr error) {
	var errors []error
	u, err := user.Current()
	if err != nil {
		errors = append(errors, err)
	}
	f, err := os.OpenFile(
		path.Join(u.HomeDir, "TEMP", "temp"+time.Now().String()+".txt"),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		errors = append(errors, err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(inErr)
	for _, err := range errors {
		log.Println(err)
	}
}
