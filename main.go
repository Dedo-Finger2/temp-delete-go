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

	for _, e := range rtEntries {
		os.RemoveAll(REGULAR_TEMP_PATH + "\\" + e.Name())
	}

	fmt.Println("Regular temp cleaned!")

	ptEntries, err := os.ReadDir(percentTempPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range ptEntries {
		os.RemoveAll(percentTempPath + "\\" + e.Name())
	}

	fmt.Println("Percent temp cleaned!")

	fmt.Scanln()
}
