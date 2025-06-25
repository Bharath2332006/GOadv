package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	inputFiles := []string{"server1.log", "server2.log", "server3.log"}
	err := ProcessLogs(inputFiles, "errors.log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Error logs have been processed successfully.")
}

func ProcessLogs(inputFiles []string, outputFile string) error {

	errorChan := make(chan string)

	var wg sync.WaitGroup

	for _, file := range inputFiles {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			err := readLogFile(filename, errorChan)
			if err != nil {
				log.Printf("Failed to read %s: %v", filename, err)
			}
		}(file)
	}

	writeErr := make(chan error, 1)
	go func() {
		err := writeErrorsToFile(outputFile, errorChan)
		writeErr <- err
	}()

	go func() {
		wg.Wait()
		close(errorChan)
	}()

	return <-writeErr
}

func readLogFile(filename string, errorChan chan<- string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			errorChan <- line
		}
	}
	return scanner.Err()
}

func writeErrorsToFile(outputFile string, errorChan <-chan string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for line := range errorChan {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
