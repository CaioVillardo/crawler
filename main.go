package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type MovideskTicket struct {
	ID          int    `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

func main() {
	baseURL := "https://desbravador.movidesk.com/public/v1."
	companyName := "desbravador"
	apiToken := "d495699b-f706-4ca6-9149-51012215ace3"

	client := &http.Client{Timeout: 10 * time.Second}

	endpoint := fmt.Sprintf("%s/%s/search/ticket", baseURL, companyName)

	values := url.Values{}
	values.Add("$select", "id,subject,description")
	values.Add("$top", "50")

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.URL.RawQuery = values.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var tickets []MovideskTicket
	err = json.NewDecoder(resp.Body).Decode(&tickets)
	if err != nil {
		panic(err)
	}

	for _, ticket := range tickets {
		fmt.Printf("Ticket ID: %d\n", ticket.ID)
		fmt.Printf("Subject: %s\n", ticket.Subject)
		fmt.Printf("Description: %s\n\n", ticket.Description)
	}
}
