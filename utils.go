package utils

import (
    "encoding/json"
    "net/http"
)

// Function to write JSON to an HTTP response
func ToJSON(w http.ResponseWriter, val interface{}) {
    w.Header().Set("Content-Type", "application/json")
    b, _ := json.Marshal(val)
    w.Write(b)
}
