/*
	Implicit interfaces make dependency injection easier

	It is easy to implement dependency inejction in Go without additional libraries 
*/


package main

import (
	"errors"
	"fmt"
)

// LogOutput simply prints a message to stdout
func LogOutput(message string) {
	fmt.Println(message)
}

// SimpleDataStore struct stores user data in a map
type SimpleDataStore struct {
	userData map[string]string
}

// UserNameForID looks up a username by their userID
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDataStore creates a new instance of SimpleDataStore
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Jose",
			"4": "santi",
		},
	}
}

// DataStore interface represents a store for retrieving user data
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Logger interface represents a logger for logging messages
type Logger interface {
	Log(message string)
}

// LoggerAdapter adapts a function to implement the Logger interface
type LoggerAdapter func(message string)

// Log calls the function to log the message
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// SimpleLogic struct depends on a logger and a datastore
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// SayHello returns a greeting for the user with the given userID
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("In say hello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("Unknown user")
	}
	return "Hello, " + name, nil
}

// SayGoodBye returns a farewell message for the user with the given userID
func (sl SimpleLogic) SayGoodBye(userID string) (string, error) {
	sl.l.Log("In say goodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("Unknown user")
	}
	return "Goodbye, " + name, nil
}

// NewSimpleLogic creates a new instance of SimpleLogic
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// Main function to demonstrate usage
func main() {
	// Create a new data store
	dataStore := NewSimpleDataStore()

	// Create a logger using LoggerAdapter
	logger := LoggerAdapter(LogOutput)

	// Create a new SimpleLogic instance
	logic := NewSimpleLogic(logger, dataStore)

	// Test with a valid user ID
	greeting, err := logic.SayHello("1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(greeting)
	}

	// Test with an invalid user ID
	goodbye, err := logic.SayGoodBye("4")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(goodbye)
	}
}
