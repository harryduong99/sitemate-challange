package apis

import (
	"encoding/json"
	"net/http"
)

type ResponseFormater struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

func ResponseWithError(response http.ResponseWriter, statusCode int, msg string) {
	ResponseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func ResponseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
