package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

// DumpRequest - Function that dumps a passed HTTP Request object
// The function will return a byte slice and string
func DumpRequest(inRequest *http.Request) (returnString string, returnBytes []byte, returnError error) {
	log.Println("Dumping out the inbound request")

	// Fetch the request
	returnBytes, returnError = httputil.DumpRequest(inRequest, true)
	if returnError != nil {
		// Error occurred when dumping the request
		log.Println("Got an error attempting to dump the request", returnError)
	} else {
		// Save the dumped request as a string
		returnString = string(returnBytes)

		// Print the request to the log
		log.Println(returnString)
	}

	return
}

// ToJSONResponse - Function to write JSON to an HTTP response
func ToJSONResponse(w http.ResponseWriter, val interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, berr := json.Marshal(val)
	w.Write(b)

	return berr
}

// ToJSON - Function returns a JSON string.
// The function will return JSON (string), JSON (byte slice) and an error value
func ToJSON(val interface{}) (returnString string, returnBytes []byte, returnError error) {
	// Encode the value to JSON
	returnBytes, returnError = json.Marshal(val)
	if returnError == nil {
		// If there's no error then create a JSON string
		returnString = string(returnBytes)
	}

	return
}

// FromJSON - Function converts JSON (a string) to a referenced (pointer to a) data structure
func FromJSON(inVal string, outPtr interface{}) (retErr error) {
  // Convert the string to a byte slice
  tmpVal := []byte(inVal)
  
	// UnMarshall the byte array to a data structure pointed to by outPtr
	retErr = json.Unmarshal(tmpVal, outPtr)
  if retErr != nil {
    log.Println("Error converting json string to a data structure:", retErr)
  }

	return
}

// CheckErr - check for an error and print to the log if the passed error is not nil
func CheckErr(inError error, inString string) {
	// Check the passed error for nil
	if inError != nil {
		log.Println("An error occurred:", inError, inString)
	}

	return
}

// CheckErrBool - check for an error; print any non nil error to the log; and return a boolean
func CheckErrBool(inError error, inString string) (retBool bool) {
	// Initialize a return boolean value
	retBool = false

	// Check the passed error for nil
	if inError != nil {
		log.Println("An error occurred:", inError, inString)
		retBool = true
	}

	return
}

// Getport fetches the port number from an environment variable so we can run on Heroku
func Getport() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4567"
		log.Printf("INFO: No PORT environment variable detected, defaulting to: %v\n", port)
	}
	return fmt.Sprint(":", port)
}
