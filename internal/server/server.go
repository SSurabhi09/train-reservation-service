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
