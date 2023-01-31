package utils

import "encoding/json"

type TodoData struct {
	UserId int `json:"userId"`
	Id     int `json:"id"`
	Title  string `json:"title"`
	Completed bool `json:"completed"`
}

func JsonToMap(jsonString string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonString), &result)
	return result
}

func JsonToTodoData(jsonString string) TodoData {
	var result TodoData
	json.Unmarshal([]byte(jsonString), &result)
	return result
}