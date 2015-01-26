package utils

import (
    "encoding/json"
    "net/http"
)

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
