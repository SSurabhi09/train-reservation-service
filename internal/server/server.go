package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	generated "github.com/SSurabhi09/train-reservation-service/internal/generated" // Import your generated package
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Implement the handler for the TicketingService
type ticketingServer struct {
	generated.UnimplementedTicketingServiceServer
}
var (
	// A global map to store reservations in memory (key: reservation ID, value: reservation details)
	reservationStore = make(map[string]*generated.Reservation)
)

// used for internal API to get all reservations
type getAllReservationResponse struct {
    Reservations []*generated.Reservation `json:"reservations"`
}


// Implement the CreateReservation method
func (s *ticketingServer) CreateReservation(ctx context.Context, req *generated.CreateReservationRequest) (
	*generated.CreateReservationResponse, error) {
	// Simple mock reservation logic
	fmt.Println("Raw request received:")

	// Print a readable version of the incoming request
	reqJSON, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("Error marshalling request: %v\n", err)
	} else {
		fmt.Printf("Request JSON: %s\n", reqJSON)
	}

	// Ensure that Reservation is not nil before accessing it
	if req.Reservation.User == nil {
		return nil, fmt.Errorf("user information is missing")
	}
	for v := range reservationStore {
		currenDetails := reservationStore[v]
		if currenDetails.Ticket.From == req.Reservation.Ticket.From && 
		currenDetails.User.Email == req.Reservation.User.Email &&
		currenDetails.Ticket.Seat == req.Reservation.Ticket.Seat {
            return nil, fmt.Errorf("a reservation for the given user and ticket already exists")
        }
	}
	// Logging the request fields
	fmt.Printf("Creating reservation for %s\n", req.Reservation.User.Email)
	reservationID := fmt.Sprintf("%s-%d", req.Reservation.User.Email,len(reservationStore)+1)

	// Save the reservation data in the in-memory store
	reservationStore[reservationID] = req.Reservation

	// Returning a mock response
	return &generated.CreateReservationResponse{
		Success: true,
		Message: "Reservation created successfully",
		ReservationId: reservationID,
	}, nil
}

// GetReservation will return the reservation details based on reservation ID
func (s *ticketingServer) GetReservation(ctx context.Context, req *generated.GetReservationRequest) (
	*generated.GetReservationResponse, error) {
	// Fetch the reservation based on the provided ID
	reservation, exists := reservationStore[req.ReservationId]
	// Debug logs
	fmt.Println("reservationStore--> server", reservationStore)
	if !exists {
		return nil, fmt.Errorf("reservation with ID %s not found", req.ReservationId)
	}

	// Returning the found reservation details
	return &generated.GetReservationResponse{
		Success:          true,
		Message:          "Reservation details fetched successfully",
		ReservationDetails: reservation,
	}, nil
}

// getAllReservations will return the reservation details of all reservations
// internal method
func (s *ticketingServer) getAllReservations() (
	*getAllReservationResponse, error) {
	if len(reservationStore) == 0 {
		return nil, fmt.Errorf("no reservations found")
	}
	var responses []*generated.Reservation
	for _, reservation := range reservationStore{
		response := &generated.Reservation{
            User: reservation.User,
            Ticket: reservation.Ticket,
        }
        responses = append(responses, response)
	}
	fmt.Println(responses)
	return &getAllReservationResponse{
		Reservations: responses,
	}, nil
}

// GetSeatAllocations retrieves the list of users and their allocated seats for a specific section
func (s *ticketingServer) GetSeatAllocations(ctx context.Context,
	section *generated.GetSeatAllocatedRequest) (
	*generated.GetSeatAllocatedResponse, error) {
	if len(reservationStore) == 0 {
		return nil, fmt.Errorf("no reservations found")
	}
	allReservations, err := s.getAllReservations()
	if err!= nil {
        return nil, fmt.Errorf("error getting all reservations: %v", err)
    }
	var seatAllocated []*generated.SeatAllocated
    for _, reservation := range allReservations.Reservations {
        if reservation.Ticket.Seat[len(reservation.Ticket.Seat)-1:] == section.Section {
            // append User to seatAllocated array
            seatAllocated = append(seatAllocated, &generated.SeatAllocated{
                User: reservation.User,
                Seat: reservation.Ticket.Seat,
            })
        }
    }
	fmt.Println(seatAllocated)
	return &generated.GetSeatAllocatedResponse{
		SeatAllocated: seatAllocated,
	}, nil
}

// GetSeatAllocations retrieves the list of users and their allocated seats for a specific section
func (s *ticketingServer) DeleteReservation(ctx context.Context,
	req *generated.DeleteReservationRequest) (
	*generated.DeleteReservationResponse, error) {
	if len(reservationStore) == 0 {
		return nil, fmt.Errorf("no reservations found")
	}
	// Fetch the reservation based on the provided ID
	_, exists := reservationStore[req.ReservationId]
	// Debug logs
	fmt.Println("reservationStore--> server", reservationStore)
	if !exists {
		return nil, fmt.Errorf("reservation with ID %s not found", req.ReservationId)
	}
	// Remove the reservation from the in-memory store
	delete(reservationStore, req.ReservationId)
    // Returning the mock response
    fmt.Println("Reservation deleted")
    fmt.Println("reservation deleted:", req.ReservationId)
    // Returning the found reservation details
    return &generated.DeleteReservationResponse{
        Success: true,
        Message: "Reservation deleted successfully",
    }, nil
}

func(s *ticketingServer) UpdateReservation(ctx context.Context,
	updateReq *generated.UpdateReservationRequest) (
		*generated.UpdateReservationResponse, error){
		// Fetch the reservation based on the provided ID
		reservation, exists := reservationStore[updateReq.ReservationId]
		if !exists {
			return &generated.UpdateReservationResponse{
                Success: false,
                Message: "Reservation with ID not found",
            }, nil
		}
		// Update the reservation details
        reservation.Ticket.Seat = updateReq.NewReservation.Seat

        // Debug logs
        fmt.Println("reservationStore--> server", reservationStore)
        // Returning the mock response
        fmt.Println("Reservation updated")
        fmt.Println("reservation updated:", updateReq.ReservationId)
        // Returning the found reservation details
        // Returning a success response as the update is successful
        // In a real-world scenario, this should also include the updated reservation details in the response.
        // For simplicity, we are returning a success response here.
        return &generated.UpdateReservationResponse{
            Success: true,
            Message: "Reservation updated successfully",
        }, nil
}

func main() {
	// Start gRPC server
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		// Create a gRPC server and register the ticketing service
		server := grpc.NewServer()
		generated.RegisterTicketingServiceServer(server, &ticketingServer{})

		// Start serving
		fmt.Println("Server started on port :50051")
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Create the gRPC Gateway mux (HTTP/REST handler)
	mux := runtime.NewServeMux()

	// Register the gRPC Gateway handler with the HTTP server
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := generated.RegisterTicketingServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway handler: %v", err)
	}

	// Start the HTTP server (REST API) on port 7090
	fmt.Println("Starting HTTP server on port 7090...")
	if err := http.ListenAndServe(":7090", mux); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
