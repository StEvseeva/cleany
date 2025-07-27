# 🐳 Docker Setup for Cleany API

This guide explains how to run the Cleany API using Docker and Docker Compose.

## 📋 Prerequisites

- **Docker** installed on your machine
- **Docker Compose** installed on your machine
- At least **2GB** of available RAM

## 🚀 Quick Start

### 1. Start Everything with One Command

```bash
docker-compose up -d
```

This command will:
- Build the Go application Docker image
- Start PostgreSQL database
- Start the Go API server
- Set up networking between containers
- Run database migrations automatically

### 2. Check Service Status

```bash
docker-compose ps
```

You should see both services running:
- `cleany-postgres` (PostgreSQL database)
- `cleany-app` (Go API server)

### 3. View Logs

```bash
# View all logs
docker-compose logs

# View specific service logs
docker-compose logs app
docker-compose logs postgres

# Follow logs in real-time
docker-compose logs -f app
```

### 4. Access the API

Once everything is running, you can access the API at:
- **API Base URL**: `http://localhost:8080`
- **PostgreSQL**: `localhost:5432`

## 🛠️ Available Commands

### Start Services
```bash
# Start in background
docker-compose up -d

# Start with logs visible
docker-compose up
```

### Stop Services
```bash
# Stop and remove containers
docker-compose down

# Stop and remove containers + volumes (⚠️ deletes data)
docker-compose down -v
```

### Rebuild Application
```bash
# Rebuild and restart
docker-compose up -d --build

# Rebuild only the app service
docker-compose up -d --build app
```

### Check Health
```bash
# Check service health
docker-compose ps

# View health check logs
docker-compose logs app | grep health
```

## 📁 Project Structure

```
cleany/
├── Dockerfile              # Go application container
├── docker-compose.yml      # Multi-service orchestration
├── .dockerignore          # Files to exclude from build
├── main.go                # Application entry point
├── internal/
│   ├── db/                # Database layer
│   ├── server/            # HTTP handlers
│   ├── service/           # Business logic
│   └── repository/        # Data access layer
└── docs/                  # OpenAPI specification
```

## 🔧 Configuration

### Environment Variables

The application uses these environment variables (set in `docker-compose.yml`):

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `postgres` | Database hostname |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database username |
| `DB_PASSWORD` | `password` | Database password |
| `DB_NAME` | `cleany` | Database name |
| `DB_SSLMODE` | `disable` | SSL mode |

### Custom Configuration

To use custom database settings, modify the `environment` section in `docker-compose.yml`:

```yaml
environment:
  - DB_HOST=your-db-host
  - DB_PORT=5432
  - DB_USER=your-user
  - DB_PASSWORD=your-password
  - DB_NAME=your-database
  - DB_SSLMODE=disable
```

## 🗄️ Database

### PostgreSQL Container

- **Image**: `postgres:15-alpine`
- **Port**: `5432`
- **Database**: `cleany`
- **User**: `postgres`
- **Password**: `password`
- **Data Persistence**: Stored in Docker volume `postgres_data`

### Database Migrations

Database migrations are automatically applied when the PostgreSQL container starts. The migration files are located in:
```
./internal/embed/migrations/
```

### Accessing the Database

```bash
# Connect to database from host
psql -h localhost -p 5432 -U postgres -d cleany

# Connect from within the postgres container
docker-compose exec postgres psql -U postgres -d cleany
```

## 🐛 Troubleshooting

### Common Issues

#### 1. Port Already in Use
```bash
# Check what's using the port
netstat -tulpn | grep :8080
netstat -tulpn | grep :5432

# Stop conflicting services or change ports in docker-compose.yml
```

#### 2. Database Connection Failed
```bash
# Check if postgres is running
docker-compose ps postgres

# Check postgres logs
docker-compose logs postgres

# Wait for postgres to be ready
docker-compose logs -f postgres | grep "database system is ready"
```

#### 3. Application Won't Start
```bash
# Check application logs
docker-compose logs app

# Check if database is accessible from app container
docker-compose exec app wget -qO- http://localhost:8080/rooms
```

#### 4. Build Failures
```bash
# Clean build cache
docker-compose build --no-cache

# Check Dockerfile syntax
docker build -t test .
```

### Health Checks

Both services include health checks:

- **PostgreSQL**: Checks if database is ready to accept connections
- **Application**: Checks if API responds to HTTP requests

### Logs and Debugging

```bash
# View all logs
docker-compose logs

# View specific service logs
docker-compose logs app
docker-compose logs postgres

# Follow logs in real-time
docker-compose logs -f

# View logs with timestamps
docker-compose logs -t
```

## 🔄 Development Workflow

### Making Code Changes

1. **Edit your code**
2. **Rebuild the application**:
   ```bash
   docker-compose up -d --build app
   ```
3. **Check the logs**:
   ```bash
   docker-compose logs -f app
   ```

### Testing with Postman

1. **Import the Postman collection**: `postman_collection.json`
2. **Set the base URL**: `http://localhost:8080`
3. **Run the tests**

### Database Changes

1. **Add migration files** to `internal/embed/migrations/`
2. **Restart the postgres service**:
   ```bash
   docker-compose restart postgres
   ```

## 🧹 Cleanup

### Remove Everything
```bash
# Stop and remove containers, networks, and volumes
docker-compose down -v

# Remove images
docker-compose down --rmi all

# Remove all unused Docker resources
docker system prune -a
```

### Keep Data
```bash
# Stop containers but keep volumes (preserves database data)
docker-compose down

# Start again with existing data
docker-compose up -d
```

## 📊 Monitoring

### Resource Usage
```bash
# Check container resource usage
docker stats

# Check disk usage
docker system df
```

### Performance
```bash
# Monitor API performance
curl -w "@curl-format.txt" -o /dev/null -s http://localhost:8080/rooms
```

## 🔒 Security Notes

- **Default passwords** are used for development only
- **SSL is disabled** for local development
- **Non-root user** runs the application container
- **Health checks** ensure service availability

For production deployment, consider:
- Using secrets management
- Enabling SSL/TLS
- Implementing proper authentication
- Using production-grade PostgreSQL configuration

## 🎯 Next Steps

1. **Test the API** using the Postman collection
2. **Review the logs** to ensure everything is working
3. **Customize configuration** as needed
4. **Deploy to production** with proper security measures

Happy coding! 🚀 