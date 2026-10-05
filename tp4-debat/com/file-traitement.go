package tp3web

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var wd, _ = os.Getwd()
var PATH = filepath.Join(wd, "env", "debat.md")

func StoreConversation(message string) {

	file, err := os.OpenFile(PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(message + "\n"))
	if err != nil {
		log.Fatal(err)
	}

}

func readLines() ([]string, error) {
	file, err := os.Open(PATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadAll() []byte {
	content, err := os.ReadFile(PATH)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	return content
}
