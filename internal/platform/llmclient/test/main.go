package main

import (
	"fmt"

	"llmclient"
)

func main() {
	client := llmclient.NewClient("http://localhost:8090")

	resp, err := client.Plan(llmclient.PlanRequest{
		DaySessionID: "123",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", resp)
}