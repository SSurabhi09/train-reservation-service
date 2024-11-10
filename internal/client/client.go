package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Reservation represents a train reservation with user and ticket details
type Reservation struct {
    User        User   `json:"user"`
    Ticket Ticket `json:"ticket"`
}

// ReservationRequest represents the structure of the incoming reservation request
type CreateReservationRequest struct {
    Reservation Reservation `json:"reservation"`
}

// Ticket contains information about a train ticket
type Ticket struct {
    Seat    string `json:"seat"`
    From string `json:"from"`
    To string `json:"to"`
    Price float64 `json:"price"`
}

// User represents the details of a person making a reservation
type User struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email string `json:"email"`
}

// ReservationResponse represents the structure of the response sent back after a reservation request
type GetReservationResponse struct {
    Success            bool    `json:"success"`
    Message            string  `json:"message"`
    ReservationDetails struct {
        Ticket struct {
            From  string  `json:"from"`
            To    string  `json:"to"`
            Seat  string  `json:"seat"`
            Price float64 `json:"price"`
        } `json:"ticket"`
        User struct {
			// use camelCase proto compiler generated json field names
            FirstName string `json:"firstName"` 
            LastName  string `json:"lastName"`
            Email     string `json:"email"`
        } `json:"user"`
    } `json:"reservationDetails"`
}

type CreateReservationResponse struct {
    Success            bool    `json:"success"`
    Message            string  `json:"message"`
    ReservationId string `json:"reservationId"`
}

func main() {
    // Create the request body
    reqBody := CreateReservationRequest{
        Reservation: Reservation{
            Ticket: Ticket{
                From: "London",
                To:   "France",
                Seat: "2A",
            },
            User: User{
                FirstName: "John",
                LastName:  "Doe",
                Email: "johndoe@example.com",
            },
        },
    }

    // Marshal the request body to JSON
    jsonData, err := json.Marshal(reqBody)
    if err != nil {
        log.Fatalf("Error marshalling request body: %v", err)
    }
    //Send the HTTP POST request
    log.Println("Requesting for a reservation")
    resp, err := http.Post("http://localhost:7090/v1/reservations", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatalf("Error making request: %v", err)
    }
    defer resp.Body.Close()

    // Debugging purpose: Printing raw response
    // bodyBytes, _ := io.ReadAll(resp.Body)
    // fmt.Println("Raw response:", string(bodyBytes))

    // Parse the response
    var ticketResponse CreateReservationResponse
    if err := json.NewDecoder(resp.Body).Decode(&ticketResponse); err != nil {
        log.Fatalf("Error decoding response: %v", err)
    }

    // Output the response
    if ticketResponse.Success{
        fmt.Printf("%s\n", ticketResponse.Message)
        fmt.Printf("Reservation Id is: %s\n", ticketResponse.ReservationId)
    }
}
