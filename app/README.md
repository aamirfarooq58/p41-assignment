## Repository Structure

- `Dockerfile`: Defines the container image for the time service application.
- `main.go`: Contains the main Go source code for the time service functionality.
- `README.md`: Provides documentation and usage instructions for the project.

## Usage Instructions

### Prerequisites

- Docker (version 20.10 or later recommended)
- Internet connection for pulling the base image

### Building the Docker Image

To build the Docker image for the time service, run the following command in the project root directory:

```bash
docker build -t timeservice .
```

This command builds a Docker image named `timeservice` based on the instructions in the `Dockerfile`.

### Running the Container

To start the time service container, use the following command:

```bash
docker run -p 8123:8123 timeservice
```

This command runs the container and maps port 8123 from the container to port 8123 on the host machine.

### Using Pre-built Docker Image

Alternatively, you can use the pre-built Docker image available on Docker Hub:

1. Pull the image:

```bash
docker pull aamir58/simpletimeservice:latest
```

2. Run the container:

```bash
docker run -d -p 8123:8123 aamir58/simpletimeservice:latest
```

3. Access the service:

```bash
curl http://localhost:8123
```

### Accessing the Service

Once the container is running, you can access the time service at:

```
http://localhost:8123
```

Replace `localhost` with the appropriate hostname or IP address if accessing the service from a different machine.

### Configuration

The time service runs on port 8123 by default. If you need to change this, you'll need to modify the `Dockerfile` and update the `EXPOSE` instruction accordingly.

## Data Flow

The time service application follows a simple request-response flow:

1. Client sends an HTTP request to the service on port 8123.
2. The Go application receives the request.
3. The application processes the request, generating the current time and extracts IP address.
4. The application sends an HTTP response back to the client with the current time and IP information.

```
[Client] <--HTTP Request--> [Docker Container (Port 8123)] <--> [Go Time Service Application]
           <--HTTP Response-->
```

## Infrastructure

The project uses Docker for containerization. The key infrastructure component is defined in the `Dockerfile`:

- **Base Image**: `golang:1.21-alpine`
  - Purpose: Provides a minimal Alpine Linux environment with Go 1.21 installed.

- **Application Binary**: `timeservice`
  - Type: Go executable
  - Purpose: The compiled time service application.

- **User**: `appuser` (non-root)
  - Purpose: Runs the application with reduced privileges for improved security.

- **Exposed Port**: 8123
  - Purpose: Allows external access to the time service.

This containerized setup ensures consistency across different environments and simplifies deployment and scaling of the time service.
