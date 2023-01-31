package main

import (
	"fmt"
	"io"
	"net/http"

	"jnaraujo.com/testmod/src/utils"
)

func main() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s", err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s", err)
	}

	fmt.Println(string(data))

	js := utils.JsonToTodoData(string(data))

	fmt.Println(js.Title)

}