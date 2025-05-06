# Office Booking API

Simple API for booking office desks.

## Technologies
- Go
- Gin
- PostgreSQL
- pgx
- golang-migrate

## Usage

1. Create PostgreSQL database:
   
    createdb office_booking
    
2. Run migrations:
   
    migrate -path migrations -database "postgres://postgres:password@localhost:5432/office_booking?sslmode=disable" up
    
3. Run the server:
   
    go run cmd/main.go
    
4. Test endpoint:
    POST /bookings

   
    {
      "user_id": 1,
      "desk_id": 2,
      "date": "2025-05-06"
    }
