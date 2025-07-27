-- Create tables for Cleany API

-- Cleaners
CREATE TABLE IF NOT EXISTS cleaners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL
);

-- Rooms
CREATE TABLE IF NOT EXISTS rooms (
    id SERIAL PRIMARY KEY,
    floor INTEGER NOT NULL,
    "desc" VARCHAR(255)
);

-- Bookings
CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    check_in_ts TIMESTAMP,
    check_out_ts TIMESTAMP,
    guests INTEGER
);

-- Cleaning Orders
CREATE TABLE IF NOT EXISTS cleaning_orders (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    cleaning_type VARCHAR(100),
    cleaning_ts TIMESTAMP,
    notes TEXT,
    cost INTEGER NOT NULL,
    done BOOLEAN DEFAULT FALSE
);

-- Cleaner Orders junction table
CREATE TABLE IF NOT EXISTS cleaner_orders (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES cleaning_orders(id) ON DELETE CASCADE,
    cleaner_id INTEGER NOT NULL REFERENCES cleaners(id) ON DELETE CASCADE,
    UNIQUE(order_id, cleaner_id)
); 