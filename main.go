package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFilesInDir(path string, upgradeStr string) {
	fmt.Println("==> reading files from directorory: " + path)

	// open directory
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// Read files in dir
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Loop through files in dir
	for _, file := range files {

		// open file
		openFile, err := os.Open(path + "/" + file.Name())
		serverName := strings.TrimSuffix(file.Name(), ".json")

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		// scan file
		scanner := bufio.NewScanner(openFile)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		// parses file line by line
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		openFile.Close()

		// read line by line
		contains := false
		for _, eachline := range txtlines {
			// Skip servers that do not have upgrades available
			if strings.Contains(eachline, upgradeStr) {
				contains = true
			}

			if !contains {
				continue
			}

			fmt.Println("\nServer: " + serverName)
			result := strings.Split(eachline, ",")
			for i := range result {

				if strings.Contains(result[i], "Inst") {
					fmt.Println(result[i])
				} else if strings.Contains(result[i], "  linux-headers-") {
					fmt.Println(result[i])
				}
			}
		}
	}
}

func main() {
	dataPath := "/tmp/upgradable-pkgs/"
	upgradeStr := "The following packages will be upgraded"

	readFilesInDir(dataPath, upgradeStr)
}
