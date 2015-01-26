package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
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

// ToJSON - Function return a JSON string.
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
