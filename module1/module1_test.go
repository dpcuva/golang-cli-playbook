package module1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestModule1CheckEnvGOOS(t *testing.T) {
	found := OpenFileAndFindNthString("module1.txt", 0, "set GOOS")

	if !found {
		t.Errorf("'go env' does not work as expected")
	}
}

func TestModule1CheckEnvGOOSInJson(t *testing.T) {
	content, _ := ioutil.ReadFile("module1.json")
	var data map[string]interface{}

	err := json.Unmarshal(content, &data)
	if err != nil {
		t.Errorf("it looks 'module1.json' is not correctly formatted as a JSON file")
	}

	_, ok := data["GOOS"]
	if !ok {
		t.Errorf("your 'module1.json' does not include valid values that we check")
	}
	// fmt.Println(data["GOOS"].(string))

}

// OpenFileAndFindString opens a file and return if the given string is found or not
func OpenFileAndFindString(filename string, expected string) bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		trimmed := strings.Trim(t, " ")
		if trimmed == "" {
			continue
		}

		// matching logic
		if trimmed == expected {
			return true
		}
	}

	return false
}

// OpenFileAndFindNthString opens a file, look for Nth string splitted by a '=', and return if given expected string is found or not
func OpenFileAndFindNthString(filename string, nth int, expected string) bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		trimmed := strings.Trim(t, " ")
		if trimmed == "" {
			continue
		}

		// matching logic
		ss := strings.Split(trimmed, "=")
		fmt.Println(ss)
		if ss[nth] == expected {
			return true
		}
	}

	return false
}
