package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/browser"
)

const FILENAME string = ".www"

const HELP string = ` _       _ __ _
| |     / / / /
| | /| / / / /
| |/ |/ / / /
|__/|__/_/_/

www, a speed dial for your browser, in your terminal
https://github.com/jkulton/www

Usage: www <bookmark>

Set up:
  1. Create a JSON file titled '.www' in your home directory
  2. Add a single object to the file, where each key is the name of your bookmark, and each value is the URL
  3. Execute www with the name of your bookmark

Example .www file:
{
  "gh": "https://github.com/"
}
`

func getBookmarks() (map[string]string, error) {
	var bookmarks map[string]string

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error trying to access home directory: %w", err)
	}

	path := filepath.Join(home, FILENAME)
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("error trying to open %s file: %w", FILENAME, err)
	}

	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error trying to read from %s file: %w", FILENAME, err)
	}

	// TODO[jkulton]: Is json.Unmarshal performant? Use something other than JSON? (toml?)
	err = json.Unmarshal(fileContent, &bookmarks)

	if err != nil {
		return nil, fmt.Errorf("error parsing JSON from %s file: %w", FILENAME, err)
	}

	return bookmarks, nil
}

func main() {
	var keyword string
	bookmarks, err := getBookmarks()

	if err != nil {
		fmt.Println("error looking up bookmarks: ", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		keyword = os.Args[1]
	}

	if keyword == "" {
		fmt.Print(HELP)
		os.Exit(1)
	}

	url, found := bookmarks[keyword]

	if !found {
		fmt.Printf("Bookmark '%s' not found, please ensure '%s' is set in your %s file\n", keyword, keyword, FILENAME)
		os.Exit(1)
	}

	browser.OpenURL(url)
}
