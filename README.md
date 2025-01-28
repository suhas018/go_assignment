AddService with Logging Middleware

This repository contains a simple Go application implementing a logging-enabled service that demonstrates clean separation of business logic and transport mechanisms. The application provides a basic HTTP API for handling requests related to advertisement placements and bid price generation.

Features

Service Interface: Encapsulates business logic through the AddService interface.

Logging Middleware: Logs request and response data, along with execution time.

UUID-based Identification: Generates and utilizes unique IDs for each advertisement placement request.

HTTP API: Exposes a REST endpoint for handling POST /add requests.

Random Bid Price: Generates a random bid price for each request.

Transport-Business Logic Separation: Demonstrates clean code architecture.