# train-reservation-service

A simple Train ticketing service

A basic implementation is as follows
train-reservation-service/
│
├── proto/
│ └── ticketing.proto # The protobuf file with service definition
│
├── server/ # Server-side implementation
│ └── server.go # For Go server
│
├── client/ # Client-side implementation
│ └── client.go # For Go client

#### Method Requests

1. POST : http://localhost:7090/v1/reservations
   Body:
   {"reservation":{"user":{"first_name":"John","last_name":"Doe","email":"johndoe@example.com"},
   "ticket":{"from":"London","to":"France","seat":"2A"}}}
2. GET : http://localhost:7090/v1/reservations
