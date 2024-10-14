/*
	A common application of an empty interface is as a placeholder for data that read from an external source
*/


package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Create a map to store the parsed JSON data
	data := map[string]interface{}{}

	// Read the file contents into a byte slice
	contents, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err) // Handle error if file cannot be read
	}

	// Unmarshal the JSON data into the map
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Fatal(err) // Handle error if JSON parsing fails
	}

	// Print the data
	fmt.Println(data)
}

