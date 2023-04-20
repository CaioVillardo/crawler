package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/time/rate"
)

const baseUrl = "https://api.movidesk.com/public/v1"
const apiKey = "d495699b-f706-4ca6-9149-51012215ace3"
const route = "/services"

func fetchData() {
	limiter := rate.NewLimiter(rate.Limit(10), 1)
	url := fmt.Sprintf("%s%s", baseUrl, route)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", apiKey)

	for {
		err = limiter.Wait(req.Context())
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Erro na requisição: %s\n", resp.Status)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = ioutil.WriteFile("tickets.json", body, 0644)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Os dados foram salvos com sucesso!")
		break
	}

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "mydatabase"
	)

	type Ticket struct {
		Id                   int    `json:"id"`
		Title                string `json:"title"`
		Status               string `json:"status"`
		Priority             string `json:"priority"`
		Name                 string `json:"name"`
		Description          string `json:"description"`
		ParentServiceId      int    `json:"parentServiceId"`
		ServiceForTicketType int    `json:"serviceForTicketType"`
		IsVisible            int    `json:"isVisible"`
		AllowSelection       int    `json:"allowSelection"`
		AllowFinishTicket    bool   `json:"allowFinishTicket"`
		IsActive             bool   `json:"isActive"`
		AutomationMacro      string `json:"automationMacro"`
		DefaultCategory      string `json:"defaultCategory"`
		DefaultUrgency       string `json:"defaultUrgency"`
		AllowAllCategories   bool   `json:"allowAllCategories"`
	}

	var ticketsJs []Ticket

	jsonData, err := ioutil.ReadFile("tickets.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &ticketsJs)
	if err != nil {
		panic(err)
	}

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5436, "postgres", "era.a", "Crawler")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, ticket := range ticketsJs {
		sqlStatement := `
	        INSERT INTO tickets (id, title, status, priority, name, description, parentServiceId,
			serviceForTicketType, isVisible, allowSelection, allowFinishTicket, isActive, automationMacro, defaultCategory,
			defaultUrgency,allowAllCategories)
	        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
		_, err = db.Exec(sqlStatement, ticket.Id,
			ticket.Title, ticket.Status, ticket.Priority, ticket.Name, ticket.Description, ticket.ParentServiceId,
			ticket.ServiceForTicketType, ticket.IsVisible, ticket.AllowSelection, ticket.AllowFinishTicket, ticket.IsActive, ticket.AutomationMacro,
			ticket.DefaultCategory, ticket.DefaultUrgency, ticket.AllowAllCategories)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	fetchData()
}
