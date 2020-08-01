package module4

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
        "fmt"
        "github.com/codemodus/kace"
)

func TestModule4GoFormatContent(t *testing.T) {
	expected := "	github.com/codemodus/kace"
	found := OpenFileAndFindNthString("./go.mod.copy", 0, expected)
	if found != true {
		t.Errorf("the kace package is not found")
	}
}

// OpenFileAndFindNthString opens a file, look for Nth string splitted by a space, and return if given expected string is found or not
func OpenFileAndFindNthString(filename string, nth int, expected string) bool {
        fmt.Println(kace.Camel("test"))
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
		ss := strings.Split(trimmed, " ")
		if ss[nth] == expected {
                        fmt.Println(kace.Camel(ss[nth]))
			return true
		}
	}

	return false
}
