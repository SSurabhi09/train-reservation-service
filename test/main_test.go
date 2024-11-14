package main

import (
	"testing"

	"github.com/SSurabhi09/train-reservation-service/internal/client"
	"github.com/stretchr/testify/assert"
)


func TestReservationFlow(t *testing.T) {
    req := `{"reservation":{"user":{"first_name":"John","last_name":"Doe","email":"johndoe@example.com"},
        "ticket":{"from":"London","to":"France","seat":"2A"}}}`
	
    reservation, err := client.CreateReservation(req)
    assert.NoError(t, err)
    assert.Equal(t, "johndoe@example.com-1", reservation.ReservationId)

    reservationDetails, err := client.GetReservation(reservation.ReservationId)
    assert.NoError(t, err)
    assert.Equal(t, "2A", reservationDetails.ReservationDetails.Ticket.Seat)
}