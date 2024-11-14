package main

import (
	"fmt"

	"github.com/SSurabhi09/train-reservation-service/internal/client"
)

// A gRPC client caller -- similar to Postman/Talend
// flow: main -> client -> server
func main() {
	req := `{"reservation":{"user":{"first_name":"John","last_name":"Doe","email":"johndoe@example.com"},
			"ticket":{"from":"London","to":"France","seat":"2A"}}}`
	reservation, err := client.CreateReservation(req)
	if err!= nil {
        fmt.Println("Error creating reservation:", err)
        return
    }
	fmt.Println("Created reservation:", reservation.ReservationId)
	reservation_id := reservation.ReservationId

	reservationDetails, err := client.GetReservation(reservation_id)
	if err!= nil {
        fmt.Println("Error getting reservation:", err)
        return
    }
	fmt.Printf("reservation details: %+v\n", reservationDetails.ReservationDetails)

	fmt.Println("Get all rsserrrr")

}