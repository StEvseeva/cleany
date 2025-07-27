<p align="center">
  <a href="" rel="noopener">
</p>

<h3 align="center">Cleany</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> A RESTful API for managing hotel cleaning operations, including room bookings, cleaners, and cleaning orders.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

This service provides a full CRUD (Create, Read, Update, Delete) interface for managing:

Rooms: Track hotel rooms with floor and description.
Cleaners: Manage cleaner profiles (name, surname).
Bookings: Record guest bookings with check-in/check-out times.
Cleaning Orders: Assign cleaning tasks to cleaners, track completion status, and costs.
Built with Go, PostgreSQL, and OpenAPI 3.0 for API specification.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites
Before running the service, ensure you have:

Go 1.20+ installed
PostgreSQL 14+ with a configured database
Docker
Swagger UI / Postman (for API testing)

### Installing

1. Clone the Repository
bash
git clone https://github.com/yourusername/hotel-cleaning-api.git
cd hotel-cleaning-api

2. Configure Environment Variables
Create a .env file:

env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=hotel_cleaning
DB_HOST=localhost
DB_PORT=5432
PORT=8000

3. Install Dependencies
bash
go mod download

4. Run the Server
docker compose up -d

For more info about docker in these project use DOCKER_README

## 🔧 Running the tests <a name = "tests"></a>


### Break down into end to end tests

### And coding style tests

## 🎈 Usage <a name="usage"></a>

## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system.

## ⛏️ Built Using <a name = "built_using"></a>

- [PostgreSQL](https://www.postgresql.org/) - Database
- [Echo](https://echo.labstack.com/) - Server Framework

## ✍️ Authors <a name = "authors"></a>

- [@StEvseeva](https://github.com/StEvseeva) - Idea & Initial work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References
