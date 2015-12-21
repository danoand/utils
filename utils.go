package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"strings"
)

var rgxNumDigit *regexp.Regexp

// init executes code at initialization time to ensure proper execution of the utils functions
func init() {
	// rgxNumDigit is a regular expression representing a numeric digit
	rgxNumDigit = regexp.MustCompile("\\D")
}

// DumpRequest - Function that dumps a passed HTTP Request object
// The function will return a byte slice and string
func DumpRequest(inRequest *http.Request) (returnString string, returnBytes []byte, returnError error) {
	// Fetch the request
	returnBytes, returnError = httputil.DumpRequest(inRequest, true)
	if returnError != nil {
		// Error occurred when dumping the request
		log.Println("Got an error attempting to dump the request", returnError)
	} else {
		// Save the dumped request as a string
		returnString = string(returnBytes)
	}

	return
}

// DumpResponse - Function that dumps a passed HTTP Response object
// The function will return a byte slice and string
func DumpResponse(inResponse *http.Response) (returnString string, returnBytes []byte, returnError error) {
	// Fetch the response
	returnBytes, returnError = httputil.DumpResponse(inResponse, true)
	if returnError != nil {
		// Error occurred when dumping the request
		log.Println("Got an error attempting to dump the request", returnError)
	} else {
		// Save the dumped request as a string
		returnString = string(returnBytes)
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

// ToJSONnStatusResponse - Function to write JSON and a status code to an HTTP response
func ToJSONnStatusResponse(w http.ResponseWriter, cde int, val interface{}) error {
	// Set a header indicating a json formatted response body
	w.Header().Set("Content-Type", "application/json")

	// Set the status code to be consumed by the client
	w.WriteHeader(cde)

	// Marshal the body into json and write to the response
	b, berr := json.Marshal(val)
	w.Write(b)

	return berr
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

// FromJSONBytes - Function converts JSON (a byte array) to a referenced (pointer to a) data structure
func FromJSONBytes(inVal []byte, outPtr interface{}) (retErr error) {
	// UnMarshall the byte array to a data structure pointed to by outPtr
	retErr = json.Unmarshal(inVal, outPtr)
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

// GetEnvVar fetches the value of a local environment variable or returns an empty string if it does not exist
func GetEnvVar(inVar string) string {
	// Validate that an argument has been passed to the function
	if len(inVar) == 0 {
		log.Printf("WARN: Missing environment variable name.  No value to be found.\n")
		return ""
	}

	// Return the value of the environment variable
	return os.Getenv(inVar)
}

// FormatPhoneUS accepts a string containing 10 numeric digits and returns a US formatted phone number
func FormatPhoneUS(inVar string) (string, error) {
	// Declare working variables
	var tmpStr string
	var tmpStrSlc []string
	var retStr string
	var retErr error

	tmpStr = rgxNumDigit.ReplaceAllString(inVar, "")
	if len(tmpStr) != 10 {
		// The phone number to be formatted does not contain 10 digits (area code + number)
		retErr = fmt.Errorf("inbound parameter [%v] does not contain 10 numeric digits", inVar)
		return "", retErr
	}

	// Construct the return string
	tmpStrSlc = strings.Split(tmpStr, "")
	retStr = fmt.Sprintf("(%v) %v-%v", strings.Join(tmpStrSlc[0:3], ""), strings.Join(tmpStrSlc[3:6], ""), strings.Join(tmpStrSlc[6:10], ""))
	return retStr, retErr
}
