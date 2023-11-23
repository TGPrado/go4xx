package inputreader

import (
	"bufio"
	"log"
	"os"
)

func ReadInputAsync(fileName string, fileData chan<- string) {
	file := os.Stdin

	if fileName != "" {
		f, err := os.Open(fileName)
		if err != nil {
			log.Print("Error open file: ", fileName)
			os.Exit(0)
		}
		file = f
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fileData <- line
	}
}
