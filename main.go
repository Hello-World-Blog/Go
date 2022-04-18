package main

import (
	"encoding/json"
	"fmt"
)

// Mobile details
type Mobile struct {
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Price        int    `json:"price"`
	Currency     string `json:"currency"`
}

func main() {
	s1 := `{
		"model": "S20 FE",
		"manufacturer": "Samsung",
		"price": 45000,
		"currency": "INR"
		}`
	// Loading JSON into a well defined structure variable.
	var mobile Mobile
	err := json.Unmarshal([]byte(s1), &mobile)
	if err != nil {
		fmt.Println("Error unmarshalling!" + err.Error())
		return
	}
	fmt.Println("Successfully loaded JSON string into a well defined structure variable in memory")
	fmt.Printf("Details:\n%+v", mobile)
	// Loading JSON into a generic map.
	var data map[string]interface{}
	err = json.Unmarshal([]byte(s1), &data)
	if err != nil {
		fmt.Println("Error unmarshalling!" + err.Error())
		return
	}
	fmt.Println("\nSuccessfully loaded JSON into generic map in memory")
	fmt.Printf("Details:\n%v", data)
}
