# train-reservation-service

A simple Train ticketing service

A basic implementation is as follows
         train-reservation-service/internal

         │
         ├── proto/
         │ └── ticketing.proto # The protobuf file with service definition
         │
         ├── server/ # Server-side implementation
         │ └── server.go # For Go server
         │
         ├── client/ # Client-side implementation
         │ └── client.go # For Go client
         |
         |__ generated # All protobuf generated code
         |

#### Method Requests

1. POST : http://localhost:7090/v1/reservations

       Body:
            {"reservation":{"user":{"first_name":"John","last_name":"Doe","email":"johndoe@example.com"},
            "ticket":{"from":"London","to":"France","seat":"2A"}}}
       Response:
            {
            "success": true,
            "message": "Reservation created successfully",
            "reservationId": "johndoe@example.com-1"
            }
3. GET : http://localhost:7090/v1/reservations/{reservation_id}

         Response:
               {
                  "success": true,
                  "message": "Reservation details fetched successfully",
                  "reservationDetails":{
                  "user":{
                  "firstName": "John",
                  "lastName": "Doe",
                  "email": "johndoe@example.com"
                  },
                  "ticket":{"from": "London", "to": "France", "price": 0, "seat": "2A"…}
                  }
               }
