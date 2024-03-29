package old

import (
	"fmt"
	"net/http"
	"time"
)

func test() {
	start := time.Now()

	res, err := http.Get("https://jnaraujo.com")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	elapsed := time.Since(start)

	fmt.Println(res.Status)
	fmt.Printf("It took %v to request", elapsed)
}