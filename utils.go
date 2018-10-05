package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var (
	rgxNonDigit *regexp.Regexp
	rgxDigits   *regexp.Regexp
)

// init executes code at initialization time to ensure proper execution of the utils functions
func init() {
	// rgxNonDigit is a regular expression representing a non-digit
	rgxNonDigit = regexp.MustCompile(`\D`)

	// rgxDigits is a regular expression representing a string that contains only digits
	rgxDigits = regexp.MustCompile(`^\d+$`)
}

// Contains determines if a string is an element within a slice of strings
func Contains(s string, slc []string) bool {
	retBool := false

	for _, v := range slc {
		if s == v {
			// Passed string is present in the passed slice of strings; break out of loop and return
			retBool = true
			break
		}
	}

	return retBool
}

// GetFromParm fetches a string value from a map or an empty string if the key does not exist
//   use for objects like parameter maps (i.e. map[string]string)
func GetFromParm(k string, m map[string]string) string {
	emptyVal := ""

	if v, ok := m[k]; ok {
		// Key exists, return the associated value
		return v
	}

	// Key does not exist, return an empty string
	return emptyVal
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

// ToJSONReader - Function to write JSON to a reader that can be invoked downstream
// ToJSONReader takes a string and byte slice and will process the string if non-empty or otherwise will process the byte slice
func ToJSONReader(obj *interface{}) (*bytes.Reader, error) {
	var (
		err    error
		tmpVal []byte
		bufrdr *bytes.Reader
	)

	// Marshal the object into a json byte slice
	tmpVal, err = json.Marshal(obj)
	if err != nil {
		// error marshaling a go object into a byte slice
		return bufrdr, fmt.Errorf("error marshaling a go object into a byte slice; see: %v", err)
	}

	// Generate a reader to enable reading of the json byte slice downstream
	return bytes.NewReader(tmpVal), nil
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
	var retVal string
	var port = os.Getenv("PORT")

	switch {

	// CASE: no port variable set; return a default value
	case port == "":
		retVal = ":4567"
		log.Printf("INFO: No PORT environment variable detected, defaulting to: %v\n", retVal)

	// CASE: port variable is a numeric only value (e.g. '8080'); return with prepended colon
	case rgxDigits.MatchString(port):
		retVal = fmt.Sprintf(":%v", port)

	// CASE: port variable is NOT a numeric only value (e.g. 'localhost:7878'); return value
	case !rgxDigits.MatchString(port):
		retVal = port
	}

	return retVal
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

	tmpStr = rgxNonDigit.ReplaceAllString(inVar, "")
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

// FileLine returns a string reflecting the filename and line number of the function call
func FileLine() string {
	var file = "unknown file"
	var line = -1
	var filepath []string

	if _, f, l, ok := runtime.Caller(1); ok {
		filepath = strings.Split(f, "/")
		line = l

		file = filepath[len(filepath)-1]
	}

	return fmt.Sprintf("%v: %v", file, line)
}
