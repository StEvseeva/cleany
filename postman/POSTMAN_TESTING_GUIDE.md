# 🧪 Postman Testing Guide for Cleany API

This guide will help you test the Hotel Cleaning Service API using the provided Postman collection.

## 📋 Prerequisites

1. **Postman** installed on your machine
2. **PostgreSQL** database running
3. **Go application** running on `http://localhost:8080`
4. **Database tables** created (see database setup below)

## 🚀 Quick Start

### 1. Import the Collection

1. Open Postman
2. Click **Import** button
3. Select the `postman_collection.json` file
4. The collection will be imported with all endpoints organized by resource type

### 2. Set Up Environment Variables

The collection uses a variable `{{base_url}}` which is set to `http://localhost:8080` by default. You can modify this in:

1. Click on the collection name
2. Go to **Variables** tab
3. Update the `base_url` value if your server runs on a different port

### 3. Database Setup

Before testing, ensure your PostgreSQL database has the required tables:

```sql
-- Create rooms table
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    floor INTEGER NOT NULL,
    desc TEXT
);

-- Create cleaners table
CREATE TABLE cleaners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL
);

-- Create bookings table
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL REFERENCES rooms(id),
    check_in_ts TIMESTAMP,
    check_out_ts TIMESTAMP,
    guests INTEGER
);

-- Create cleaning_orders table
CREATE TABLE cleaning_orders (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER NOT NULL REFERENCES bookings(id),
    cleaning_type VARCHAR(100),
    cleaning_ts TIMESTAMP,
    notes TEXT,
    cost INTEGER NOT NULL,
    done BOOLEAN DEFAULT FALSE
);

-- Create cleaner_orders junction table
CREATE TABLE cleaner_orders (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES cleaning_orders(id),
    cleaner_id INTEGER NOT NULL REFERENCES cleaners(id),
    UNIQUE(order_id, cleaner_id)
);
```

## 🧪 Testing Scenarios

### **Recommended Testing Order**

1. **Rooms** - Start with room management
2. **Cleaners** - Add cleaning staff
3. **Bookings** - Create room bookings
4. **Cleaning Orders** - Manage cleaning tasks
5. **Test Scenarios** - Validate error handling

### **1. Rooms Testing**

#### Create a Room
- **Endpoint**: `POST /rooms`
- **Sample Data**:
```json
{
  "floor": 2,
  "desc": "Deluxe room with city view"
}
```
- **Expected Response**: 201 Created with room ID

#### Get All Rooms
- **Endpoint**: `GET /rooms`
- **Expected Response**: 200 OK with array of rooms

#### Get Room by ID
- **Endpoint**: `GET /rooms/1`
- **Expected Response**: 200 OK with room details

#### Update Room
- **Endpoint**: `PUT /rooms/1`
- **Sample Data**:
```json
{
  "floor": 3,
  "desc": "Updated deluxe room with mountain view"
}
```

### **2. Cleaners Testing**

#### Create a Cleaner
- **Endpoint**: `POST /cleaners`
- **Sample Data**:
```json
{
  "name": "John",
  "surname": "Doe"
}
```

#### Get All Cleaners
- **Endpoint**: `GET /cleaners`

### **3. Bookings Testing**

#### Create a Booking
- **Endpoint**: `POST /bookings`
- **Sample Data**:
```json
{
  "room_id": 1,
  "check_in_ts": "2024-01-15T14:00:00Z",
  "check_out_ts": "2024-01-17T11:00:00Z",
  "guests": 2
}
```

**Important**: Make sure the `room_id` exists before creating a booking!

### **4. Cleaning Orders Testing**

#### Create a Cleaning Order
- **Endpoint**: `POST /cleaning_orders`
- **Sample Data**:
```json
{
  "booking_id": 1,
  "cleaning_type": "deep_clean",
  "cleaning_ts": "2024-01-17T12:00:00Z",
  "notes": "Guest requested extra attention to bathroom",
  "cost": 150,
  "done": false
}
```

**Important**: Make sure the `booking_id` exists before creating a cleaning order!

#### Assign Cleaner to Order
- **Endpoint**: `POST /cleaning_orders/1/cleaners`
- **Sample Data**:
```json
{
  "cleaner_id": 1
}
```

## 🔍 Error Testing

The collection includes specific error scenarios:

### **Invalid Room ID**
- **Endpoint**: `GET /rooms/999`
- **Expected Response**: 404 Not Found

### **Invalid Booking Dates**
- **Endpoint**: `POST /bookings`
- **Sample Data**:
```json
{
  "room_id": 1,
  "check_in_ts": "2024-01-17T14:00:00Z",
  "check_out_ts": "2024-01-15T11:00:00Z",
  "guests": 2
}
```
- **Expected Response**: 400 Bad Request (check-in after check-out)

## 📊 Complete Workflow Test

Follow this sequence to test the complete workflow:

1. **Create Room** → Get room ID (e.g., 1)
2. **Create Cleaner** → Get cleaner ID (e.g., 1)
3. **Create Booking** → Use room ID from step 1
4. **Create Cleaning Order** → Use booking ID from step 3
5. **Assign Cleaner** → Use cleaner ID from step 2 and order ID from step 4
6. **Update Cleaning Order** → Mark as done
7. **Remove Cleaner** → Test unassignment
8. **Delete entities** → Test cleanup

## 🛠️ Troubleshooting

### **Common Issues**

1. **Connection Refused**
   - Ensure the Go application is running
   - Check if the port is correct (default: 8080)

2. **Database Connection Error**
   - Verify PostgreSQL is running
   - Check database credentials in `main.go`
   - Ensure database and tables exist

3. **404 Not Found**
   - Check if the resource ID exists
   - Verify the endpoint URL is correct

4. **400 Bad Request**
   - Check request body format
   - Verify required fields are provided
   - Ensure referenced IDs exist (room_id, booking_id, etc.)

### **Database Connection**

If you need to modify database settings, update the `DefaultConfig()` function in `internal/db/db.go`:

```go
func DefaultConfig() *Config {
    return &Config{
        Host:     "localhost",
        Port:     5432,
        User:     "your_username",
        Password: "your_password",
        DBName:   "cleany",
        SSLMode:  "disable",
    }
}
```

## 📝 Notes

- All timestamps should be in ISO 8601 format (e.g., `2024-01-15T14:00:00Z`)
- IDs are auto-generated by the database
- The API validates relationships (e.g., booking must reference existing room)
- Error responses include descriptive messages
- DELETE operations return 204 No Content on success

## 🎯 Success Criteria

A successful test run should demonstrate:

✅ All CRUD operations work for each resource  
✅ Proper error handling for invalid data  
✅ Relationship validation (foreign keys)  
✅ Business logic validation (e.g., check-in before check-out)  
✅ Cleaner assignment/unassignment functionality  

Happy testing! 🚀 