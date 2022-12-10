package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("couldn't close request's body: %w", err)
		}
	}(response.Body)

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("couldn't close file: %w", err)
		}
	}(file)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
