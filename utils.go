package utils

import (
    "encoding/json"
    "net/http"
)

// Function to write JSON to an HTTP response
func ToJSONResponse(w http.ResponseWriter, val interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    b, berr := json.Marshal(val)
    w.Write(b)
    
    return berr
}
