package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	generated "github.com/SSurabhi09/train-reservation-service/internal/generated"
	"github.com/stretchr/testify/assert"
)

func TestCreateReservationAPI(t *testing.T) {
	// Prepare the request body for creating a reservation
	reqBody := &generated.CreateReservationRequest{
		Reservation: &generated.Reservation{
			User: &generated.User{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
			},
			Ticket: &generated.Ticket{
				Seat:  "1A",
				From:  "London",
				To:    "France",
				Price: 100.0,
			},
		},
	}

	// Marshal the request body into JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshalling request body: %v", err)
	}

	// Send the POST request to the gRPC Gateway
	resp, err := http.Post("http://localhost:7090/v1/reservations", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Assert the response status code
	assert.Equal(t, 200, resp.StatusCode, "Expected 200 OK")

	// Decode the response body
	var ticketResponse generated.CreateReservationResponse
	err = json.NewDecoder(resp.Body).Decode(&ticketResponse)
	assert.Nil(t, err, "Error decoding response body")

	// Assert that the response contains a reservation ID
	assert.True(t, ticketResponse.Success)
	assert.Equal(t, "Reservation created successfully", ticketResponse.Message)
	assert.NotEmpty(t, ticketResponse.ReservationId, "Reservation ID should not be empty")
}
