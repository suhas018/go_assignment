# AdService with Logging Middleware

This repository contains a simple Go application implementing a logging-enabled service that demonstrates clean separation of business logic and transport mechanisms. The application provides a basic HTTP API for handling requests related to advertisement placements and bid price generation.

---

## Features

- **Service Interface**: Encapsulates business logic through the `AddService` interface.
- **Logging Middleware**: Logs request and response data, along with execution time.
- **UUID-based Identification**: Generates and utilizes unique IDs for each advertisement placement request.
- **HTTP API**: Exposes a REST endpoint for handling `POST /add` requests.
- **Random Bid Price**: Generates a random bid price for each request.
- **Transport-Business Logic Separation**: Demonstrates clean code architecture.

---

## Project Structure

```
.
├── main.go         # Entry point of the application
├── README.md       # Project documentation
```

---

## Prerequisites

- **Go**: Ensure that Go is installed on your system. You can download it [here](https://golang.org/dl/).
- **Go Modules**: The project uses `go.mod` for dependency management.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/<your-github-username>/addservice-logging.git
   cd addservice-logging
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

---

## Usage

1. Run the application:

   ```bash
   go run main.go
   ```

2. The server will start on `http://localhost:3000`.

3. Make a POST request to the `/add` endpoint. Example:

   ```bash
   curl -X POST http://localhost:3000/add
   ```

4. A successful response will return a JSON object containing:

   - `addID`: A unique ID for the advertisement.
   - `bidPrice`: The generated bid price.

---

## Response Example

```json
{
  "addID": "bdf0c8cd-7a54-4c91-9269-79a0d589f4b5",
  "bidPrice": 70.134563
}
```

---

## Design Highlights

- **Clean Architecture**: Separates business logic (via the `AddService` interface) and transport logic (via HTTP handlers).
- **Middleware**: Implements a logging middleware to log request/response details and execution time for better observability.
- **Error Handling**: Centralized error handling ensures that the application responds gracefully to unexpected issues.

---

## Dependencies

- [google/uuid](https://github.com/google/uuid): For generating unique IDs.
- [log/slog](https://pkg.go.dev/log/slog): For structured logging.

---

